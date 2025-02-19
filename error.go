package govalidator

import (
	"encoding/json"
	"strings"
)

// ValidationError represents a contract for handling validation errors.
type ValidationError interface {
	// AddError records a validation error for a specific field and rule.
	AddError(field, rule string, message ...string)

	// HasError checks if there are any validation errors.
	HasError() bool

	// Failed checks if a specific field has any validation errors.
	Failed(field string) bool

	// FailedOn checks if a specific field failed on a particular validation rule.
	FailedOn(field, rule string) bool

	// Rules returns a map of validation rules applied to each field.
	Rules() map[string][]string

	// Messages returns a map of error messages for each field.
	Messages() map[string][]string

	// Errors returns a nested map containing validation errors for each field and rule.
	Errors() map[string]map[string]string

	// MarshalJSON converts validation errors into a JSON format.
	MarshalJSON() ([]byte, error)

	// String returns a string representation of the validation errors.
	String() string
}

// NewError creates and returns a new instance of ValidationError.
func NewError() ValidationError {
	return &vErrors{
		err: make(map[string]map[string]string),
	}
}

type vErrors struct {
	err map[string]map[string]string
}

func (e *vErrors) AddError(field, rule string, message ...string) {
	m := ""
	if len(message) > 0 {
		m = message[0]
	}

	_, exists := e.err[field]
	if exists {
		e.err[field][rule] = m
	} else {
		e.err[field] = map[string]string{rule: m}
	}
}

func (e *vErrors) HasError() bool {
	return len(e.err) > 0
}

func (e *vErrors) Failed(field string) bool {
	_, exists := e.err[field]
	return exists
}

func (e *vErrors) FailedOn(field, rule string) bool {
	_, exists := e.err[field][rule]
	return exists
}

func (e *vErrors) Rules() map[string][]string {
	rules := make(map[string][]string)
	for field, errs := range e.err {
		rules[field] = make([]string, 0)
		for rule := range errs {
			rules[field] = append(rules[field], rule)
		}
	}
	return rules
}

func (e *vErrors) Messages() map[string][]string {
	messages := make(map[string][]string)
	for field, errs := range e.err {
		messages[field] = make([]string, 0)
		for _, message := range errs {
			messages[field] = append(messages[field], message)
		}
	}
	return messages
}

func (e *vErrors) Errors() map[string]map[string]string {
	return e.err
}

func (e *vErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.err)
}

func (e *vErrors) String() string {
	var builder strings.Builder
	for field, errs := range e.err {
		builder.WriteString(field + ":\n")
		for rule, message := range errs {
			builder.WriteString("    " + rule + ": " + message + "\n")
		}
	}
	return builder.String()
}
