package dto

type Source struct {
	Id     string `json:"_id"`
	Key    string `json:"key"`
	Hash   string `json:"hash"`
	FSize  int64  `json:"fsize"`
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

type Token struct {
	Token       string `json:"token"`
	OriginUrl   string `json:"origin_id"`    // 关联发布动态id
	CallbackUrl string `json:"callback_url"` // 回调
}
