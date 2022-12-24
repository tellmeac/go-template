package store

import "fmt"

// ErrNotFound indicates that a resource was not found.
type ErrNotFound struct {
	resource string
	ID       string
	wrapped  error
}

func NewErrNotFound(resource, id string) *ErrNotFound {
	return &ErrNotFound{
		resource: resource,
		ID:       id,
	}
}

func (e *ErrNotFound) Wrap(err error) *ErrNotFound {
	e.wrapped = err
	return e
}

func (e *ErrNotFound) Error() string {
	if e.wrapped != nil {
		return fmt.Sprintf("resource: %s id: %s error: %s", e.resource, e.ID, e.wrapped)
	}

	return fmt.Sprintf("resource: %s id: %s", e.resource, e.ID)
}

// ErrConflict indicates a conflict that occurred.
type ErrConflict struct {
	Resource string // The resource which created the conflict.
	err      error  // Internal error.
	meta     string // Any additional metadata.
}

func NewErrConflict(resource string, err error, meta string) *ErrConflict {
	return &ErrConflict{
		Resource: resource,
		err:      err,
		meta:     meta,
	}
}

func (e *ErrConflict) Error() string {
	msg := e.Resource + "exists " + e.meta
	if e.err != nil {
		msg += " " + e.err.Error()
	}
	return msg
}

func (e *ErrConflict) Unwrap() error {
	return e.err
}
