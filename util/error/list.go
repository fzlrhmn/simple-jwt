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
		Code:   "UPWK-10000", // You know it gets real when the error code is not 4 digit anymore
		Detail: "Unexpected panic",
	}.ProduceStackTrace().ToContainer()
}

// RequestBodyUnparseable should be used to indicate that the request body is unparseable
// UPWK-1000
func RequestBodyUnparseable() Container {
	return Error{
		Status: http.StatusBadRequest,
		Code:   "UPWK-1000",
		Detail: "Failed parsing json body",
	}.ProduceStackTrace().ToContainer()
}

// RequestBodyFailsValidation should be used to indicate that the request body doesn't pass validation
// This error is expecting error that comes from github.com/go-playground/validator.
// UPWK-1001
func RequestBodyFailsValidation(e error) Container {
	return Error{
		Status: http.StatusBadRequest,
		Code:   "UPWK-1001",
	}.ProduceStackTrace().ToContainer().Wrap(e)
}

// MissingPrimaryID should be used to indicate that the request
// is missing primary id
func MissingPrimaryID() Container {
	return Error{
		Status: http.StatusBadRequest,
		Code:   "UPWK-1003",
		Detail: "Missing username",
	}.ProduceStackTrace().ToContainer()
}
