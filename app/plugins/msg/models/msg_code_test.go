package models

import (
	"strings"
	"testing"
)

func TestMsgCodeHashCode(t *testing.T) {
	m := &MsgCode{Code: "123456"}
	if err := m.HashCode(); err != nil {
		t.Fatal(err)
	}
	if m.Code == "123456" {
		t.Fatal("code should not be stored in plaintext")
	}
	if !strings.HasPrefix(m.Code, "$2a$") && !strings.HasPrefix(m.Code, "$2b$") && !strings.HasPrefix(m.Code, "$2y$") {
		t.Fatalf("code should be bcrypt hash, got %q", m.Code)
	}
	if !m.CheckCode("123456") {
		t.Fatal("CheckCode should pass for original code")
	}
	if m.CheckCode("654321") {
		t.Fatal("CheckCode should fail for wrong code")
	}
}

func TestMsgCodeHashCodeIdempotent(t *testing.T) {
	m := &MsgCode{Code: "888888"}
	if err := m.HashCode(); err != nil {
		t.Fatal(err)
	}
	first := m.Code
	if err := m.HashCode(); err != nil {
		t.Fatal(err)
	}
	if m.Code != first {
		t.Fatal("HashCode should not double-hash an already hashed value (BeforeUpdate path)")
	}
}

func TestMsgCodeEmptyCode(t *testing.T) {
	m := &MsgCode{}
	if err := m.HashCode(); err != nil {
		t.Fatal(err)
	}
	if m.CheckCode("") || m.CheckCode("123") {
		t.Fatal("empty code should never pass")
	}
}
