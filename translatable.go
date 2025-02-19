package govalidator

// Translatable defines an interface for translating validation error messages.
type Translatable interface {
	// TranslateError returns a localized error message for a given rule and field.
	TranslateError(locale, rule, field string) string
}

// TranslatableField defines an interface for translating field names.
type TranslatableField interface {
	// TranslateTitle returns a localized display name for a given field.
	TranslateTitle(locale, field string) string
}
