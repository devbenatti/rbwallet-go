package validator

type Validator interface {
	Validate(v interface{}) ValidationResult
}
