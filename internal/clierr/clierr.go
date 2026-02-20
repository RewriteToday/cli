package clierr

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
)

type Code int

const (
	CodeSuccess      Code = 0
	CodeUsage        Code = 2
	CodeAuthRequired Code = 3
	CodeNotFound     Code = 4
	CodeNetwork      Code = 5
	CodeInternal     Code = 10
)

type Error struct {
	Code Code
	Err  error
}

func (e *Error) Error() string {
	if e == nil || e.Err == nil {
		return "unknown error"
	}

	return e.Err.Error()
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Err
}

func (c Code) String() string {
	switch c {
	case CodeSuccess:
		return "success"
	case CodeUsage:
		return "usage_error"
	case CodeAuthRequired:
		return "auth_required"
	case CodeNotFound:
		return "not_found"
	case CodeNetwork:
		return "network_error"
	default:
		return "internal_error"
	}
}

func Wrap(code Code, err error) error {
	if err == nil {
		return nil
	}

	var coded *Error

	if errors.As(err, &coded) {
		return err
	}

	return &Error{
		Code: code,
		Err:  err,
	}
}

func Errorf(code Code, format string, args ...any) error {
	return &Error{
		Code: code,
		Err:  fmt.Errorf(format, args...),
	}
}

func CodeOf(err error) Code {
	if err == nil {
		return CodeSuccess
	}

	var coded *Error
	if errors.As(err, &coded) {
		return coded.Code
	}

	if isNetworkError(err) {
		return CodeNetwork
	}

	return CodeInternal
}

func ExitCode(err error) int {
	return int(CodeOf(err))
}

func isNetworkError(err error) bool {
	if err == nil {
		return false
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		return true
	}

	var opErr *net.OpError
	if errors.As(err, &opErr) {
		return true
	}

	var urlErr *url.Error
	return errors.As(err, &urlErr)
}
