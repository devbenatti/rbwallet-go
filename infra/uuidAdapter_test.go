package infra

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

}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
