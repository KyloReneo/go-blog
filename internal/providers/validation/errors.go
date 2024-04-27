package validation

func ErrorMessages() map[string]string {
	return map[string]string{
		"requierd": "The field is required.",
		"email":    "The field must have a valid email address.",
		"min":      "Should be more than the limit.",
		"max":      "Should be less than the limit.",
	}
}
