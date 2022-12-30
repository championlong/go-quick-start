package model

// DingDingResponse 官方钉钉响应体
type DingDingResponse struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
}

type At struct {
	IsAtAll   bool     `json:"isAtAll,omitempty"`   //@所有人
	AtMobiles []string `json:"atMobiles,omitempty"` //@手机号
	AtUsersId []string `json:"atUserIds,omitempty"` //@userid
}
