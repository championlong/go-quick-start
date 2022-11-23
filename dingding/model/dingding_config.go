package model

// DingdingKindConfig 钉钉官方配置
type DingdingConfig struct {
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