package govalidator

import (
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mekramy/goi18n"
	"github.com/mekramy/gojalaali"
	"github.com/mekramy/govalidator/funcs"
)

// Options defines a function type that modifies I18nValidator.
type Options func(*I18nValidator)

// WithTranslator configures the I18nValidator with a translator and a translation key prefix.
func WithTranslator(translator goi18n.Translator, prefix string) Options {
	prefix = strings.TrimSpace(prefix)
	return func(iv *I18nValidator) {
		iv.translator = translator
		iv.prefix = prefix
	}
}

// WithFiberTagResolver returns an option for configuring the I18nValidator to resolve field names from various tags.
// It checks the "field", "json", "form", and "xml" tags in priority order. If no valid tag is found, it defaults to the field name.
// Fields with the "-" tag are ignored.
func WithFiberTagResolver() Options {
	return func(iv *I18nValidator) {
		// Register a function that resolves field names from tags
		iv.validator.RegisterTagNameFunc(func(field reflect.StructField) string {
			var name string

			// Check the "field", "json", "form", and "xml" tags in order
			if n := strings.SplitN(field.Tag.Get("field"), ",", 2)[0]; n != "" {
				name = n
			} else if n := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]; n != "" {
				name = n
			} else if n := strings.SplitN(field.Tag.Get("form"), ",", 2)[0]; n != "" {
				name = n
			} else if n := strings.SplitN(field.Tag.Get("xml"), ",", 2)[0]; n != "" {
				name = n
			}

			// Ignore fields with "-" tag
			if name == "-" {
				return ""
			}

			// Return the resolved name or default to field name
			if name != "" {
				return name
			}
			return field.Name
		})
	}
}

// WithUsernameValidator adds a validation rule for usernames (letters, numbers, underscores).
func WithUsernameValidator(messages map[string]string, rule ...string) Options {
	tag := resolveParams("username", rule...)
	messages = resolveMessages(
		messages,
		"Only letters, numbers, and underscores are allowed",
	)

	return func(iv *I18nValidator) {
		iv.AddValidation(tag, func(fl validator.FieldLevel) bool {
			return funcs.IsValidUsername(fl.Field().String())
		})
		for l, m := range messages {
			iv.AddTranslation(l, tag, m)
		}
	}
}

// WithAlphaNumericValidator adds validation for English letters and numbers.
func WithAlphaNumericValidator(messages map[string]string, rule ...string) Options {
	tag := resolveParams("alnum", rule...)
	messages = resolveMessages(
		messages,
		"Only english letters and numbers are allowed",
	)
	return func(iv *I18nValidator) {
		iv.AddValidation(tag, func(fl validator.FieldLevel) bool {
			return funcs.IsAlphaNumeric(
				fl.Field().String(),
				toChars(fl.Param())...,
			)
		})
		for l, m := range messages {
			iv.AddTranslation(l, tag, m)
		}
	}
}

// WithAlphaNumericPersianValidator adds validation for English, Persian letters, and numbers.
func WithAlphaNumericPersianValidator(messages map[string]string, rule ...string) Options {
	tag := resolveParams("alnum_fa", rule...)
	messages = resolveMessages(
		messages,
		"Only english letters, persian letters, and numbers are allowed",
	)

	return func(iv *I18nValidator) {
		iv.AddValidation(tag, func(fl validator.FieldLevel) bool {
			return funcs.IsAlphaNumericWithPersian(
				fl.Field().String(),
				toChars(fl.Param())...,
			)
		})
		for l, m := range messages {
			iv.AddTranslation(l, tag, m)
		}
	}
}

// WithIranianPhoneValidator adds validation for 11-digit Iranian phone numbers.
func WithIranianPhoneValidator(messages map[string]string, rule ...string) Options {
	tag := resolveParams("phone", rule...)
	messages = resolveMessages(
		messages,
		"Must be a valid 11-digit iranian phone number",
	)

	return func(iv *I18nValidator) {
		iv.AddValidation(tag, func(fl validator.FieldLevel) bool {
			return funcs.IsValidIranianPhone(fl.Field().String())
		})
		for l, m := range messages {
			iv.AddTranslation(l, tag, m)
		}
	}
}

// WithIranianMobileValidator adds validation for 11-digit Iranian mobile numbers.
func WithIranianMobileValidator(messages map[string]string, rule ...string) Options {
	tag := resolveParams("mobile", rule...)
	messages = resolveMessages(
		messages,
		"Must be a valid 11-digit iranian mobile number",
	)

	return func(iv *I18nValidator) {
		iv.AddValidation(tag, func(fl validator.FieldLevel) bool {
			return funcs.IsValidIranianMobile(fl.Field().String())
		})
		for l, m := range messages {
			iv.AddTranslation(l, tag, m)
		}
	}
}

// WithIranianPostalCodeValidator adds validation for 10-digit Iranian postal codes.
func WithIranianPostalCodeValidator(messages map[string]string, rule ...string) Options {
	tag := resolveParams("postal_code", rule...)
	messages = resolveMessages(
		messages,
		"Must be a valid 10-digit iranian postal code",
	)

	return func(iv *I18nValidator) {
		iv.AddValidation(tag, func(fl validator.FieldLevel) bool {
			return funcs.IsValidIranianPostalCode(fl.Field().String())
		})
		for l, m := range messages {
			iv.AddTranslation(l, tag, m)
		}
	}
}

// WithIranianIdNumberValidator adds validation for Iranian birth certificate numbers.
func WithIranianIdNumberValidator(messages map[string]string, rule ...string) Options {
	tag := resolveParams("id_number", rule...)
	messages = resolveMessages(
		messages,
		"Must be a valid iranian birth certificate number",
	)

	return func(iv *I18nValidator) {
		iv.AddValidation(tag, func(fl validator.FieldLevel) bool {
			return funcs.IsValidIranianIdNumber(fl.Field().String())
		})
		for l, m := range messages {
			iv.AddTranslation(l, tag, m)
		}
	}
}

// WithIranianNationalCodeValidator adds validation for 10-digit Iranian national ID numbers.
func WithIranianNationalCodeValidator(messages map[string]string, rule ...string) Options {
	tag := resolveParams("national_code", rule...)
	messages = resolveMessages(
		messages,
		"Must be a valid 10 digit iranian national id number",
	)

	return func(iv *I18nValidator) {
		iv.AddValidation(tag, func(fl validator.FieldLevel) bool {
			return funcs.IsValidIranianNationalCode(fl.Field().String())
		})
		for l, m := range messages {
			iv.AddTranslation(l, tag, m)
		}
	}
}

// WithIranianCreditNumberValidator adds validation for 16-digit Iranian credit card numbers.
func WithIranianCreditNumberValidator(messages map[string]string, rule ...string) Options {
	tag := resolveParams("credit_number", rule...)
	messages = resolveMessages(
		messages,
		"Must be a valid 16 digit iranian credit card number",
	)

	return func(iv *I18nValidator) {
		iv.AddValidation(tag, func(fl validator.FieldLevel) bool {
			return funcs.IsValidIranianBankCard(fl.Field().String())
		})
		for l, m := range messages {
			iv.AddTranslation(l, tag, m)
		}
	}
}

// WithIranianIBANValidator adds validation for 24-digit Iranian IBAN numbers.
func WithIranianIBANValidator(messages map[string]string, rule ...string) Options {
	tag := resolveParams("iban", rule...)
	messages = resolveMessages(
		messages,
		"Must be a valid 24 digit iranian IBAN number",
	)

	return func(iv *I18nValidator) {
		iv.AddValidation(tag, func(fl validator.FieldLevel) bool {
			return funcs.IsValidIranianIBAN(fl.Field().String())
		})
		for l, m := range messages {
			iv.AddTranslation(l, tag, m)
		}
	}
}

// WithJalaaliValidator adds validation for Jalaali datetime strings.
func WithJalaaliValidator(messages map[string]string, rule ...string) Options {
	tag := resolveParams("jalaali", rule...)
	messages = resolveMessages(
		messages,
		"Must be a valid jalaali datetime",
	)

	return func(iv *I18nValidator) {
		iv.AddValidation(tag, func(fl validator.FieldLevel) bool {
			layout := fl.Param()
			if layout == "" {
				layout = time.RFC3339
			}
			d, err := gojalaali.Parse(layout, fl.Field().String())
			return err == nil && !d.IsZero()
		})
		for l, m := range messages {
			iv.AddTranslation(l, tag, m)
		}
	}
}
