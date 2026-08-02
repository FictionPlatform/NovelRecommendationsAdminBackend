package database

import "testing"

func TestMaskDSN(t *testing.T) {
	cases := map[string]string{
		"host=127.0.0.1 port=5432 user=postgres dbname=public password=123456 sslmode=disable": "host=127.0.0.1 port=5432 user=postgres dbname=public password=*** sslmode=disable",
		"postgres://user:secret@127.0.0.1:5432/db?sslmode=disable":                            "postgres://user:***@127.0.0.1:5432/db?sslmode=disable",
	}
	for in, want := range cases {
		if got := maskDSN(in); got != want {
			t.Errorf("maskDSN(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestExtractDSNPassword(t *testing.T) {
	if got := extractDSNPassword("host=x password=123456"); got != "123456" {
		t.Errorf("kv extract = %q", got)
	}
	if got := extractDSNPassword("postgres://u:abc@h/db"); got != "abc" {
		t.Errorf("url extract = %q", got)
	}
	if got := extractDSNPassword("host=x"); got != "" {
		t.Errorf("empty extract = %q", got)
	}
}

func TestValidateDSNPassword(t *testing.T) {
	weak := []string{"host=x password=123456", "host=x password=root", "host=x password=short1", "host=x", "postgres://u:qwerty@h/db"}
	for _, s := range weak {
		if err := validateDSNPassword(s); err == nil {
			t.Errorf("expected weak rejected: %q", s)
		}
	}
	if err := validateDSNPassword("host=x password=MyStr0ngP@ss!"); err != nil {
		t.Errorf("expected strong accepted: %v", err)
	}
}

func TestResolveEnv(t *testing.T) {
	t.Setenv("DB_PASSWORD", "secret123")
	if got := resolveEnv("host=x password=${DB_PASSWORD}"); got != "host=x password=secret123" {
		t.Errorf("resolveEnv = %q", got)
	}
	if got := resolveEnv("host=x password=$DB_PASSWORD"); got != "host=x password=$DB_PASSWORD" {
		t.Errorf("bare $ must not expand: %q", got)
	}
}
