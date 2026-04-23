package utils

type AppError struct {
	Message string
	Code int
}

func NewError(message string , code int) *AppError{
	return &AppError{
		Message: message,
		Code: code,
	}
}