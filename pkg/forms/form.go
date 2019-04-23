package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// Form holds form data and errors
type Form struct {
	url.Values
	Errors errors
}

// New initializes a new custom Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks that data fields in the form are pesent and not blank
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MaxLength checks if a given field has specific length
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maxumum is %d characters)", d))
	}
}

// PermittedValues checks if a given field match a set of specific Values
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

// Valid returns true if there are no Errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
