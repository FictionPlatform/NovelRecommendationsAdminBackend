package search

import (
	"reflect"
	"testing"
)

type basicFields struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int
}

type nestedStruct struct {
	User basicFields
	Ptr  *basicFields
}

type nullableFields struct {
	Ptr   *int `search:"type:isnull;column:deleted_at;table:t"`
	Int   int  `search:"type:isnull;column:deleted_at;table:t"`
	Str   string
	Slice []string `search:"type:isnull;column:x;table:t"`
}

func TestResolveSearchQueryBasicTypeField(t *testing.T) {
	c := &GormCondition{GormPublic: GormPublic{}, Join: make([]*GormJoin, 0)}
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("basic type field without search tag panicked: %v", r)
		}
	}()
	ResolveSearchQuery(basicFields{ID: 1, Name: "x"}, c)
}

func TestResolveSearchQueryNestedStruct(t *testing.T) {
	c := &GormCondition{GormPublic: GormPublic{}, Join: make([]*GormJoin, 0)}
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("nested struct field panicked: %v", r)
		}
	}()
	ResolveSearchQuery(nestedStruct{User: basicFields{ID: 2}, Ptr: &basicFields{Name: "p"}}, c)
}

func TestResolveSearchQueryNilPointer(t *testing.T) {
	c := &GormCondition{GormPublic: GormPublic{}, Join: make([]*GormJoin, 0)}
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("nil pointer field panicked: %v", r)
		}
	}()
	var p *basicFields
	ResolveSearchQuery(nestedStruct{User: basicFields{ID: 2}, Ptr: p}, c)
}

func TestResolveSearchQueryNonStruct(t *testing.T) {
	c := &GormCondition{GormPublic: GormPublic{}, Join: make([]*GormJoin, 0)}
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("non-struct input panicked: %v", r)
		}
	}()
	ResolveSearchQuery(42, c)
	ResolveSearchQuery("str", c)
	ResolveSearchQuery(nil, c)
}

func TestResolveSearchQueryIsNullNonNilable(t *testing.T) {
	c := &GormCondition{GormPublic: GormPublic{}, Join: make([]*GormJoin, 0)}
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("isnull on non-nilable type panicked: %v", r)
		}
	}()
	ResolveSearchQuery(nullableFields{Int: 1, Str: "s"}, c)
	ResolveSearchQuery(nullableFields{Int: 1, Str: "s", Slice: []string{"a"}}, c)
	v := 1
	ResolveSearchQuery(nullableFields{Ptr: &v, Int: 1}, c)
	if len(c.Where) == 0 {
		t.Fatal("expected isnull where conditions")
	}
}

func TestResolveSearchQueryIsZeroValueSkipped(t *testing.T) {
	c := &GormCondition{GormPublic: GormPublic{}, Join: make([]*GormJoin, 0)}
	ResolveSearchQuery(nullableFields{}, c)
	if len(c.Where) != 0 {
		t.Fatalf("zero-value fields should be skipped, got %d conditions", len(c.Where))
	}
}

func TestReflectIsZeroSemantics(t *testing.T) {
	var p *int
	if !reflect.ValueOf(p).IsNil() {
		t.Fatal("nil pointer IsNil should be true")
	}
}
