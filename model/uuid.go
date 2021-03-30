package model

import "strings"

type Uuid struct {
	val string
}

func (u *Uuid) Val() string {
	return u.val
}

func NewUuid(val string) Uuid {
	return Uuid{val: strings.ToUpper(val)}
}
