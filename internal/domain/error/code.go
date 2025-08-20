package error

import "net/http"

const (
	// Client errors (4xx)
	CodeBadRequest   = "BAD_REQUEST"
	CodeUnauthorized = "UNAUTHORIZED"
	CodeForbidden    = "FORBIDDEN"
	CodeNotFound     = "NOT_FOUND"
	CodeConflict     = "CONFLICT"

	// Server errors (5xx)
	CodeInternalServer = "INTERNAL_SERVER_ERROR"
	CodeDatabaseError  = "DATABASE_ERROR"
	CodeTimeout        = "TIMEOUT"
)

var ErrorCodeToStatus = map[string]int{
	CodeBadRequest:     http.StatusBadRequest,
	CodeNotFound:       http.StatusNotFound,
	CodeConflict:       http.StatusConflict,
	CodeUnauthorized:   http.StatusUnauthorized,
	CodeForbidden:      http.StatusForbidden,
	CodeInternalServer: http.StatusInternalServerError,
	CodeTimeout:        http.StatusGatewayTimeout,
}
