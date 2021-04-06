package validator

type ValidationResult struct {
	errors map[string]string
}

func (vr *ValidationResult) IsValid() bool {
	return vr.errors == nil
}

func (vr *ValidationResult) Errors() map[string]string {
	return vr.errors
}
