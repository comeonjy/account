// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/v1/account.proto

package v1

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

// Validate checks the field values on Result with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Result) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResultValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ResultValidationError is the validation error returned by Result.Validate if
// the designated constraints aren't met.
type ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultValidationError) ErrorName() string { return "ResultValidationError" }

// Error satisfies the builtin error interface
func (e ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultValidationError{}

// Validate checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Empty) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// EmptyValidationError is the validation error returned by Empty.Validate if
// the designated constraints aren't met.
type EmptyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyValidationError) ErrorName() string { return "EmptyValidationError" }

// Error satisfies the builtin error interface
func (e EmptyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmpty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyValidationError{}

// Validate checks the field values on SendVerificationCodeReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SendVerificationCodeReq) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetAccount()) < 1 {
		return SendVerificationCodeReqValidationError{
			field:  "Account",
			reason: "value length must be at least 1 runes",
		}
	}

	if _, ok := _SendVerificationCodeReq_Type_InLookup[m.GetType()]; !ok {
		return SendVerificationCodeReqValidationError{
			field:  "Type",
			reason: "value must be in list [mobile email]",
		}
	}

	return nil
}

// SendVerificationCodeReqValidationError is the validation error returned by
// SendVerificationCodeReq.Validate if the designated constraints aren't met.
type SendVerificationCodeReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendVerificationCodeReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendVerificationCodeReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendVerificationCodeReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendVerificationCodeReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendVerificationCodeReqValidationError) ErrorName() string {
	return "SendVerificationCodeReqValidationError"
}

// Error satisfies the builtin error interface
func (e SendVerificationCodeReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendVerificationCodeReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendVerificationCodeReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendVerificationCodeReqValidationError{}

var _SendVerificationCodeReq_Type_InLookup = map[string]struct{}{
	"mobile": {},
	"email":  {},
}

// Validate checks the field values on LoginReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LoginReq) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetAccount()) < 1 {
		return LoginReqValidationError{
			field:  "Account",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetPassword() != "" {

		if utf8.RuneCountInString(m.GetPassword()) < 1 {
			return LoginReqValidationError{
				field:  "Password",
				reason: "value length must be at least 1 runes",
			}
		}

	}

	if m.GetCode() != "" {

		if utf8.RuneCountInString(m.GetCode()) < 1 {
			return LoginReqValidationError{
				field:  "Code",
				reason: "value length must be at least 1 runes",
			}
		}

	}

	if _, ok := _LoginReq_Type_InLookup[m.GetType()]; !ok {
		return LoginReqValidationError{
			field:  "Type",
			reason: "value must be in list [mobile email password]",
		}
	}

	return nil
}

// LoginReqValidationError is the validation error returned by
// LoginReq.Validate if the designated constraints aren't met.
type LoginReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReqValidationError) ErrorName() string { return "LoginReqValidationError" }

// Error satisfies the builtin error interface
func (e LoginReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReqValidationError{}

var _LoginReq_Type_InLookup = map[string]struct{}{
	"mobile":   {},
	"email":    {},
	"password": {},
}

// Validate checks the field values on LoginResp with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LoginResp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	return nil
}

// LoginRespValidationError is the validation error returned by
// LoginResp.Validate if the designated constraints aren't met.
type LoginRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRespValidationError) ErrorName() string { return "LoginRespValidationError" }

// Error satisfies the builtin error interface
func (e LoginRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRespValidationError{}

// Validate checks the field values on GetMiniQRCodeReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetMiniQRCodeReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	return nil
}

// GetMiniQRCodeReqValidationError is the validation error returned by
// GetMiniQRCodeReq.Validate if the designated constraints aren't met.
type GetMiniQRCodeReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMiniQRCodeReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMiniQRCodeReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMiniQRCodeReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMiniQRCodeReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMiniQRCodeReqValidationError) ErrorName() string { return "GetMiniQRCodeReqValidationError" }

// Error satisfies the builtin error interface
func (e GetMiniQRCodeReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMiniQRCodeReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMiniQRCodeReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMiniQRCodeReqValidationError{}

// Validate checks the field values on GetMiniQRCodeResp with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetMiniQRCodeResp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Qrcode

	return nil
}

// GetMiniQRCodeRespValidationError is the validation error returned by
// GetMiniQRCodeResp.Validate if the designated constraints aren't met.
type GetMiniQRCodeRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMiniQRCodeRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMiniQRCodeRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMiniQRCodeRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMiniQRCodeRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMiniQRCodeRespValidationError) ErrorName() string {
	return "GetMiniQRCodeRespValidationError"
}

// Error satisfies the builtin error interface
func (e GetMiniQRCodeRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMiniQRCodeResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMiniQRCodeRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMiniQRCodeRespValidationError{}

// Validate checks the field values on MiniLoginReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MiniLoginReq) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetCode()) < 1 {
		return MiniLoginReqValidationError{
			field:  "Code",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// MiniLoginReqValidationError is the validation error returned by
// MiniLoginReq.Validate if the designated constraints aren't met.
type MiniLoginReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MiniLoginReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MiniLoginReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MiniLoginReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MiniLoginReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MiniLoginReqValidationError) ErrorName() string { return "MiniLoginReqValidationError" }

// Error satisfies the builtin error interface
func (e MiniLoginReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMiniLoginReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MiniLoginReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MiniLoginReqValidationError{}

// Validate checks the field values on MiniLoginResp with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MiniLoginResp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	if v, ok := interface{}(m.GetUserInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MiniLoginRespValidationError{
				field:  "UserInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// MiniLoginRespValidationError is the validation error returned by
// MiniLoginResp.Validate if the designated constraints aren't met.
type MiniLoginRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MiniLoginRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MiniLoginRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MiniLoginRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MiniLoginRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MiniLoginRespValidationError) ErrorName() string { return "MiniLoginRespValidationError" }

// Error satisfies the builtin error interface
func (e MiniLoginRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMiniLoginResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MiniLoginRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MiniLoginRespValidationError{}

// Validate checks the field values on UserInfo with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *UserInfo) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetNickName()) < 1 {
		return UserInfoValidationError{
			field:  "NickName",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetAvatarUrl()) < 1 {
		return UserInfoValidationError{
			field:  "AvatarUrl",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// UserInfoValidationError is the validation error returned by
// UserInfo.Validate if the designated constraints aren't met.
type UserInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserInfoValidationError) ErrorName() string { return "UserInfoValidationError" }

// Error satisfies the builtin error interface
func (e UserInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserInfoValidationError{}

// Validate checks the field values on UpdatesUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UpdatesUserReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NickName

	// no validation rules for AvatarUrl

	return nil
}

// UpdatesUserReqValidationError is the validation error returned by
// UpdatesUserReq.Validate if the designated constraints aren't met.
type UpdatesUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatesUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatesUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatesUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatesUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatesUserReqValidationError) ErrorName() string { return "UpdatesUserReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdatesUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatesUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatesUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatesUserReqValidationError{}

// Validate checks the field values on GetByIDReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GetByIDReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetByIDReqValidationError is the validation error returned by
// GetByIDReq.Validate if the designated constraints aren't met.
type GetByIDReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetByIDReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetByIDReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetByIDReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetByIDReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetByIDReqValidationError) ErrorName() string { return "GetByIDReqValidationError" }

// Error satisfies the builtin error interface
func (e GetByIDReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetByIDReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetByIDReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetByIDReqValidationError{}

// Validate checks the field values on GetByIDResp with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetByIDResp) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetByIDRespValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetByIDRespValidationError is the validation error returned by
// GetByIDResp.Validate if the designated constraints aren't met.
type GetByIDRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetByIDRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetByIDRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetByIDRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetByIDRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetByIDRespValidationError) ErrorName() string { return "GetByIDRespValidationError" }

// Error satisfies the builtin error interface
func (e GetByIDRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetByIDResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetByIDRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetByIDRespValidationError{}

// Validate checks the field values on UserModel with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *UserModel) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	return nil
}

// UserModelValidationError is the validation error returned by
// UserModel.Validate if the designated constraints aren't met.
type UserModelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserModelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserModelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserModelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserModelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserModelValidationError) ErrorName() string { return "UserModelValidationError" }

// Error satisfies the builtin error interface
func (e UserModelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserModel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserModelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserModelValidationError{}
