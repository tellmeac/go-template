package model

import (
	"encoding/json"
	"fmt"
)

type AppError struct {
	message string
	err     error
	code    int
}

func NewError(msg string) *AppError {
	return &AppError{
		message: msg,
	}
}

func (a *AppError) WithCode(code int) *AppError {
	a.code = code
	return a
}

func (a *AppError) Wrap(err error) *AppError {
	a.err = err
	return a
}

func (a *AppError) Error() string {
	return fmt.Sprintf("%s: %s", a.message, a.err)
}

func (a *AppError) String() string {
	return fmt.Sprintf("%s: %s", a.message, a.err)
}

func (a *AppError) StatusCode() int {
	return a.code
}

func (a *AppError) MarshalJSON() ([]byte, error) {
	type errBody struct {
		Error string `json:"error"`
	}

	return json.Marshal(errBody{Error: a.String()})
}
