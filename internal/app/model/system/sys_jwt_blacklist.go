package system

import (
	"github.com/championlong/go-quick-start/internal/app/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
