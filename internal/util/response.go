package util

import (
	derror "clean-architecture/internal/domain/error"
	"clean-architecture/internal/infrastructure/repository"
	"context"
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func NewErrorResponse(err error) (*ErrorResponse, int) {
	if err == nil {
		return &ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Unknown Error",
			Detail:  "No error provided",
		}, http.StatusInternalServerError
	}

	if status, ok := internalError(err); ok {
		return newErrorResponse(status, err), status
	}

	// Default to internal server error
	return &ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: "Internal Server Error",
		Detail:  err.Error(),
	}, http.StatusInternalServerError
}

func newErrorResponse(status int, err error) *ErrorResponse {
	return &ErrorResponse{
		Status:  status,
		Message: http.StatusText(status),
		Detail:  err.Error(),
	}
}

func internalError(err error) (int, bool) {
	if err == nil {
		return 0, false
	}

	// First, check domain errors
	if status, ok := internalDomainError(err); ok {
		return status, true
	}

	// Then, check standard errors
	var status int

	switch {
	// 4xx - Client Errors
	case // 400
		errors.Is(err, ErrInvalidArgument),
		errors.Is(err, repository.ErrInvalidArgument):
		status = http.StatusBadRequest
	case // 401
		errors.Is(err, ErrUnauthorized):
		status = http.StatusUnauthorized
	case // 403
		errors.Is(err, ErrForbidden):
		status = http.StatusForbidden
	case // 404
		errors.Is(err, ErrNotFound),
		errors.Is(err, repository.ErrNotFound):
		status = http.StatusNotFound
	case // 406
		errors.Is(err, ErrNotAcceptable),
		errors.Is(err, repository.ErrNotAcceptable):
		status = http.StatusNotAcceptable
	case // 409
		errors.Is(err, ErrAlreadyExists),
		errors.Is(err, repository.ErrAlreadyExists):
		status = http.StatusConflict
	case // 412
		errors.Is(err, ErrFailedPrecondition),
		errors.Is(err, repository.ErrFailedPrecondition):
		status = http.StatusPreconditionFailed
	case // 429
		errors.Is(err, ErrResourceExhausted):
		status = http.StatusTooManyRequests

	// 5xx - Server Errors
	case // 500
		errors.Is(err, ErrInternal),
		errors.Is(err, repository.ErrInternal):
		status = http.StatusInternalServerError
	case // 501
		errors.Is(err, ErrNotImplemented),
		errors.Is(err, repository.ErrNotImplemented):
		status = http.StatusNotImplemented
	case // 502
		errors.Is(err, ErrUnavailable):
		status = http.StatusBadGateway
	case // 504
		errors.Is(err, ErrDeadlineExceeded),
		errors.Is(err, context.Canceled),
		errors.Is(err, context.DeadlineExceeded),
		errors.Is(err, repository.ErrCanceled),
		errors.Is(err, repository.ErrDeadlineExceeded):
		status = http.StatusGatewayTimeout
	default:
		return 0, false
	}

	return status, true
}

func internalDomainError(err error) (int, bool) {
	var de *derror.DomainError
	if errors.As(err, &de) {
		return de.Status, true
	}
	return 0, false
}
