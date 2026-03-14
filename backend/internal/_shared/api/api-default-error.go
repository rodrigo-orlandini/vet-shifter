package api

type ApiErrorResponse struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}
