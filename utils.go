package govalidator

import (
	"reflect"
	"strconv"
	"strings"
)

// toChars converts a string into a slice of single-character strings.
// It correctly handles Unicode characters, including Persian and emojis.
func toChars(s string) []string {
	var result []string
	for _, r := range s {
		result = append(result, string(r))
	}
	return result
}

// resolveParams returns the first non-empty, trimmed string from provided values.
// Falls back to 'fallback' if all values are empty.
func resolveParams(fallback string, vals ...string) string {
	if len(vals) > 0 {
		if strings.TrimSpace(vals[0]) != "" {
			return strings.TrimSpace(vals[0])
		}
	}
	return fallback
}

// resolveMessages returns the provided messages map or a default message if nil.
func resolveMessages(messages map[string]string, def string) map[string]string {
	if messages == nil {
		return map[string]string{"": def}
	}
	return messages
}

// parseNumeric attempts to parse a string into an integer or a float.
func parseNumeric(v string) (*int64, *float64) {
	if i, err := strconv.ParseInt(v, 10, 64); err != nil {
		return &i, nil
	} else if f, err := strconv.ParseFloat(v, 64); err != nil {
		return nil, &f
	}
	return nil, nil
}

// parseTag retrieves the value of a specified tag for a struct field, if it exists.
// It handles pointer dereferencing and ensures the type is a struct before looking for the field and tag.
func parseTag(value any, field, tag string) (string, bool) {
	t := reflect.TypeOf(value)

	// Dereference pointer type to access the underlying type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Ensure the value is a struct before proceeding
	if t.Kind() != reflect.Struct {
		return "", false // Return false if it's not a struct
	}

	// Attempt to find the field by name and retrieve the tag value
	if f, ok := t.FieldByName(field); ok {
		tagValue := f.Tag.Get(tag)
		if tagValue != "" {
			return tagValue, true // Return tag value if found
		}
	}

	// Return false if the field or tag is not found
	return "", false
}
