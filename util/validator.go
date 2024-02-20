package util

import (
	"regexp"

	"github.com/go-playground/validator"
)

func NewValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("phone", validatePhone)
	v.RegisterValidation("password", validatePassword)

	return v
}

func validatePhone(fl validator.FieldLevel) bool {
	phoneField := fl.Field().String()

	regex, _ := regexp.Compile(`^\+62([0-9]{8,11})$`)
	result := regex.MatchString(phoneField)
	return result
}

func validatePassword(fl validator.FieldLevel) bool {
	field := fl.Field().String()

	regexAlpha, _ := regexp.Compile(`[A-Za-z]+`)
	regexDigit, _ := regexp.Compile(`[0-9]+`)
	regexSpecialChar, _ := regexp.Compile(`[\+\-_><{}\(\)\\=;:,.'"!@#$%^*|&?]+`)

	if result := regexAlpha.MatchString(field); !result {
		return result
	}

	if result := regexDigit.MatchString(field); !result {
		return result
	}

	if result := regexSpecialChar.MatchString(field); !result {
		return result
	}
	// regex, _ := regexp.Compile(`^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$`)
	// alpha := `/[A-Za-z]/i
	// numer := `/[0-9]/
	// special := `/[!@#$%\^&*(),\.?;":{}|<>\']/
	// if (password.match(alpha) && password.match(numer) && password.match(speicial))
	// result := regex.MatchString(field)
	return true
	// return result
}
