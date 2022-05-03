package model

// DingdingConfig 钉钉配置
type DingdingConfig struct {
	// WebHook 钉钉官方配置
	WebHook DingdingKindConfig `mapstructure:"web-hook" json:"web-hook" yaml:"web-hook"`
}

// DingdingKindConfig 钉钉官方配置
type DingdingKindConfig struct {
	// Url 钉钉发送URl
	Url string `mapstructure:"url" json:"url" yaml:"url"`
	// DingdingQuery 机器人列表
	DingdingQuery map[string]DingdingQueryConfig `mapstructure:"dingding-query" json:"dingding-query" yaml:"dingding-query"`
}

// DingdingQueryConfig 钉钉机器人配置
type DingdingQueryConfig struct {
	// IsKeysWord 是否使用关键字发送
	IsKeysWord bool `mapstructure:"is-keys-word" json:"is-keys-word" yaml:"is-keys-word"`
	// KeysWord 关键字
	KeysWord string `mapstructure:"keys-word" json:"keys-word" yaml:"keys-word"`
	// IsEncrypt 是否使用加签发送
	IsEncrypt bool `mapstructure:"is-encrypt" json:"is-encrypt" yaml:"is-encrypt"`
	// Encrypt 加签信息
	Encrypt string `mapstructure:"encrypt" json:"encrypt" yaml:"encrypt"`
	// AccessToken 请求token信息
	AccessToken string `mapstructure:"access-token" json:"access-token" yaml:"access-token"`
}

// UrlConfig 基础的url配置，包含url、query、params
type UrlConfig struct {
	// Url 公司钉钉发送URL地址
	Url string `mapstructure:"url" json:"url" yaml:"url"`
	// Query 请求参数
	Query map[string]string `mapstructure:"query" json:"query" yaml:"query"`
	// Params 参数
	Params map[string]string `mapstructure:"params" json:"params" yaml:"params"`
}
