package uuid

import (
	"testing"
	"unicode"
)

func TestGenerate(t *testing.T) {
	ua := UuidAdapter{}
	uuid := ua.Generate()

	if !IsUpper(uuid.Val()) {
		t.Error("expect UUID should be in Upper")
	}

	if len(uuid.Val()) != 36 {
		t.Errorf("expect UUID should have 36 characteres, found '%d'", len(uuid.Val()))
	}

}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
