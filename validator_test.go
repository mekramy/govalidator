package govalidator_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/mekramy/goi18n"
	"github.com/mekramy/govalidator"
	"golang.org/x/text/language"
)

func TestValidator(t *testing.T) {
	// initialize
	v := govalidator.NewValidator(
		validator.New(),
		govalidator.WithTranslator(
			goi18n.NewTranslator("en", language.English),
			"",
		),
	)
	v.AddValidation("is_valid", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "valid"
	})
	v.AddTranslation("en", "is_valid", "{field} must be valid")

	t.Run("Var", func(t *testing.T) {
		err := v.Var("en", "my_field", "valid", "required,is_valid")
		if err.HasInternalError() {
			t.Fatal(err.InternalError())
		} else if err.HasValidationErrors() {
			t.Fatal("expected no errors, got some")
		}
	})

	t.Run("VarWithValue", func(t *testing.T) {
		err := v.VarWithValue("en", "my_field", "value1", "value2", "eqfield")
		if err.HasInternalError() {
			t.Fatal(err.InternalError())
		} else if !err.HasValidationErrors() {
			t.Fatal("expected validation errors, got none")
		}
	})

	t.Run("Struct", func(t *testing.T) {
		type TestStruct struct {
			Field string `validate:"required,is_valid"`
		}
		ts := TestStruct{Field: "valid"}
		err := v.Struct("en", ts)
		if err.HasInternalError() {
			t.Fatal(err.InternalError())
		} else if err.HasValidationErrors() {
			t.Fatal("expected no errors, got some")
		}
	})

	t.Run("InvalidStruct", func(t *testing.T) {
		type TestStruct struct {
			Field string `validate:"required,is_valid"`
		}
		ts := TestStruct{Field: "invalid"}
		err := v.Struct("en", ts)
		if err.HasInternalError() {
			t.Fatal(err.InternalError())
		} else if !err.HasValidationErrors() {
			t.Fatal("expected validation errors, got none")
		}
	})

	t.Run("StructExcept", func(t *testing.T) {
		type TestStruct struct {
			Field1 string `validate:"required,is_valid"`
			Field2 string `validate:"required"`
		}
		ts := TestStruct{Field1: "valid", Field2: ""}
		err := v.StructExpect("en", ts, "Field2")
		if err.HasInternalError() {
			t.Fatal(err.InternalError())
		} else if err.HasValidationErrors() {
			t.Fatal("expected no errors, got some")
		}
	})

	t.Run("StructPartial", func(t *testing.T) {
		type TestStruct struct {
			Field1 string `validate:"required,is_valid"`
			Field2 string `validate:"required"`
		}
		ts := TestStruct{Field1: "invalid", Field2: ""}
		err := v.StructPartial("en", ts, "Field1")
		if err.HasInternalError() {
			t.Fatal(err.InternalError())
		} else if !err.HasValidationErrors() {
			t.Fatal("expected validation errors, got none")
		}
	})
}
