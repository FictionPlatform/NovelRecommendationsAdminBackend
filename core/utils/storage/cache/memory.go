package cache

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"go-admin/core/runtime"
	"strconv"
	"sync"
	"time"
)

type item struct {
	Value   interface{}
	Expired *time.Time
}

// NewMemory memory模式
func NewMemory() *Memory {
	m := &Memory{
		items: new(sync.Map),
	}
	go m.cleanupLoop()
	return m
}

// cleanupInterval 过期条目后台清理间隔
const cleanupInterval = 30 * time.Second

// cleanupLoop 定期主动清理过期条目：惰性删除只能清理被访问过的 key，
// 验证码等随机 key 必须由后台协程回收，否则内存无限增长
func (m *Memory) cleanupLoop() {
	ticker := time.NewTicker(cleanupInterval)
	defer ticker.Stop()
	for range ticker.C {
		now := time.Now()
		m.items.Range(func(k, v interface{}) bool {
			if it, ok := v.(*item); ok && it.Expired != nil && it.Expired.Before(now) {
				m.items.Delete(k)
			}
			return true
		})
	}
}

type Memory struct {
	items *sync.Map
	mutex sync.RWMutex
}

func (*Memory) String() string {
	return "memory"
}

func (m *Memory) connect() {
}

func (m *Memory) Exist(prefix, key string) bool {
	key = prefix + runtime.IntervalTenant + key
	item, err := m.getItem(key)
	if err != nil || item == nil {
		return false
	}
	return true
}

func (m *Memory) Get(prefix, key string) (string, error) {
	key = prefix + runtime.IntervalTenant + key
	item, err := m.getItem(key)
	if err != nil || item == nil {
		return "", err
	}
	v, e := cast.ToStringE(item.Value)
	if e != nil {
		return "", e
	}
	return v, nil
}

// Set
// @Description:
// @receiver m
// @param key
// @param val
// @param expire 单位秒，若小于0，则表示不清除，始终停留在内存中
// @return error
func (m *Memory) Set(prefix, key string, val interface{}, expire int) error {
	key = prefix + runtime.IntervalTenant + key
	s, err := cast.ToStringE(val)
	if err != nil {
		return err
	}

	var exp *time.Time
	if expire <= 0 {
		exp = nil
	} else {
		exp2 := time.Now().Add(time.Duration(expire) * time.Second)
		exp = &exp2
	}

	item := &item{
		Value:   s,
		Expired: exp,
	}
	return m.setItem(key, item)
}

func (m *Memory) Del(prefix, key string) error {
	key = prefix + runtime.IntervalTenant + key
	return m.del(key)
}

func (m *Memory) HashSet(expire int, prefix, key string, values map[string]interface{}) error {
	key = prefix + runtime.IntervalTenant + key
	var exp *time.Time
	if expire <= 0 {
		exp = nil
	} else {
		exp2 := time.Now().Add(time.Duration(expire) * time.Second)
		exp = &exp2
	}

	itemTmps := &item{
		Value:   values,
		Expired: exp,
	}

	return m.setItem(key, itemTmps)
}

func (m *Memory) HashGetAll(prefix, key string) (map[string]string, error) {
	key = prefix + runtime.IntervalTenant + key
	item, err := m.getItem(key)
	if err != nil {
		return nil, err
	}
	if item == nil || item.Value == nil {
		return nil, errors.New("data is empty")
	}
	tmp := (item.Value).(map[string]interface{})
	result := map[string]string{}
	for k, v := range tmp {
		result[k] = v.(string)
	}
	return result, nil

}

func (m *Memory) HashGet(prefix, key, field string) (string, error) {
	key = prefix + runtime.IntervalTenant + key
	item, err := m.getItem(key)
	if err != nil || item == nil {
		return "", err
	}
	v := ((item.Value).(map[string]interface{}))[field]
	if v == nil {
		return "", errors.New("value is empty")
	}

	return v.(string), nil
}

func (m *Memory) HashDel(prefix, key, field string) error {
	key = prefix + runtime.IntervalTenant + key
	it, err := m.getItem(key)
	if err != nil || it == nil {
		return err
	}
	// copy-on-write：不修改共享 map，避免与其他读方并发访问共享 *item 引发竞态
	oldMap, ok := (it.Value).(map[string]interface{})
	if !ok {
		return fmt.Errorf("value of %s type error", key)
	}
	newMap := make(map[string]interface{}, len(oldMap))
	for k, v := range oldMap {
		newMap[k] = v
	}
	delete(newMap, field)
	return m.setItem(key, &item{
		Value:   newMap,
		Expired: it.Expired,
	})
}

func (m *Memory) Increase(prefix, key string) error {
	key = prefix + runtime.IntervalTenant + key
	return m.calculate(key, 1)
}

func (m *Memory) Decrease(prefix, key string) error {
	key = prefix + runtime.IntervalTenant + key
	return m.calculate(key, -1)
}

// calculate 读改写必须使用写锁，RLock 会导致并发 Increase/Decrease 丢失更新；
// 采用 copy-on-write 替换整条记录，避免并发读写共享 *item 指针造成数据竞态
func (m *Memory) calculate(key string, num int) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	it, err := m.getItem(key)
	if err != nil {
		return err
	}

	if it == nil {
		err = fmt.Errorf("%s not exist", key)
		return err
	}
	var n int
	n, err = cast.ToIntE(it.Value)
	if err != nil {
		return err
	}
	n += num
	return m.setItem(key, &item{
		Value:   strconv.Itoa(n),
		Expired: it.Expired,
	})
}

func (m *Memory) Expire(prefix, key string, expire int) error {
	key = prefix + runtime.IntervalTenant + key
	m.mutex.Lock()
	defer m.mutex.Unlock()
	it, err := m.getItem(key)
	if err != nil {
		return err
	}
	if it == nil {
		err = fmt.Errorf("%s not exist", key)
		return err
	}

	exp := time.Now().Add(time.Duration(expire) * time.Second)
	// copy-on-write 替换整条记录，不修改共享 *item
	return m.setItem(key, &item{
		Value:   it.Value,
		Expired: &exp,
	})
}

func (m *Memory) getItem(key string) (*item, error) {
	var err error
	i, ok := m.items.Load(key)
	if !ok {
		return nil, nil
	}
	switch i.(type) {
	case *item:
		item := i.(*item)
		if item.Expired == nil {
			return item, nil
		}
		if item.Expired.Before(time.Now()) {
			//过期
			_ = m.del(key)
			//过期后删除
			return nil, nil
		}
		return item, nil
	default:
		err = fmt.Errorf("value of %s type error", key)
		return nil, err
	}
}

func (m *Memory) setItem(key string, item *item) error {
	m.items.Store(key, item)
	return nil
}

func (m *Memory) del(key string) error {
	m.items.Delete(key)
	return nil
}
