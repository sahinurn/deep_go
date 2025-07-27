package main

import (
	"fmt"
	"strings"
)

type MultiError struct {
	errors []error
}

func (e *MultiError) Error() string {
	if e == nil || len(e.errors) == 0 {
		return ""
	}

	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("%d errors occured:\n", len(e.errors)))

	for _, err := range e.errors {
		_, _ = sb.WriteString("\t* ")
		_, _ = sb.WriteString(err.Error())
	}

	_, _ = sb.WriteString("\n")

	return sb.String()
}

func Append(err error, errs ...error) *MultiError {
	mr := &MultiError{}
	if e, ok := err.(*MultiError); ok {
		mr = e
	}

	mr.errors = append(mr.errors, errs...)

	return mr
}
