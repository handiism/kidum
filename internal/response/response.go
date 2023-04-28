package response

type Status string

const (
	StatusFail    Status = "Fail"
	StatusSuccess Status = "Success"
	StatusError   Status = "Error"
)

type Response struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(message string, data any) Response {
	return Response{
		Status:  StatusSuccess,
		Message: message,
		Data:    data,
	}
}

func Fail(message string) Response {
	return Response{
		Status:  StatusFail,
		Message: message,
		Data:    nil,
	}
}

func Error(message string) Response {
	return Response{
		Status:  StatusError,
		Message: message,
		Data:    nil,
	}
}
