package utility

type ErrorResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty" `
}

func NewErrorResponse(message, error string) ErrorResponse {
	return ErrorResponse{
		Message: message,
		Success: false,
		Error:   error,
	}
}

func NewSuccessResponse(message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Message: message,
		Success: true,
		Data:    data,
	}
}
