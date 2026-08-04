package runtime

import (
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
	"testing"
)

func TestGetDbByKeyExactFirst(t *testing.T) {
	app := NewConfig()
	app.SetDb("*", &gorm.DB{})
	app.SetDb("admin.example.com", &gorm.DB{})

	if app.GetDbByKey("admin.example.com") == app.GetDbByKey("*") {
		t.Fatal("exact host must win over wildcard key")
	}
	if app.GetDbByKey("other.example.com") != app.GetDbByKey("*") {
		t.Fatal("unmatched host must fall back to wildcard key")
	}
}

func TestGetDbByKeyMissing(t *testing.T) {
	app := NewConfig()
	app.SetDb("a.example.com", &gorm.DB{})
	if app.GetDbByKey("a.example.com") == nil {
		t.Fatal("exact key must be found")
	}
	if app.GetDbByKey("b.example.com") != nil {
		t.Fatal("missing key must return nil (no wildcard)")
	}
}

func TestGetCasbinKeyExactFirst(t *testing.T) {
	app := NewConfig()
	app.SetCasbin("*", &casbin.SyncedEnforcer{})
	app.SetCasbin("admin.example.com", &casbin.SyncedEnforcer{})

	if app.GetCasbinKey("admin.example.com") == app.GetCasbinKey("*") {
		t.Fatal("exact host must win over wildcard key")
	}
	if app.GetCasbinKey("other.example.com") != app.GetCasbinKey("*") {
		t.Fatal("unmatched host must fall back to wildcard key")
	}
	if app.GetCasbinKey("missing.example.com") == nil {
		t.Fatal("missing host with wildcard must not return nil")
	}
}

func TestGetCasbinKeyMissing(t *testing.T) {
	app := NewConfig()
	app.SetCasbin("a.example.com", &casbin.SyncedEnforcer{})
	if app.GetCasbinKey("a.example.com") == nil {
		t.Fatal("exact key must be found")
	}
	if app.GetCasbinKey("b.example.com") != nil {
		t.Fatal("missing key must return nil (no wildcard)")
	}
}
