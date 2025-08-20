package error

import (
	"fmt"
	"time"
)

func New(code, message string, status int) *DomainError {
	return &DomainError{
		Code:      code,
		Status:    status,
		Message:   message,
		Stack:     captureStackTrace(1),
		Timestamp: time.Now(),
	}
}

func Newf(code, format string, status int, args ...interface{}) *DomainError {
	return &DomainError{
		Code:      code,
		Status:    status,
		Message:   fmt.Sprintf(format, args...),
		Stack:     captureStackTrace(1),
		Timestamp: time.Now(),
	}
}

func WithDetails(code, message string, status int, details map[string]interface{}) *DomainError {
	return &DomainError{
		Code:      code,
		Status:    status,
		Message:   message,
		Details:   details,
		Stack:     captureStackTrace(1),
		Timestamp: time.Now().UTC(),
	}
}

func Wrap(cause error, code, message string, status int) *DomainError {
	return &DomainError{
		Code:      code,
		Status:    status,
		Message:   message,
		Cause:     cause,
		Stack:     captureStackTrace(1),
		Timestamp: time.Now().UTC(),
	}
}

func Wrapf(cause error, code, format string, status int, args ...interface{}) *DomainError {
	return &DomainError{
		Code:      code,
		Status:    status,
		Message:   fmt.Sprintf(format, args...),
		Cause:     cause,
		Stack:     captureStackTrace(1),
		Timestamp: time.Now().UTC(),
	}
}

func BadRequest(message string) *DomainError {
	return New(
		CodeBadRequest,
		message,
		ErrorCodeToStatus[CodeBadRequest],
	)
}

func BadRequestf(format string, args ...interface{}) *DomainError {
	return Newf(
		CodeBadRequest,
		format,
		ErrorCodeToStatus[CodeBadRequest],
		args...,
	)
}

func Unauthorized(message string) *DomainError {
	return New(
		CodeUnauthorized,
		message,
		ErrorCodeToStatus[CodeUnauthorized],
	)
}

func Forbidden(message string) *DomainError {
	return New(
		CodeForbidden,
		message,
		ErrorCodeToStatus[CodeForbidden],
	)
}

func NotFound(resource string) *DomainError {
	return Newf(
		CodeNotFound,
		"%s not found",
		ErrorCodeToStatus[CodeNotFound],
		resource,
	)
}

func Conflict(message string) *DomainError {
	return New(
		CodeConflict,
		message,
		ErrorCodeToStatus[CodeConflict],
	)
}

func InternalServer(cause error) *DomainError {
	return New(
		CodeInternalServer,
		cause.Error(),
		ErrorCodeToStatus[CodeInternalServer],
	)
}

func InternalServerf(format string, args ...interface{}) *DomainError {
	return Newf(
		CodeInternalServer,
		format,
		ErrorCodeToStatus[CodeInternalServer],
		args...,
	)
}

func DatabaseError(cause error) *DomainError {
	return Wrap(
		cause,
		CodeDatabaseError,
		"Database operation failed",
		ErrorCodeToStatus[CodeDatabaseError],
	)
}

func Timeout(message string) *DomainError {
	return New(
		CodeTimeout,
		message,
		ErrorCodeToStatus[CodeTimeout],
	)
}

func Timeoutf(format string, args ...interface{}) *DomainError {
	return Newf(
		CodeTimeout,
		format,
		ErrorCodeToStatus[CodeTimeout],
		args...,
	)
}
