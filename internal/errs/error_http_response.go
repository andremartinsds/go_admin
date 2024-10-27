package errs

type HttpResponse struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}
