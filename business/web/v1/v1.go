//  Package v1 represents types used by the web application for v1.
package v1

import "errors"

// ErrorResponse is the form used for API responses from failures in the API.
type ErrorResponse struct {
	Error  string `json:"error"`
	Fields string `json:"fields,omitempty"`
}

// RequestError is used to pass an error during the request through the
// application with web specific context.
type RequestError struct {
	Err    error
	Status int
}

// NewRequestError wraps a provided error with an HTTP status code. This
// function should be used when handlers encounter expected errors.
func NewRequestError(err error, status int) error {
	return &RequestError{err, status}
}

// Error implements the error interface. It uses the default message of the
// wrapped error. This is what will be shown in the services' logs.
func (err *RequestError) Error() string {
	return err.Err.Error()
}

// AsRequestError checks if an error of type RequestError exists.
func AsRequestError(err error) bool {
	var requestErrors *RequestError
	return errors.As(err, &requestErrors)
}

// GetRequestError returns a copy of the RequestError pointer.
func GetRequestError(err error) *RequestError {
	var requestError *RequestError
	if !errors.As(err, &requestError) {
		return nil
	}
	return requestError
}
