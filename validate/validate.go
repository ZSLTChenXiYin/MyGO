package validate

import (
	"github.com/go-playground/validator"
)

type Validator struct {
	val *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{val: validator.New()}
}

func (v *Validator) Init() error {
	if err := v.val.RegisterValidation(VALIDATE_TAG_CAPTCHA, validateTagCaptcha); err != nil {
		return err
	}

	if err := v.val.RegisterValidation(VALIDATE_TAG_GENDER, validateTagGender); err != nil {
		return err
	}

	return nil
}

func (v *Validator) Validate(a any) error {
	return v.val.Struct(a)
}

func (v *Validator) Validator() *validator.Validate {
	return v.val
}
