package constants

import "net/http"

var StatusCode = struct {
	INTERNAL_SERVER     int
	VALIDATION          int
	UNAUTHORIZED        int
	SUCCESS             int
	SERVICE_UNAVAILABLE int
	DOWNSTREAM          int
}{
	INTERNAL_SERVER:     http.StatusInternalServerError,
	VALIDATION:          http.StatusUnprocessableEntity,
	UNAUTHORIZED:        http.StatusUnauthorized,
	SUCCESS:             http.StatusOK,
	SERVICE_UNAVAILABLE: http.StatusServiceUnavailable,
	DOWNSTREAM:          550,
}
