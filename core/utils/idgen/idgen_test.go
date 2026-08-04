package idgen

import (
	"strings"
	"testing"
)

func TestInviteId(t *testing.T) {
	seen := make(map[string]struct{})
	for i := 0; i < 200; i++ {
		id := InviteId()
		if len(id) != 6 {
			t.Fatalf("invite id %q length = %d, want 6", id, len(id))
		}
		for _, r := range id {
			if !strings.ContainsRune(inviteChars, r) {
				t.Fatalf("invite id %q contains char %q outside charset", id, r)
			}
		}
		if _, dup := seen[id]; dup {
			t.Fatalf("duplicate invite id %q (200 samples, 62^6 space)", id)
		}
		seen[id] = struct{}{}
	}
}
