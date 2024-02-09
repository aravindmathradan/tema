package validator

import (
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

type errorCode string

const (
	EBLANKFIELD         errorCode = "blank_field"
	EVALUENOTPERMITTED  errorCode = "value_not_permitted"
	EMAXCHARS           errorCode = "max_chars"
	EMINCHARS           errorCode = "min_chars"
	EINVALIDEMAIL       errorCode = "invalid_email"
	EEMAILALREADYEXISTS errorCode = "email_already_exists"
	EINVALIDTOKEN       errorCode = "invalid_token"
	EINVALIDFILTER      errorCode = "invalid_page_filter"
	ENOTFOUND           errorCode = "not_found"
	EACCOUNTINACTIVE    errorCode = "inactive_account"
	EALREADYACTIVE      errorCode = "already_active"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type FieldError struct {
	SubCode errorCode `json:"sub_code"`
	Message string    `json:"message"`
}

type Validator struct {
	Errors map[string]FieldError
}

func New() *Validator {
	return &Validator{Errors: make(map[string]FieldError)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key string, subCode errorCode, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = FieldError{
			SubCode: subCode,
			Message: message,
		}
	}
}

func (v *Validator) Check(ok bool, key string, subCode errorCode, message string) {
	if !ok {
		v.AddError(key, subCode, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func MaxBytes(value string, n int) bool {
	return len(value) <= n
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// Generic function which returns true if a specific value is in a list of
// permitted values.
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}

// Generic function which returns true if all values in a slice are unique.
func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}
