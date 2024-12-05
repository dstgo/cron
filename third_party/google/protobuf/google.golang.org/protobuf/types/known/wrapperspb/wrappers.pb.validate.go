// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: wrappers.proto

package wrapperspb

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

// Validate checks the field values on DoubleValue with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DoubleValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DoubleValue with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DoubleValueMultiError, or
// nil if none found.
func (m *DoubleValue) ValidateAll() error {
	return m.validate(true)
}

func (m *DoubleValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return DoubleValueMultiError(errors)
	}

	return nil
}

// DoubleValueMultiError is an error wrapping multiple validation errors
// returned by DoubleValue.ValidateAll() if the designated constraints aren't met.
type DoubleValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DoubleValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DoubleValueMultiError) AllErrors() []error { return m }

// DoubleValueValidationError is the validation error returned by
// DoubleValue.Validate if the designated constraints aren't met.
type DoubleValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DoubleValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DoubleValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DoubleValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DoubleValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DoubleValueValidationError) ErrorName() string { return "DoubleValueValidationError" }

// Error satisfies the builtin error interface
func (e DoubleValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDoubleValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DoubleValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DoubleValueValidationError{}

// Validate checks the field values on FloatValue with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FloatValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FloatValue with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FloatValueMultiError, or
// nil if none found.
func (m *FloatValue) ValidateAll() error {
	return m.validate(true)
}

func (m *FloatValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return FloatValueMultiError(errors)
	}

	return nil
}

// FloatValueMultiError is an error wrapping multiple validation errors
// returned by FloatValue.ValidateAll() if the designated constraints aren't met.
type FloatValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FloatValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FloatValueMultiError) AllErrors() []error { return m }

// FloatValueValidationError is the validation error returned by
// FloatValue.Validate if the designated constraints aren't met.
type FloatValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FloatValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FloatValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FloatValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FloatValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FloatValueValidationError) ErrorName() string { return "FloatValueValidationError" }

// Error satisfies the builtin error interface
func (e FloatValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFloatValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FloatValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FloatValueValidationError{}

// Validate checks the field values on Int64Value with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Int64Value) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Int64Value with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Int64ValueMultiError, or
// nil if none found.
func (m *Int64Value) ValidateAll() error {
	return m.validate(true)
}

func (m *Int64Value) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return Int64ValueMultiError(errors)
	}

	return nil
}

// Int64ValueMultiError is an error wrapping multiple validation errors
// returned by Int64Value.ValidateAll() if the designated constraints aren't met.
type Int64ValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Int64ValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Int64ValueMultiError) AllErrors() []error { return m }

// Int64ValueValidationError is the validation error returned by
// Int64Value.Validate if the designated constraints aren't met.
type Int64ValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Int64ValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Int64ValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Int64ValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Int64ValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Int64ValueValidationError) ErrorName() string { return "Int64ValueValidationError" }

// Error satisfies the builtin error interface
func (e Int64ValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInt64Value.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Int64ValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Int64ValueValidationError{}

// Validate checks the field values on UInt64Value with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UInt64Value) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UInt64Value with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UInt64ValueMultiError, or
// nil if none found.
func (m *UInt64Value) ValidateAll() error {
	return m.validate(true)
}

func (m *UInt64Value) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return UInt64ValueMultiError(errors)
	}

	return nil
}

// UInt64ValueMultiError is an error wrapping multiple validation errors
// returned by UInt64Value.ValidateAll() if the designated constraints aren't met.
type UInt64ValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UInt64ValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UInt64ValueMultiError) AllErrors() []error { return m }

// UInt64ValueValidationError is the validation error returned by
// UInt64Value.Validate if the designated constraints aren't met.
type UInt64ValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UInt64ValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UInt64ValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UInt64ValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UInt64ValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UInt64ValueValidationError) ErrorName() string { return "UInt64ValueValidationError" }

// Error satisfies the builtin error interface
func (e UInt64ValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUInt64Value.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UInt64ValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UInt64ValueValidationError{}

// Validate checks the field values on Int32Value with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Int32Value) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Int32Value with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Int32ValueMultiError, or
// nil if none found.
func (m *Int32Value) ValidateAll() error {
	return m.validate(true)
}

func (m *Int32Value) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return Int32ValueMultiError(errors)
	}

	return nil
}

// Int32ValueMultiError is an error wrapping multiple validation errors
// returned by Int32Value.ValidateAll() if the designated constraints aren't met.
type Int32ValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Int32ValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Int32ValueMultiError) AllErrors() []error { return m }

// Int32ValueValidationError is the validation error returned by
// Int32Value.Validate if the designated constraints aren't met.
type Int32ValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Int32ValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Int32ValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Int32ValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Int32ValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Int32ValueValidationError) ErrorName() string { return "Int32ValueValidationError" }

// Error satisfies the builtin error interface
func (e Int32ValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInt32Value.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Int32ValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Int32ValueValidationError{}

// Validate checks the field values on UInt32Value with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UInt32Value) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UInt32Value with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UInt32ValueMultiError, or
// nil if none found.
func (m *UInt32Value) ValidateAll() error {
	return m.validate(true)
}

func (m *UInt32Value) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return UInt32ValueMultiError(errors)
	}

	return nil
}

// UInt32ValueMultiError is an error wrapping multiple validation errors
// returned by UInt32Value.ValidateAll() if the designated constraints aren't met.
type UInt32ValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UInt32ValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UInt32ValueMultiError) AllErrors() []error { return m }

// UInt32ValueValidationError is the validation error returned by
// UInt32Value.Validate if the designated constraints aren't met.
type UInt32ValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UInt32ValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UInt32ValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UInt32ValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UInt32ValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UInt32ValueValidationError) ErrorName() string { return "UInt32ValueValidationError" }

// Error satisfies the builtin error interface
func (e UInt32ValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUInt32Value.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UInt32ValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UInt32ValueValidationError{}

// Validate checks the field values on BoolValue with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BoolValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BoolValue with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BoolValueMultiError, or nil
// if none found.
func (m *BoolValue) ValidateAll() error {
	return m.validate(true)
}

func (m *BoolValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return BoolValueMultiError(errors)
	}

	return nil
}

// BoolValueMultiError is an error wrapping multiple validation errors returned
// by BoolValue.ValidateAll() if the designated constraints aren't met.
type BoolValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BoolValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BoolValueMultiError) AllErrors() []error { return m }

// BoolValueValidationError is the validation error returned by
// BoolValue.Validate if the designated constraints aren't met.
type BoolValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BoolValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BoolValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BoolValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BoolValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BoolValueValidationError) ErrorName() string { return "BoolValueValidationError" }

// Error satisfies the builtin error interface
func (e BoolValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBoolValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BoolValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BoolValueValidationError{}

// Validate checks the field values on StringValue with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StringValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StringValue with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StringValueMultiError, or
// nil if none found.
func (m *StringValue) ValidateAll() error {
	return m.validate(true)
}

func (m *StringValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return StringValueMultiError(errors)
	}

	return nil
}

// StringValueMultiError is an error wrapping multiple validation errors
// returned by StringValue.ValidateAll() if the designated constraints aren't met.
type StringValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StringValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StringValueMultiError) AllErrors() []error { return m }

// StringValueValidationError is the validation error returned by
// StringValue.Validate if the designated constraints aren't met.
type StringValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringValueValidationError) ErrorName() string { return "StringValueValidationError" }

// Error satisfies the builtin error interface
func (e StringValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringValueValidationError{}

// Validate checks the field values on BytesValue with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BytesValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BytesValue with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BytesValueMultiError, or
// nil if none found.
func (m *BytesValue) ValidateAll() error {
	return m.validate(true)
}

func (m *BytesValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return BytesValueMultiError(errors)
	}

	return nil
}

// BytesValueMultiError is an error wrapping multiple validation errors
// returned by BytesValue.ValidateAll() if the designated constraints aren't met.
type BytesValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BytesValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BytesValueMultiError) AllErrors() []error { return m }

// BytesValueValidationError is the validation error returned by
// BytesValue.Validate if the designated constraints aren't met.
type BytesValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BytesValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BytesValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BytesValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BytesValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BytesValueValidationError) ErrorName() string { return "BytesValueValidationError" }

// Error satisfies the builtin error interface
func (e BytesValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBytesValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BytesValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BytesValueValidationError{}
