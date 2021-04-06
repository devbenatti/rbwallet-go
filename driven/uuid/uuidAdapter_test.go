package uuid

import (
	"testing"
	"unicode"
)

func TestGenerate(t *testing.T) {
	ua := UuidAdapter{}
	uuid := ua.Generate()

	if !IsUpper(uuid.String()) {
		t.Error("expect UUID should be in Upper")
	}

	if len(uuid.String()) != 36 {
		t.Errorf("expect UUID should have 36 characteres, found '%d'", len(uuid.String()))
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
