package api

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(data interface{}) Response {
	return Response{
		Code:    0,
		Message: "ok",
		Data:    data,
	}
}

func FailResponse(code int64, message string) Response {
	return Response{
		Code:    code,
		Message: message,
	}
}
