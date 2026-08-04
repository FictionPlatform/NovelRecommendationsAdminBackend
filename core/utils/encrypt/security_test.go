package encrypt

import (
	"strings"
	"testing"
)

func TestGenerateRandomKey(t *testing.T) {
	for _, tc := range []struct {
		name string
		fn   func() string
		size int
	}{
		{"key20", GenerateRandomKey20, 20},
		{"key16", GenerateRandomKey16, 16},
		{"key6", GenerateRandomKey6, 6},
	} {
		seen := make(map[string]struct{})
		for i := 0; i < 200; i++ {
			k := tc.fn()
			if len(k) != tc.size {
				t.Fatalf("%s length = %d, want %d", tc.name, len(k), tc.size)
			}
			for _, r := range k {
				if !strings.ContainsRune(symbol, r) {
					t.Fatalf("%s contains char %q outside charset", tc.name, r)
				}
			}
			if _, dup := seen[k]; dup {
				t.Fatalf("%s duplicate (200 samples)", tc.name)
			}
			seen[k] = struct{}{}
		}
	}
}
