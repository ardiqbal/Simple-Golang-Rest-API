package util

type PostHttpResponse struct {
	Success bool     `json:"success"`
	Message []string `json:"message"`
}

type GetHttpResponse struct {
	Success bool        `json:"success"`
	Message []string    `json:"message"`
	Data    interface{} `json:"data"`
}
