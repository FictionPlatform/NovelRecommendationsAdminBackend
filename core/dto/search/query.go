package search

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	// FromQueryTag tag标记
	FromQueryTag = "search"
)

// ResolveSearchQuery 解析
/**
 * 	exact / iexact 等于
 * 	contains / icontains 包含
 *	gt / gte 大于 / 大于等于
 *	lt / lte 小于 / 小于等于
 *	startswith / istartswith 以…起始
 *	endswith / iendswith 以…结束
 *	in
 *	isnull
 *  order 排序		e.g. order[key]=desc     order[key]=asc
 */
func ResolveSearchQuery(q interface{}, condition Condition) {
	if condition == nil || q == nil {
		return
	}
	qType := reflect.TypeOf(q)
	qValue := reflect.ValueOf(q)
	// 指针/空值防护：递归传入基础类型、nil 指针或未导出字段时安全返回
	for qType.Kind() == reflect.Ptr {
		if qValue.IsNil() {
			return
		}
		qType = qType.Elem()
		qValue = qValue.Elem()
	}
	if qType.Kind() != reflect.Struct {
		return
	}
	var tag string
	var ok bool
	var t *resolveSearchTag
	for i := 0; i < qType.NumField(); i++ {
		tag, ok = "", false
		tag, ok = qType.Field(i).Tag.Lookup(FromQueryTag)
		if !ok {
			// 无 search tag：仅对可导出结构体类型递归解析（基础类型/未导出字段直接跳过，避免 panic）
			if qValue.Field(i).CanInterface() {
				kind := qValue.Field(i).Kind()
				if kind == reflect.Struct || kind == reflect.Ptr {
					ResolveSearchQuery(qValue.Field(i).Interface(), condition)
				}
			}
			continue
		}
		switch tag {
		case "-":
			continue
		}
		t = makeTag(tag)
		if qValue.Field(i).IsZero() {
			continue
		}
		//解析
		switch t.Type {
		case "left":
			//左关联
			join := condition.SetJoinOn(t.Type, fmt.Sprintf(
				"left join  %s on  %s.%s =  %s.%s",
				t.Join,
				t.Join,
				t.On[0],
				t.Table,
				t.On[1],
			))
			ResolveSearchQuery(qValue.Field(i).Interface(), join)
		case "inner":
			//左关联
			join := condition.SetJoinOn(t.Type, fmt.Sprintf(
				"inner join %s on %s.%s = %s.%s",
				t.Join,
				t.Join,
				t.On[0],
				t.Table,
				t.On[1],
			))
			ResolveSearchQuery(qValue.Field(i).Interface(), join)
		case "exact", "iexact":
			condition.SetWhere(fmt.Sprintf("%s.%s = ?", t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
		case "contains", "icontains":
			condition.SetWhere(fmt.Sprintf("%s.%s like ?", t.Table, t.Column), []interface{}{"%" + qValue.Field(i).String() + "%"})
		case "leftcontains", "lefticontains":
			condition.SetWhere(fmt.Sprintf("%s.%s like ?", t.Table, t.Column), []interface{}{"%" + qValue.Field(i).String()})
		case "rightcontains", "righticontains":
			condition.SetWhere(fmt.Sprintf("%s.%s like ?", t.Table, t.Column), []interface{}{qValue.Field(i).String() + "%"})
		case "gt":
			condition.SetWhere(fmt.Sprintf("%s.%s > ?", t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
		case "gte":
			condition.SetWhere(fmt.Sprintf("%s.%s >= ?", t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
		case "lt":
			condition.SetWhere(fmt.Sprintf("%s.%s < ?", t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
		case "lte":
			condition.SetWhere(fmt.Sprintf("%s.%s <= ?", t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
		case "startswith", "istartswith":
			condition.SetWhere(fmt.Sprintf("%s.%s like ?", t.Table, t.Column), []interface{}{qValue.Field(i).String() + "%"})
		case "endswith", "iendswith":
			condition.SetWhere(fmt.Sprintf("%s.%s like ?", t.Table, t.Column), []interface{}{"%" + qValue.Field(i).String()})
		case "in":
			condition.SetWhere(fmt.Sprintf("%s.%s in (?)", t.Table, t.Column), []interface{}{qValue.Field(i).Interface()})
		case "isnull":
			// IsNil 仅对可空类型（chan/func/interface/map/pointer/slice）合法，非可空类型不调用避免 panic
			field := qValue.Field(i)
			switch field.Kind() {
			case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
				if !(field.IsZero() && field.IsNil()) {
					condition.SetWhere(fmt.Sprintf("%s.%s isnull", t.Table, t.Column), make([]interface{}, 0))
				}
			default:
				if !field.IsZero() {
					condition.SetWhere(fmt.Sprintf("%s.%s isnull", t.Table, t.Column), make([]interface{}, 0))
				}
			}
		case "order":
			switch strings.ToLower(qValue.Field(i).String()) {
			case "desc", "asc":
				condition.SetOrder(fmt.Sprintf("%s.%s %s", t.Table, t.Column, qValue.Field(i).String()))
			}
		}
	}
}
