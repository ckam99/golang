package response

type ErrorResponse struct {
	Namespace string `json:"namespace"`
	Field     string `json:"field"`
	Tag       string `json:"tag"`
	Value     string `json:"value"`
}
