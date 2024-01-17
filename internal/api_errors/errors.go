package api_errors

import "net/http"

var (
	ErrInternalServerError = "10000"
	ErrUnauthorizedAccess  = "10001"
	ErrInvalidUserID       = "10002"
	ErrValidation          = "10003"
	ErrDeleteFailed        = "10004"

	// Service errors begin with 30000
)

type MessageAndStatus struct {
	Message string
	Status  int
}

var MapErrorCodeMessage = map[string]MessageAndStatus{
	ErrInternalServerError: {"Internal Server Error", http.StatusInternalServerError},
	ErrUnauthorizedAccess:  {"Unauthorized Access", http.StatusUnauthorized},
	ErrInvalidUserID:       {"Invalid User ID", http.StatusBadRequest},
	ErrValidation:          {"Validation Error", http.StatusBadRequest},
	ErrDeleteFailed:        {"Delete Failed", http.StatusInternalServerError},
}
