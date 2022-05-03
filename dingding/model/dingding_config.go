package model

// DingdingConfig 钉钉配置
type DingdingConfig struct {
	// WebHook 钉钉官方配置
	WebHook DingdingKindConfig
}

// DingdingKindConfig 钉钉官方配置
type DingdingKindConfig struct {
	// Url 钉钉发送URl
	Url string
	// DingdingQuery 机器人列表
	DingdingQuery map[string]DingdingQueryConfig
}

// DingdingQueryConfig 钉钉机器人配置
type DingdingQueryConfig struct {
	// IsKeysWord 是否使用关键字发送
	IsKeysWord bool
	// KeysWord 关键字
	KeysWord string
	// IsEncrypt 是否使用加签发送
	IsEncrypt bool
	// Encrypt 加签信息
	Encrypt string
	// AccessToken 请求token信息
	AccessToken string
}

// UrlConfig 基础的url配置，包含url、query、params
type UrlConfig struct {
	// Url 公司钉钉发送URL地址
	Url string
	// Query 请求参数：user_name(申请的服务名称)、access_token(申请的token)
	Query  map[string]string
	Params map[string]string
}
