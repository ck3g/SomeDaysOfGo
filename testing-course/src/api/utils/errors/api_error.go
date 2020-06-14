package errors

// APIError represents API Error structure
type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
