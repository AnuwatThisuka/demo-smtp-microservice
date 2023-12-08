package types

type SuccessResponse struct {
	Data string `json:"data"`
}

type ErrorResponse struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}
