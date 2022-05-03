package model

// ActionCardBtn 卡片类型
type ActionCardBtn struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

// DingDingResponse 官方钉钉响应体
type DingDingResponse struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
}
