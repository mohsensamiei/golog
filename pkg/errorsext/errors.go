package errorsext

import "github.com/pkg/errors"

type StackError struct {
	Message    string
	StackTrace string `yaml:"stack_trace,omitempty" json:"stack_trace,omitempty"`
}
type StackTracer interface {
	error
	StackTrace() errors.StackTrace
}
