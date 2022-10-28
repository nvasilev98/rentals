package api

type Error struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error: Error{message},
	}
}
