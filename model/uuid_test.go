package model

import "testing"

func TestNewUuid(t *testing.T) {
	uuid := NewUuid("fb6eb1ca-6f55-414f-ab78-f62d227b6ba5")
	expected := "FB6EB1CA-6F55-414F-AB78-F62D227B6BA5"

	if uuid.Val() != expected {
		t.Errorf("expect '%s', result '%s'", expected, uuid.Val())
	}
}
