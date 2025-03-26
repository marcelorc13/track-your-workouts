package models

type (
	DBResponse struct {
		Success bool
		Message string
		Data    any
	}

	HttpResponse struct {
		Status  int
		Message string
		Data    any
	}
)
