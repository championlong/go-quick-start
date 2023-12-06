package config

import (
	"encoding/json"

	"github.com/championlong/go-quick-start/internal/pkg/options"
	"github.com/championlong/go-quick-start/pkg/log"
	"github.com/championlong/robot-talk-sdk/config"
)

type Server struct {
	JWT     options.JWT    `mapstructure:"jwt"         json:"jwt"         yaml:"jwt"`
	Zap     log.ZapOptions `mapstructure:"zap"         json:"zap"         yaml:"zap"`
	Redis   options.Redis  `mapstructure:"redis"       json:"redis"       yaml:"redis"`
	Email   options.Email  `mapstructure:"email"       json:"email"       yaml:"email"`
	Casbin  Casbin         `mapstructure:"casbin"      json:"casbin"      yaml:"casbin"`
	System  System         `mapstructure:"system"      json:"system"      yaml:"system"`
	Captcha Captcha        `mapstructure:"captcha"     json:"captcha"     yaml:"captcha"`
	// auto
	AutoCode Autocode `mapstructure:"autocode"    json:"autocode"    yaml:"autocode"`
	// gorm
	Mysql  options.Mysql           `mapstructure:"mysql"       json:"mysql"       yaml:"mysql"`
	Pgsql  options.Pgsql           `mapstructure:"pgsql"       json:"pgsql"       yaml:"pgsql"`
	Oracle options.Oracle          `mapstructure:"oracle"      json:"oracle"      yaml:"oracle"`
	Sqlite options.Sqlite          `mapstructure:"sqlite"      json:"sqlite"      yaml:"sqlite"`
	DBList []options.SpecializedDB `mapstructure:"db-list"     json:"db-list"     yaml:"db-list"`
	// oss
	Local      options.Local      `mapstructure:"local"       json:"local"       yaml:"local"`
	Qiniu      options.Qiniu      `mapstructure:"qiniu"       json:"qiniu"       yaml:"qiniu"`
	AliyunOSS  options.AliyunOSS  `mapstructure:"aliyun-oss"  json:"aliyun-oss"  yaml:"aliyun-oss"`
	HuaWeiObs  options.HuaWeiObs  `mapstructure:"hua-wei-obs" json:"hua-wei-obs" yaml:"hua-wei-obs"`
	TencentCOS options.TencentCOS `mapstructure:"tencent-cos" json:"tencent-cos" yaml:"tencent-cos"`
	AwsS3      options.AwsS3      `mapstructure:"aws-s3"      json:"aws-s3"      yaml:"aws-s3"`

	Excel Excel         `mapstructure:"excel" json:"excel" yaml:"excel"`
	Timer options.Timer `mapstructure:"timer" json:"timer" yaml:"timer"`

	Dingding config.PlatformConfig `mapstructure:"dingding" json:"dingding" yaml:"dingding"`
	// 跨域配置
	Cors options.CORS `mapstructure:"cors"     json:"cors"     yaml:"cors"`
}

func (o *Server) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}
