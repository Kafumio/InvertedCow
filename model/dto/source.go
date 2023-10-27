package dto

// Source 七牛云 业务回调请求参数
type Source struct {
	Id     string `json:"_id"`
	Key    string `json:"key"`
	Hash   string `json:"hash"`
	FSize  int64  `json:"fsize"`
	Bucket string `json:"bucket"`
	PID    string `json:"pid"`
}

// Token 七牛云授权Token
type Token struct {
	Token string `json:"token"`
	PID   uint   `json:"pid"` // 关联发布动态id
}
