package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	jsonValidator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var _ Validator = (*JsonValidator)(nil)

type JsonValidator struct {
	validator *jsonValidator.Validate
	trans     ut.Translator
}

func NewJsonValidator() *JsonValidator {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validator := jsonValidator.New()

	en_translations.RegisterDefaultTranslations(validator, trans)

	return &JsonValidator{
		validator: validator,
		trans:     trans,
	}
}

func (jv *JsonValidator) Validate(v interface{}) ValidationResult {
	err := jv.validator.Struct(v)
	if err != nil {
		errs := err.(jsonValidator.ValidationErrors)
		return ValidationResult{
			errors: errs.Translate(jv.trans),
		}
	}
	return ValidationResult{}
}
