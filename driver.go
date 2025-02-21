package govalidator

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mekramy/goi18n"
)

type I18nValidator struct {
	prefix     string
	translator goi18n.Translator
	validator  *validator.Validate
}

func (v *I18nValidator) AddValidation(rule string, f validator.Func) {
	rule = strings.TrimSpace(rule)
	if rule == "" {
		return
	}

	v.validator.RegisterValidation(rule, f)
}

func (v *I18nValidator) AddTranslation(locale, rule, message string, options ...goi18n.PluralOption) {
	rule = strings.TrimSpace(rule)
	if rule == "" || v.translator == nil {
		return
	}

	if v.prefix == "" {
		v.translator.AddMessage(locale, rule, message, options...)
	} else {
		v.translator.AddMessage(locale, v.prefix+"."+rule, message, options...)
	}
}

func (v *I18nValidator) Struct(locale string, value any) ValidationError {
	return v.parseStructErrors(
		locale,
		value,
		v.validator.Struct(value),
	)
}

func (v *I18nValidator) StructExpect(locale string, value any, fields ...string) ValidationError {
	return v.parseStructErrors(
		locale,
		value,
		v.validator.StructExcept(value, fields...),
	)
}

func (v *I18nValidator) StructPartial(locale string, value any, fields ...string) ValidationError {
	return v.parseStructErrors(
		locale,
		value,
		v.validator.StructPartial(value, fields...),
	)
}

func (v *I18nValidator) Var(locale, name string, value any, rules string) ValidationError {
	return v.parseVariableErrors(
		locale,
		name,
		value,
		v.validator.Var(value, rules),
	)
}

func (v *I18nValidator) VarWithValue(locale, name string, value any, other any, rules string) ValidationError {
	return v.parseVariableErrors(
		locale,
		name,
		value,
		v.validator.VarWithValue(value, other, rules),
	)
}

// translate generates a localized error message based on the provided value, field, and parameters.
func (v *I18nValidator) translate(locale, name, rule, field string, param, value any, count int) string {
	// Return empty string if translator not passed to I18nValidator
	if v.translator == nil {
		return ""
	}

	// Try resolving error translation using the Translatable interface
	if t, ok := value.(Translatable); ok {
		if res := t.TranslateError(locale, rule, field); res != "" {
			return res
		}
	}

	// Use the main translator to generate the message with pluralization support
	// If a prefix is set, prepend it to the rule
	if v.prefix != "" {
		rule = v.prefix + "." + rule
	}

	// Next, attempt to translate the field name using TranslatableField interface
	if t, ok := value.(TranslatableField); ok {
		if n := t.TranslateTitle(locale, field); n != "" {
			name = n
		}
	}

	return v.translator.Plural(locale, rule, count, map[string]any{
		"field": name,
		"param": param,
	})
}

// parseStructErrors processes and translates validation errors
// based on the provided locale and value for struct.
func (v *I18nValidator) parseStructErrors(locale string, value any, err error) ValidationError {
	// Skip nil error
	if err == nil {
		return NewEmptyError()
	}

	// Assert the error as validator.ValidationErrors
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return NewError(err)
	}

	// Initialize the result validation error
	res := NewEmptyError()

	// Iterate over each validation error and process
	for _, field := range errs {
		// Parse the field parameter and handle its type (int or float)
		var count int
		var param any = field.Param()
		if i, f := parseNumeric(field.Param()); i != nil {
			param = *i
			count = int(*i)
		} else if f != nil {
			param = *f
			count = int(*f)
		}

		// Add the translated error if translator available or raw error to the result
		if v.translator == nil {
			res.AddError(field.Field(), field.Tag(), field.Error())
		} else {
			res.AddError(
				field.Field(),
				field.Tag(),
				v.translate(
					locale, field.Field(), field.Tag(),
					field.StructField(), param, value, count,
				),
			)
		}

	}

	// Return the aggregated validation errors
	return res
}

// parseVariableErrors processes and translates validation errors based on the provided locale and value for variable.
func (v *I18nValidator) parseVariableErrors(locale, name string, value any, err error) ValidationError {
	// Skip nil error
	if err == nil {
		return NewEmptyError()
	}

	// Assert the error as validator.ValidationErrors
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return NewError(err)
	}

	// Initialize the result validation error
	res := NewEmptyError()

	// Iterate over each validation error and process
	for _, field := range errs {
		// Parse the field parameter and handle its type (int or float)
		var count int
		var param any = field.Param()
		if i, f := parseNumeric(field.Param()); i != nil {
			param = *i
			count = int(*i)
		} else if f != nil {
			param = *f
			count = int(*f)
		}

		// Add the translated error if translator available or raw error to the result
		if v.translator == nil {
			res.AddError(field.Field(), field.Tag(), field.Error())
		} else {
			res.AddError(
				name,
				field.Tag(),
				v.translate(
					locale, name, field.Tag(), name,
					param, value, count,
				),
			)
		}
	}

	// Return the aggregated validation errors
	return res
}
