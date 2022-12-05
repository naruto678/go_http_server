package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRgx = regexp.MustCompile(".*@.*\\.com$")

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}

}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) MaxLen(field string, length int) {
	value := f.Get(field)
	if utf8.RuneCountInString(value) < length {
		f.Errors.Add(field, fmt.Sprintf("Field %s cannot have more than %d characters", field, length))
	}
}

func (f *Form) PermittedValues(field string, values ...string) {
	current_val := f.Get(field)
	for _, val := range values {
		if val == current_val {
			return
		}
	}
	f.Errors.Add(field, fmt.Sprintf("Field %s can only have values [%v]", field, values))
}

func (f *Form) IsValid() bool {
	return len(f.Errors) == 0
}

func (f *Form) MinLength(field string, length int) {
	value := f.Get(field)
	if utf8.RuneCountInString(value) < length {
		f.Errors.Add(field, fmt.Sprintf("%s has minimum field length of %d", field, length))
		return
	}
}

func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field is invalid")
	}

}
