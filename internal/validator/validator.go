package validator

import "regexp"

// Errors is a types alias of map[string]string
type Errors = map[string]string

type Validator struct {
	Errors Errors
}

// New creates and returns an instance of a Validator.
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// Valid returns true if the Validator.Errors map doesn't contain any entries.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError adds a new error to Validator.Errors with a given key and message.
// If key already exists in the map, the message won't be added
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check checks a boolean expression. If the expression evaluates to `false`,
// an error will be added to Validator.Errors
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// In checks if a string value is present in a list of strings.
func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

// Matches checks if a string value matches a specific regexp pattern.
func Matches(rx *regexp.Regexp, value string) bool {
	return rx.MatchString(value)
}
