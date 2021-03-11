package handler

// HTTPError error struct
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

func (h HTTPError) Error() string {
	return h.Message
}

//NewHTTPError returns new error
func NewHTTPError(code int, message string) HTTPError {
	return HTTPError{
		Code:    code,
		Message: message,
	}
}
