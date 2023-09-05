package util

type ApiResponse struct {
	success bool
	data    interface{}
	message string
}
type ApiError struct {
	ApiResponse
	code string
}

// Create a new api error instance
func NewError(message string, code string) ApiError {
	err := ApiError{
		ApiResponse: ApiResponse{
			message: message,
			success: false,
		},
		code: code,
	}
	return err
}

type ApiSuccess struct {
	ApiResponse
}

// Create a new api success instance
func NewSuccess(message string, data interface{}) ApiSuccess {
	res := ApiSuccess{
		ApiResponse: ApiResponse{
			message: message,
			success: true,
			data:    data,
		},
	}
	return res
}
