//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/core/v3/config_source.proto

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

// Validate checks the field values on ApiConfigSource with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ApiConfigSource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ApiConfigSource with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ApiConfigSourceMultiError, or nil if none found.
func (m *ApiConfigSource) ValidateAll() error {
	return m.validate(true)
}

func (m *ApiConfigSource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := ApiConfigSource_ApiType_name[int32(m.GetApiType())]; !ok {
		err := ApiConfigSourceValidationError{
			field:  "ApiType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := ApiVersion_name[int32(m.GetTransportApiVersion())]; !ok {
		err := ApiConfigSourceValidationError{
			field:  "TransportApiVersion",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetGrpcServices() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ApiConfigSourceValidationError{
						field:  fmt.Sprintf("GrpcServices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ApiConfigSourceValidationError{
						field:  fmt.Sprintf("GrpcServices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ApiConfigSourceValidationError{
					field:  fmt.Sprintf("GrpcServices[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetRefreshDelay()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApiConfigSourceValidationError{
					field:  "RefreshDelay",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApiConfigSourceValidationError{
					field:  "RefreshDelay",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRefreshDelay()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApiConfigSourceValidationError{
				field:  "RefreshDelay",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if d := m.GetRequestTimeout(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = ApiConfigSourceValidationError{
				field:  "RequestTimeout",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := ApiConfigSourceValidationError{
					field:  "RequestTimeout",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if all {
		switch v := interface{}(m.GetRateLimitSettings()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApiConfigSourceValidationError{
					field:  "RateLimitSettings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApiConfigSourceValidationError{
					field:  "RateLimitSettings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRateLimitSettings()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApiConfigSourceValidationError{
				field:  "RateLimitSettings",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for SetNodeOnFirstMessageOnly

	for idx, item := range m.GetConfigValidators() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ApiConfigSourceValidationError{
						field:  fmt.Sprintf("ConfigValidators[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ApiConfigSourceValidationError{
						field:  fmt.Sprintf("ConfigValidators[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ApiConfigSourceValidationError{
					field:  fmt.Sprintf("ConfigValidators[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ApiConfigSourceMultiError(errors)
	}

	return nil
}

// ApiConfigSourceMultiError is an error wrapping multiple validation errors
// returned by ApiConfigSource.ValidateAll() if the designated constraints
// aren't met.
type ApiConfigSourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApiConfigSourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApiConfigSourceMultiError) AllErrors() []error { return m }

// ApiConfigSourceValidationError is the validation error returned by
// ApiConfigSource.Validate if the designated constraints aren't met.
type ApiConfigSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApiConfigSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApiConfigSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApiConfigSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApiConfigSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApiConfigSourceValidationError) ErrorName() string { return "ApiConfigSourceValidationError" }

// Error satisfies the builtin error interface
func (e ApiConfigSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApiConfigSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApiConfigSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApiConfigSourceValidationError{}

// Validate checks the field values on AggregatedConfigSource with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AggregatedConfigSource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AggregatedConfigSource with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AggregatedConfigSourceMultiError, or nil if none found.
func (m *AggregatedConfigSource) ValidateAll() error {
	return m.validate(true)
}

func (m *AggregatedConfigSource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AggregatedConfigSourceMultiError(errors)
	}

	return nil
}

// AggregatedConfigSourceMultiError is an error wrapping multiple validation
// errors returned by AggregatedConfigSource.ValidateAll() if the designated
// constraints aren't met.
type AggregatedConfigSourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AggregatedConfigSourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AggregatedConfigSourceMultiError) AllErrors() []error { return m }

// AggregatedConfigSourceValidationError is the validation error returned by
// AggregatedConfigSource.Validate if the designated constraints aren't met.
type AggregatedConfigSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AggregatedConfigSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AggregatedConfigSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AggregatedConfigSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AggregatedConfigSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AggregatedConfigSourceValidationError) ErrorName() string {
	return "AggregatedConfigSourceValidationError"
}

// Error satisfies the builtin error interface
func (e AggregatedConfigSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAggregatedConfigSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AggregatedConfigSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AggregatedConfigSourceValidationError{}

// Validate checks the field values on SelfConfigSource with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SelfConfigSource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SelfConfigSource with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SelfConfigSourceMultiError, or nil if none found.
func (m *SelfConfigSource) ValidateAll() error {
	return m.validate(true)
}

func (m *SelfConfigSource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := ApiVersion_name[int32(m.GetTransportApiVersion())]; !ok {
		err := SelfConfigSourceValidationError{
			field:  "TransportApiVersion",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SelfConfigSourceMultiError(errors)
	}

	return nil
}

// SelfConfigSourceMultiError is an error wrapping multiple validation errors
// returned by SelfConfigSource.ValidateAll() if the designated constraints
// aren't met.
type SelfConfigSourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SelfConfigSourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SelfConfigSourceMultiError) AllErrors() []error { return m }

// SelfConfigSourceValidationError is the validation error returned by
// SelfConfigSource.Validate if the designated constraints aren't met.
type SelfConfigSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SelfConfigSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SelfConfigSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SelfConfigSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SelfConfigSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SelfConfigSourceValidationError) ErrorName() string { return "SelfConfigSourceValidationError" }

// Error satisfies the builtin error interface
func (e SelfConfigSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSelfConfigSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SelfConfigSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SelfConfigSourceValidationError{}

// Validate checks the field values on RateLimitSettings with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RateLimitSettings) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RateLimitSettings with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RateLimitSettingsMultiError, or nil if none found.
func (m *RateLimitSettings) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitSettings) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMaxTokens()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RateLimitSettingsValidationError{
					field:  "MaxTokens",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RateLimitSettingsValidationError{
					field:  "MaxTokens",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMaxTokens()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitSettingsValidationError{
				field:  "MaxTokens",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetFillRate(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			err := RateLimitSettingsValidationError{
				field:  "FillRate",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return RateLimitSettingsMultiError(errors)
	}

	return nil
}

// RateLimitSettingsMultiError is an error wrapping multiple validation errors
// returned by RateLimitSettings.ValidateAll() if the designated constraints
// aren't met.
type RateLimitSettingsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitSettingsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitSettingsMultiError) AllErrors() []error { return m }

// RateLimitSettingsValidationError is the validation error returned by
// RateLimitSettings.Validate if the designated constraints aren't met.
type RateLimitSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitSettingsValidationError) ErrorName() string {
	return "RateLimitSettingsValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitSettingsValidationError{}

// Validate checks the field values on PathConfigSource with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PathConfigSource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PathConfigSource with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PathConfigSourceMultiError, or nil if none found.
func (m *PathConfigSource) ValidateAll() error {
	return m.validate(true)
}

func (m *PathConfigSource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetPath()) < 1 {
		err := PathConfigSourceValidationError{
			field:  "Path",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetWatchedDirectory()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PathConfigSourceValidationError{
					field:  "WatchedDirectory",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PathConfigSourceValidationError{
					field:  "WatchedDirectory",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWatchedDirectory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PathConfigSourceValidationError{
				field:  "WatchedDirectory",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PathConfigSourceMultiError(errors)
	}

	return nil
}

// PathConfigSourceMultiError is an error wrapping multiple validation errors
// returned by PathConfigSource.ValidateAll() if the designated constraints
// aren't met.
type PathConfigSourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PathConfigSourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PathConfigSourceMultiError) AllErrors() []error { return m }

// PathConfigSourceValidationError is the validation error returned by
// PathConfigSource.Validate if the designated constraints aren't met.
type PathConfigSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PathConfigSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PathConfigSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PathConfigSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PathConfigSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PathConfigSourceValidationError) ErrorName() string { return "PathConfigSourceValidationError" }

// Error satisfies the builtin error interface
func (e PathConfigSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPathConfigSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PathConfigSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PathConfigSourceValidationError{}

// Validate checks the field values on ConfigSource with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ConfigSource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfigSource with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ConfigSourceMultiError, or
// nil if none found.
func (m *ConfigSource) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfigSource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetAuthorities() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ConfigSourceValidationError{
						field:  fmt.Sprintf("Authorities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ConfigSourceValidationError{
						field:  fmt.Sprintf("Authorities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigSourceValidationError{
					field:  fmt.Sprintf("Authorities[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetInitialFetchTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfigSourceValidationError{
					field:  "InitialFetchTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfigSourceValidationError{
					field:  "InitialFetchTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInitialFetchTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigSourceValidationError{
				field:  "InitialFetchTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := ApiVersion_name[int32(m.GetResourceApiVersion())]; !ok {
		err := ConfigSourceValidationError{
			field:  "ResourceApiVersion",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	oneofConfigSourceSpecifierPresent := false
	switch v := m.ConfigSourceSpecifier.(type) {
	case *ConfigSource_Path:
		if v == nil {
			err := ConfigSourceValidationError{
				field:  "ConfigSourceSpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofConfigSourceSpecifierPresent = true
		// no validation rules for Path
	case *ConfigSource_PathConfigSource:
		if v == nil {
			err := ConfigSourceValidationError{
				field:  "ConfigSourceSpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofConfigSourceSpecifierPresent = true

		if all {
			switch v := interface{}(m.GetPathConfigSource()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ConfigSourceValidationError{
						field:  "PathConfigSource",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ConfigSourceValidationError{
						field:  "PathConfigSource",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPathConfigSource()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigSourceValidationError{
					field:  "PathConfigSource",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ConfigSource_ApiConfigSource:
		if v == nil {
			err := ConfigSourceValidationError{
				field:  "ConfigSourceSpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofConfigSourceSpecifierPresent = true

		if all {
			switch v := interface{}(m.GetApiConfigSource()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ConfigSourceValidationError{
						field:  "ApiConfigSource",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ConfigSourceValidationError{
						field:  "ApiConfigSource",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetApiConfigSource()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigSourceValidationError{
					field:  "ApiConfigSource",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ConfigSource_Ads:
		if v == nil {
			err := ConfigSourceValidationError{
				field:  "ConfigSourceSpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofConfigSourceSpecifierPresent = true

		if all {
			switch v := interface{}(m.GetAds()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ConfigSourceValidationError{
						field:  "Ads",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ConfigSourceValidationError{
						field:  "Ads",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAds()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigSourceValidationError{
					field:  "Ads",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ConfigSource_Self:
		if v == nil {
			err := ConfigSourceValidationError{
				field:  "ConfigSourceSpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofConfigSourceSpecifierPresent = true

		if all {
			switch v := interface{}(m.GetSelf()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ConfigSourceValidationError{
						field:  "Self",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ConfigSourceValidationError{
						field:  "Self",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSelf()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigSourceValidationError{
					field:  "Self",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofConfigSourceSpecifierPresent {
		err := ConfigSourceValidationError{
			field:  "ConfigSourceSpecifier",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ConfigSourceMultiError(errors)
	}

	return nil
}

// ConfigSourceMultiError is an error wrapping multiple validation errors
// returned by ConfigSource.ValidateAll() if the designated constraints aren't met.
type ConfigSourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigSourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigSourceMultiError) AllErrors() []error { return m }

// ConfigSourceValidationError is the validation error returned by
// ConfigSource.Validate if the designated constraints aren't met.
type ConfigSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigSourceValidationError) ErrorName() string { return "ConfigSourceValidationError" }

// Error satisfies the builtin error interface
func (e ConfigSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigSourceValidationError{}

// Validate checks the field values on ExtensionConfigSource with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ExtensionConfigSource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExtensionConfigSource with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExtensionConfigSourceMultiError, or nil if none found.
func (m *ExtensionConfigSource) ValidateAll() error {
	return m.validate(true)
}

func (m *ExtensionConfigSource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetConfigSource() == nil {
		err := ExtensionConfigSourceValidationError{
			field:  "ConfigSource",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if a := m.GetConfigSource(); a != nil {

	}

	if all {
		switch v := interface{}(m.GetDefaultConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ExtensionConfigSourceValidationError{
					field:  "DefaultConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ExtensionConfigSourceValidationError{
					field:  "DefaultConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDefaultConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtensionConfigSourceValidationError{
				field:  "DefaultConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ApplyDefaultConfigWithoutWarming

	if len(m.GetTypeUrls()) < 1 {
		err := ExtensionConfigSourceValidationError{
			field:  "TypeUrls",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ExtensionConfigSourceMultiError(errors)
	}

	return nil
}

// ExtensionConfigSourceMultiError is an error wrapping multiple validation
// errors returned by ExtensionConfigSource.ValidateAll() if the designated
// constraints aren't met.
type ExtensionConfigSourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExtensionConfigSourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExtensionConfigSourceMultiError) AllErrors() []error { return m }

// ExtensionConfigSourceValidationError is the validation error returned by
// ExtensionConfigSource.Validate if the designated constraints aren't met.
type ExtensionConfigSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtensionConfigSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtensionConfigSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtensionConfigSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtensionConfigSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtensionConfigSourceValidationError) ErrorName() string {
	return "ExtensionConfigSourceValidationError"
}

// Error satisfies the builtin error interface
func (e ExtensionConfigSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtensionConfigSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtensionConfigSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtensionConfigSourceValidationError{}
