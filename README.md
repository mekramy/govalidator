# govalidator

`govalidator` is a Go package that provides a flexible and extensible validation framework with support for localization. It leverages the `go-playground/validator` package for validation and `mekramy/goi18n` for localization.

## Features

- Custom validation rules
- Localized error messages
- Struct and variable validation
- Support for various data types and formats

## Installation

To install `govalidator`, use the following command:

```sh
go get github.com/mekramy/govalidator
```

## Usage

### Basic Usage

Here's a basic example of how to use `govalidator`:

```go
package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
    "github.com/mekramy/goi18n"
    "github.com/mekramy/govalidator"
    "golang.org/x/text/language"
)

func main() {
    // Initialize the validator
    v := govalidator.NewValidator(
        validator.New(),
        govalidator.WithTranslator(
            goi18n.NewTranslator("en", language.English),
            "",
        ),
    )

    // Add custom validation rule
    v.AddValidation("is_valid", func(fl validator.FieldLevel) bool {
        return fl.Field().String() == "valid"
    })
    v.AddTranslation("en", "is_valid", "{field} must be valid")

    // Validate a variable
    err := v.Var("en", "my_field", "valid", "required,is_valid")
    if err.HasInternalError() {
        fmt.Println("Internal error:", err.InternalError())
    } else if err.HasValidationErrors() {
        fmt.Println("Validation errors:", err)
    } else {
        fmt.Println("Validation passed")
    }
}
```

### Struct Validation

You can also validate structs with `govalidator`:

```go
package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
    "github.com/mekramy/goi18n"
    "github.com/mekramy/govalidator"
    "golang.org/x/text/language"
)

type TestStruct struct {
    Field string `validate:"required,is_valid"`
}

func main() {
    // Initialize the validator
    v := govalidator.NewValidator(
        validator.New(),
        govalidator.WithTranslator(
            goi18n.NewTranslator("en", language.English),
            "",
        ),
    )

    // Add custom validation rule
    v.AddValidation("is_valid", func(fl validator.FieldLevel) bool {
        return fl.Field().String() == "valid"
    })
    v.AddTranslation("en", "is_valid", "{field} must be valid")

    // Validate a struct
    ts := TestStruct{Field: "valid"}
    err := v.Struct("en", ts)
    if err.HasInternalError() {
        fmt.Println("Internal error:", err.InternalError())
    } else if err.HasValidationErrors() {
        fmt.Println("Validation errors:", err)
    } else {
        fmt.Println("Validation passed")
    }
}
```

### Custom Validators

You can add custom validators to `govalidator`:

```go
package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
    "github.com/mekramy/goi18n"
    "github.com/mekramy/govalidator"
    "golang.org/x/text/language"
)

func main() {
    // Initialize the validator
    v := govalidator.NewValidator(
        validator.New(),
        govalidator.WithTranslator(
            goi18n.NewTranslator("en", language.English),
            "",
        ),
    )

    // Add custom validation rule
    v.AddValidation("is_valid", func(fl validator.FieldLevel) bool {
        return fl.Field().String() == "valid"
    })
    v.AddTranslation("en", "is_valid", "{field} must be valid")

    // Validate a variable
    err := v.Var("en", "my_field", "invalid", "required,is_valid")
    if err.HasInternalError() {
        fmt.Println("Internal error:", err.InternalError())
    } else if err.HasValidationErrors() {
        fmt.Println("Validation errors:", err)
    } else {
        fmt.Println("Validation passed")
    }
}
```

## License

This project is licensed under the ISC License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## Acknowledgements

- [go-playground/validator](https://github.com/go-playground/validator)
- [mekramy/goi18n](https://github.com/mekramy/goi18n)
