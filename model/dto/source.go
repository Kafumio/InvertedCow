package dto

type Source struct {
}

type Token struct {
	Token       string `json:"token"`
	OriginUrl   string `json:"origin_id"`    // 关联发布动态id
	CallbackUrl string `json:"callback_url"` // 回调
}
