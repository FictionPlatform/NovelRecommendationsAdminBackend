package queue

import (
	"go-admin/core/utils/idgen"
	"go-admin/core/utils/log"
	"go-admin/core/utils/storage"
	"sync"
	"time"
)

// maxConsumerRetries 单条消息最大重试次数，超过后丢弃并记录错误，防止坏消息永久占用消费协程
const maxConsumerRetries = 3

// enqueueTimeout 入队最长等待时间，超时丢弃，避免消费过慢时 goroutine 无限堆积
const enqueueTimeout = 100 * time.Millisecond

type queue chan storage.Messager

// NewMemory 内存模式
func NewMemory(poolNum uint) *Memory {
	return &Memory{
		queue:   new(sync.Map),
		PoolNum: poolNum,
	}
}

type Memory struct {
	queue   *sync.Map
	wait    sync.WaitGroup
	mutex   sync.RWMutex
	PoolNum uint
}

func (*Memory) String() string {
	return "memory"
}

func (m *Memory) makeQueue() queue {
	if m.PoolNum <= 0 {
		return make(queue)
	}
	return make(queue, m.PoolNum)
}

func (m *Memory) Append(message storage.Messager) error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	memoryMessage := new(Message)
	memoryMessage.SetID(message.GetID())
	memoryMessage.SetStream(message.GetStream())
	memoryMessage.SetValues(message.GetValues())
	v, ok := m.queue.Load(message.GetStream())
	if !ok {
		v = m.makeQueue()
		m.queue.Store(message.GetStream(), v)
	}
	var q queue
	switch v.(type) {
	case queue:
		q = v.(queue)
	default:
		q = m.makeQueue()
		m.queue.Store(message.GetStream(), q)
	}
	// 有界等待入队：队列满时最多等待 enqueueTimeout 后丢弃，避免 goroutine 无限堆积导致内存耗尽
	go func(gm storage.Messager, gq queue) {
		gm.SetID(idgen.UUID())
		select {
		case gq <- gm:
		case <-time.After(enqueueTimeout):
			log.Errorf("memory queue [%s] is full, drop message", gm.GetStream())
		}
	}(memoryMessage, q)
	return nil
}

func (m *Memory) Register(name string, f storage.ConsumerFunc) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	v, ok := m.queue.Load(name)
	if !ok {
		v = m.makeQueue()
		m.queue.Store(name, v)
	}
	var q queue
	switch v.(type) {
	case queue:
		q = v.(queue)
	default:
		q = m.makeQueue()
		m.queue.Store(name, q)
	}
	// 消费失败在本地重试有限次数后丢弃，不再回投到正在读取的同一 channel（原实现会自死锁/无限重试）
	go func(q queue, gf storage.ConsumerFunc) {
		for message := range q {
			var err error
			for retry := 0; retry < maxConsumerRetries; retry++ {
				if err = gf(message); err == nil {
					break
				}
				// 短暂退避，避免坏消息造成忙等
				time.Sleep(50 * time.Millisecond)
			}
			if err != nil {
				log.Errorf("memory queue consume message error, drop: %s", err.Error())
			}
		}
	}(q, f)
}

func (m *Memory) Run() {
	m.wait.Add(1)
	m.wait.Wait()
}

func (m *Memory) Shutdown() {
	m.wait.Done()
}
