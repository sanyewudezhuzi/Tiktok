package serializer

type Response struct {
	StatusCode int         `json:"status_code"` // 0-成功，其他-失败
	StatusMsg  string      `json:"status_msg"`
	Data       interface{} `json:"data"`
}
