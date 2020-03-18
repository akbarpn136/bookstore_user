package errors

import "net/http"

type RestErrors struct {
	Message []string `json:"message"`
	Code    int      `json:"code"`
	Error   string   `json:"error"`
}

func BadRequest(message []string) *RestErrors {
	return &RestErrors{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NotFound(message []string) *RestErrors {
	return &RestErrors{
		Message: message,
		Code:    http.StatusNotFound,
		Error:   "not_found",
	}
}

func NotAcceptable(message []string) *RestErrors {
	return &RestErrors{
		Message: message,
		Code:    http.StatusNotAcceptable,
		Error:   "not_acceptable",
	}
}
