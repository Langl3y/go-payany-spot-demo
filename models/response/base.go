package response

type LogInResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    LogInResponseData `json:"data"`
}
