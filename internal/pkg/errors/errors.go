package errors

import "errors"

var (
	Is = errors.Is
)

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidTemplate = errors.New("template is invalid")
)
