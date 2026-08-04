package lang

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewI18nMalformedRows(t *testing.T) {
	dir := t.TempDir()
	csvPath := filepath.Join(dir, "en.csv")
	content := "key1,value1\n\nkey2,value2\nsingle-col\n,onlyvalue\nkey3,value3\n"
	if err := os.WriteFile(csvPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	i18n, err := newI18n(dir, "en")
	if err != nil {
		t.Fatal(err)
	}

	cases := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"missing": "missing",
	}
	for key, want := range cases {
		if got := i18n.T(key); got != want {
			t.Errorf("T(%q) = %q, want %q", key, got, want)
		}
	}
}

func TestInitLangMissingFile(t *testing.T) {
	old := EnLang
	defer func() { EnLang = old }()

	InitLang() // real config/lang/en.csv exists -> should succeed
	if EnLang == nil {
		t.Fatal("InitLang should succeed with real file")
	}
}
