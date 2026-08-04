package strutils

import (
	"regexp"
	"testing"
)

var digit6 = regexp.MustCompile(`^\d{6}$`)

func TestGenerateValidateCode(t *testing.T) {
	for i := 0; i < 1000; i++ {
		code := GenerateValidateCode()
		if !digit6.MatchString(code) {
			t.Fatalf("code %q is not 6 digits", code)
		}
	}
}

func TestGenerateValidateCodeDistinct(t *testing.T) {
	seen := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		code := GenerateValidateCode()
		if _, dup := seen[code]; dup {
			t.Fatalf("duplicate code %q (100 samples, 10^6 space)", code)
		}
		seen[code] = struct{}{}
	}
}
