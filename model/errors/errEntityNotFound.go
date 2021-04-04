package model

import "errors"

type ErrEntityNotFound struct {
	err error
}

func NewErrEntityNotFound(msg string) *ErrEntityNotFound {
	return &ErrEntityNotFound{
		err: errors.New(msg),
	}
}

func (errAc *ErrEntityNotFound) Error() string {
	return errAc.err.Error()
}
