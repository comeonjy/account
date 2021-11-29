// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: configs/config.proto

package configs

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Config) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := _Config_Mode_InLookup[m.GetMode()]; !ok {
		return ConfigValidationError{
			field:  "Mode",
			reason: "value must be in list [debug normal]",
		}
	}

	if utf8.RuneCountInString(m.GetGrpcAddr()) < 3 {
		return ConfigValidationError{
			field:  "GrpcAddr",
			reason: "value length must be at least 3 runes",
		}
	}

	if utf8.RuneCountInString(m.GetHttpAddr()) < 3 {
		return ConfigValidationError{
			field:  "HttpAddr",
			reason: "value length must be at least 3 runes",
		}
	}

	if utf8.RuneCountInString(m.GetPprofAddr()) < 3 {
		return ConfigValidationError{
			field:  "PprofAddr",
			reason: "value length must be at least 3 runes",
		}
	}

	// no validation rules for ApmUrl

	// no validation rules for MysqlConf

	// no validation rules for TenSecretId

	// no validation rules for TenSecretKey

	// no validation rules for TenSmsConf

	// no validation rules for WechatMiniAppid

	// no validation rules for WechatMiniSecret

	// no validation rules for JwtKey

	// no validation rules for YunpianApiKey

	// no validation rules for RedisOption

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return ConfigValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	return nil
}

func (m *Config) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *Config) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}

var _Config_Mode_InLookup = map[string]struct{}{
	"debug":  {},
	"normal": {},
}
