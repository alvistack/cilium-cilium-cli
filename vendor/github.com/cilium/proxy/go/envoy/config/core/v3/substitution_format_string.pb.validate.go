// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/core/v3/substitution_format_string.proto

package corev3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on SubstitutionFormatString with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SubstitutionFormatString) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubstitutionFormatString with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SubstitutionFormatStringMultiError, or nil if none found.
func (m *SubstitutionFormatString) ValidateAll() error {
	return m.validate(true)
}

func (m *SubstitutionFormatString) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OmitEmptyValues

	if !_SubstitutionFormatString_ContentType_Pattern.MatchString(m.GetContentType()) {
		err := SubstitutionFormatStringValidationError{
			field:  "ContentType",
			reason: "value does not match regex pattern \"^[^\\x00\\n\\r]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetFormatters() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SubstitutionFormatStringValidationError{
						field:  fmt.Sprintf("Formatters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SubstitutionFormatStringValidationError{
						field:  fmt.Sprintf("Formatters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SubstitutionFormatStringValidationError{
					field:  fmt.Sprintf("Formatters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.Format.(type) {

	case *SubstitutionFormatString_TextFormat:
		// no validation rules for TextFormat

	case *SubstitutionFormatString_JsonFormat:

		if m.GetJsonFormat() == nil {
			err := SubstitutionFormatStringValidationError{
				field:  "JsonFormat",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetJsonFormat()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SubstitutionFormatStringValidationError{
						field:  "JsonFormat",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SubstitutionFormatStringValidationError{
						field:  "JsonFormat",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetJsonFormat()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SubstitutionFormatStringValidationError{
					field:  "JsonFormat",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *SubstitutionFormatString_TextFormatSource:

		if all {
			switch v := interface{}(m.GetTextFormatSource()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SubstitutionFormatStringValidationError{
						field:  "TextFormatSource",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SubstitutionFormatStringValidationError{
						field:  "TextFormatSource",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTextFormatSource()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SubstitutionFormatStringValidationError{
					field:  "TextFormatSource",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := SubstitutionFormatStringValidationError{
			field:  "Format",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return SubstitutionFormatStringMultiError(errors)
	}
	return nil
}

// SubstitutionFormatStringMultiError is an error wrapping multiple validation
// errors returned by SubstitutionFormatString.ValidateAll() if the designated
// constraints aren't met.
type SubstitutionFormatStringMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubstitutionFormatStringMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubstitutionFormatStringMultiError) AllErrors() []error { return m }

// SubstitutionFormatStringValidationError is the validation error returned by
// SubstitutionFormatString.Validate if the designated constraints aren't met.
type SubstitutionFormatStringValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubstitutionFormatStringValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubstitutionFormatStringValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubstitutionFormatStringValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubstitutionFormatStringValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubstitutionFormatStringValidationError) ErrorName() string {
	return "SubstitutionFormatStringValidationError"
}

// Error satisfies the builtin error interface
func (e SubstitutionFormatStringValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubstitutionFormatString.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubstitutionFormatStringValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubstitutionFormatStringValidationError{}

var _SubstitutionFormatString_ContentType_Pattern = regexp.MustCompile("^[^\x00\n\r]*$")
