package system

import (
	"database/sql"
	"fmt"

	"github.com/championlong/backend-common/app/global"
	"github.com/championlong/backend-common/app/model/system"
	"github.com/championlong/backend-common/app/model/system/request"
)

type InitDBService struct{}

// InitDB 创建数据库并初始化 总入口
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (initDBService *InitDBService) InitDB(conf request.InitDB) error {
	switch conf.DBType {
	case "mysql":
		return initDBService.initMysqlDB(conf)
	case "pgsql":
		return initDBService.initPgsqlDB(conf)
	default:
		return initDBService.initMysqlDB(conf)
	}
}

// initTables 初始化表
// Author [SliverHorn](https://github.com/SliverHorn)
func (initDBService *InitDBService) initTables() error {
	return global.GVA_DB.AutoMigrate(
		system.SysUser{},
		system.SysAuthority{},
		system.JwtBlacklist{},
	)
}

// createDatabase 创建数据库(mysql)
// Author [SliverHorn](https://github.com/SliverHorn)
// Author: [songzhibin97](https://github.com/songzhibin97)

func (initDBService *InitDBService) createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}
