package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	validate *validator.Validate
}

func (v *Validator) valid(value string, tag string) bool {
	err := v.validate.Var(value, tag)
	if err != nil {
		return false
	}

	return true
}
