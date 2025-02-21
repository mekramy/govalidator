package govalidator

import (
	"encoding/json"
	"strings"
)

// ValidationError defines an interface for managing validation errors.
type ValidationError interface {
	// HasError checks if there is any validation or internal error.
	HasError() bool

	// HasInternalError checks if there is a internal system error.
	HasInternalError() bool

	// HasValidationErrors returns true if there are validation errors.
	HasValidationErrors() bool

	// IsFailed checks if a specific field has validation errors.
	IsFailed(field string) bool

	// IsFailedOn checks if a specific field failed on a given validation rule.
	IsFailedOn(field, rule string) bool

	// InternalError returns the internal system error related to the validation process, if any.
	InternalError() error

	// Errors returns a nested map of validation errors for each field and rule.
	Errors() map[string]map[string]string

	// Messages returns a map of validation error messages for each field.
	Messages() map[string][]string

	// Rules returns a map of validation error rules for each field.
	Rules() map[string][]string

	// MarshalJSON serializes the validation errors into JSON format.
	MarshalJSON() ([]byte, error)

	// String returns a string representation of all validation errors.
	String() string

	// AddError records a validation error for a specific field and validation rule.
	AddError(field, rule string, message ...string)
}

// NewError creates a new ValidationError with an internal error.
func NewError(err error) ValidationError {
	return &vErrors{
		interr: err,
		valerr: make(map[string]map[string]string),
	}
}

// NewEmptyError creates a new empty ValidationError without any internal error.
func NewEmptyError() ValidationError {
	return &vErrors{
		interr: nil,
		valerr: make(map[string]map[string]string),
	}
}

// vError handles validation errors and implements the ValidationError interface.
type vErrors struct {
	interr error
	valerr map[string]map[string]string
}

func (e *vErrors) HasError() bool {
	return e.interr != nil || len(e.valerr) > 0
}

func (e *vErrors) HasInternalError() bool {
	return e.interr != nil
}

func (e *vErrors) HasValidationErrors() bool {
	return len(e.valerr) > 0
}

func (e *vErrors) IsFailed(field string) bool {
	_, exists := e.valerr[field]
	return exists
}

func (e *vErrors) IsFailedOn(field, rule string) bool {
	_, exists := e.valerr[field][rule]
	return exists
}

func (e *vErrors) InternalError() error {
	return e.interr
}

func (e *vErrors) Errors() map[string]map[string]string {
	return e.valerr
}

func (e *vErrors) Messages() map[string][]string {
	messages := make(map[string][]string)
	for field, errs := range e.valerr {
		messages[field] = make([]string, 0)
		for _, message := range errs {
			messages[field] = append(messages[field], message)
		}
	}
	return messages
}

func (e *vErrors) Rules() map[string][]string {
	rules := make(map[string][]string)
	for field, errs := range e.valerr {
		rules[field] = make([]string, 0)
		for rule := range errs {
			rules[field] = append(rules[field], rule)
		}
	}
	return rules
}

func (e *vErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.valerr)
}

func (e *vErrors) String() string {
	var builder strings.Builder
	for field, errs := range e.valerr {
		builder.WriteString(field + ":\n")
		for rule, message := range errs {
			builder.WriteString("    " + rule + ": " + message + "\n")
		}
	}
	return builder.String()
}

func (e *vErrors) AddError(field, rule string, message ...string) {
	msg := resolveParams("", message...)
	_, exists := e.valerr[field]
	if exists {
		e.valerr[field][rule] = msg
	} else {
		e.valerr[field] = map[string]string{rule: msg}
	}

}
