package pkg

import "fmt"

// NoSuchContextError is returned when a  context cannot be found.
type NoSuchContextError struct {
	Message string
}

func (e NoSuchContextError) Error() string {
	return e.Message
}

// NewNoSuchContextError creates a new `NoSuchContextError`.
func NewNoSuchContextError(name string) NoSuchContextError {
	message := fmt.Sprintf("context not found for specified name: %s", name)
	return NoSuchContextError{Message: message}
}

// CannotSetContextError is a custom error type.
type CannotSetContextError struct {
	Message string
}

func (e CannotSetContextError) Error() string {
	return e.Message
}

// NewCannotSetContextError creates a new `CannotSetContextError`.
func NewCannotSetContextError(name string) CannotSetContextError {
	message := fmt.Sprintf("no context exists with the name: %s", name)
	return CannotSetContextError{Message: message}
}
