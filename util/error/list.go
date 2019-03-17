package error

import "net/http"

// Unexpected should be used to indicate generic unexpected error
func Unexpected() Container {
	return Error{
		Status: http.StatusInternalServerError,
		Code:   "UPWK-9999",
		Detail: "Unexpected error",
	}.ProduceStackTrace().ToContainer()
}

// NotFound should be used to indicate not found route error.
func NotFound() Container {
	return Error{
		Status: http.StatusNotFound,
		Code:   "UPWK-1234",
		Detail: "Route not found",
	}.ProduceStackTrace().ToContainer()
}

// UnexpectedPanic should be used to indicate panic
func UnexpectedPanic() Container {
	return Error{
		Status: http.StatusInternalServerError,
		Code:   "USR-10000", // You know it gets real when the error code is not 4 digit anymore
		Detail: "Unexpected panic",
	}.ProduceStackTrace().ToContainer()
}
