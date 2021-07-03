/**
* @Author: lik
* @Date: 2021/3/5 19:09
* @Version 1.0
 */
package model

import (
	"fmt"
	"gitee.com/open-product/dtcloud-api/app/global/errno"
	"gitee.com/open-product/dtcloud-api/app/global/variable"
	"gorm.io/gorm"
	"strings"
	"time"
)

type BaseModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	Id        int64     `gorm:"primary key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func useDbConn(sqlType string) *gorm.DB {
	var db *gorm.DB
	sqlType = strings.Trim(sqlType, " ")
	if sqlType == "" {
		sqlType = variable.ConfigGormYml.GetString("Gormv2.UseDbType")
	}
	switch strings.ToLower(sqlType) {
	case "mysql":
		if variable.GormDBMysql == nil {
			variable.ZapLog.Fatal(fmt.Sprintf(errno.ErrorsGormNotInitGlobalPointer, sqlType, sqlType))
		}
		db = variable.GormDBMysql
	case "sqlserver":
		if variable.GormDBSqlserver == nil {
			variable.ZapLog.Fatal(fmt.Sprintf(errno.ErrorsGormNotInitGlobalPointer, sqlType, sqlType))
		}
		db = variable.GormDBSqlserver
	case "postgres", "postgre", "postgresql":
		if variable.GormDBPostgreSql == nil {
			variable.ZapLog.Fatal(fmt.Sprintf(errno.ErrorsGormNotInitGlobalPointer, sqlType, sqlType))
		}
		db = variable.GormDBPostgreSql

	default:
		variable.ZapLog.Error(errno.ErrorsDbDriverNotExists + sqlType)
	}
	return db
}
