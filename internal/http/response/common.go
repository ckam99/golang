package response

type ValidationError struct {
	Namespace string `json:"namespace"`
	Field     string `json:"field"`
	Tag       string `json:"tag"`
	Value     string `json:"value"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
