package exceptions

type HttpException struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
