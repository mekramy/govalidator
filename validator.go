package govalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/mekramy/goi18n"
)

// Validator defines an interface for localized validation functionality.
type Validator interface {
	// AddValidation registers a custom validation rule with a custom validation function.
	// Parameters:
	//   rule: The name of the validation rule.
	//   f: The validation function to be applied.
	AddValidation(rule string, f validator.Func)

	// AddTranslation adds a translation message for a validation rule in a specified locale.
	// Parameters:
	//   locale: The locale for the translation.
	//   rule: The validation rule to associate with the translation.
	//   message: The translated message.
	//   options: Optional pluralization options.
	AddTranslation(locale, rule, message string, options ...goi18n.PluralOption)

	// Struct validates an entire struct based on its defined validation rules.
	// Returns validation errors in the provided locale, defaulting to the translator locale if empty.
	// Parameters:
	//   locale: The locale for error messages.
	//   value: The struct to validate.
	// Returns:
	//   ValidationError: The validation errors for the struct.
	Struct(locale string, value any) ValidationError

	// StructExpect validates a struct while ignoring specified fields.
	// Returns validation errors in the provided locale, defaulting to the translator locale if empty.
	// Parameters:
	//   locale: The locale for error messages.
	//   value: The struct to validate.
	//   fields: The list of fields to ignore.
	// Returns:
	//   ValidationError: The validation errors for the struct, excluding ignored fields.
	StructExpect(locale string, value any, fields ...string) ValidationError

	// StructPartial validates only specified fields of a struct.
	// Returns validation errors in the provided locale, defaulting to the translator locale if empty.
	// Parameters:
	//   locale: The locale for error messages.
	//   value: The struct to validate.
	//   fields: The list of fields to validate.
	// Returns:
	//   ValidationError: The validation errors for the specified fields.
	StructPartial(locale string, value any, fields ...string) ValidationError

	// Var validates a single variable against a rule with optional custom messages.
	// Parameters:
	//   locale: The locale for error messages.
	//   name: The name of the field being validated.
	//   value: The value to validate.
	//   rule: The validation rule to apply.
	// Returns:
	//   ValidationError: The validation errors, if any.
	Var(locale, name string, value any, rules string) ValidationError

	// VarWithValue validates a variable against another using a rule with custom error messages.
	// Parameters:
	//   locale: The locale for error messages.
	//   name: The name of the field being validated.
	//   value: The value to validate.
	//   other: The other value to compare against.
	//   rule: The validation rule to apply.
	// Returns:
	//   ValidationError: The validation errors, if any.
	VarWithValue(locale, name string, value any, other any, rules string) ValidationError
}

// NewValidator creates a new Validator instance with optional configurations.
// Initializes the I18nValidator with the provided translator and base validator, and applies any options.
func NewValidator(validator *validator.Validate, options ...Options) Validator {
	// Initialize the I18nValidator with the provided validator
	v := &I18nValidator{
		translator: nil,
		validator:  validator,
	}

	// Apply any additional options
	for _, opt := range options {
		opt(v)
	}

	// Return the configured validator instance
	return v
}
