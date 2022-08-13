package dto

type HttpError struct {
	Message string `json:"message"`
}

type ValidateError struct {
	Namespace string `json:"namespace"`
	Field     string `json:"field"`
	Tag       string `json:"tag"`
	Value     string `json:"value"`
}
