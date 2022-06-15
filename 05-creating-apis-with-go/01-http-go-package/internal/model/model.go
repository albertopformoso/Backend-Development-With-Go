package model

import "errors"

var (
	// ErrPersonCanNotBeNil person can't be null
	ErrPersonCanNotBeNil = errors.New("person can't be null")
	// ErrIDPersonDoesNotExists person doesn't exist
	ErrIDPersonDoesNotExists = errors.New("person doesn't exist")
)
