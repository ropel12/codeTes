package domain

type ResponseApi struct {
	StatusCode string      `json:"status_code"`
	Data       interface{} `json:"data"`
}
