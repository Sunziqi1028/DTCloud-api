/**
* @Author: lik
* @Date: 2021/3/7 14:59
* @Version 1.0
 */
package bootstrap

import (
	"gitee.com/open-product/dtcloud-api/app/global/errno"
	"gitee.com/open-product/dtcloud-api/app/global/variable"
	"gitee.com/open-product/dtcloud-api/app/http/validator/common/register_validator"
	"gitee.com/open-product/dtcloud-api/app/service/sys_log_hook"
	"gitee.com/open-product/dtcloud-api/app/util/gorm_v2"
	"gitee.com/open-product/dtcloud-api/app/util/zap_factory"
	"gitee.com/open-product/dtcloud-api/config/configuration"
	"log"
	"os"
)

// 检查项目必须的非编译目录是否存在，避免编译后调用的时候缺失相关目录
func checkRequiredFolders() {
	//1.检查配置文件是否存在
	if _, err := os.Stat(variable.BasePath + "/conf/config.yaml"); err != nil {
		log.Fatal(errno.ErrorsConfigYamlNotExists + err.Error())
	}
	if _, err := os.Stat(variable.BasePath + "/conf/gorm_v2.yaml"); err != nil {
		log.Fatal(errno.ErrorsConfigGormNotExists1 + err.Error())
	}
	//2.检查public目录是否存在
	if _, err := os.Stat(variable.BasePath + "/public/"); err != nil {
		log.Fatal(errno.ErrorsPublicNotExists + err.Error())
	}
	//3.检查Storage/logs 目录是否存在
	if _, err := os.Stat(variable.BasePath + "/storage/logs/"); err != nil {
		log.Fatal(errno.ErrorsStorageLogsNotExists + err.Error())
	}
}

func init() {
	//check()
	// 1. 初始化 项目根路径，参见 variable 常量包，相关路径：app\global\variable\variable.go

	//2.检查配置文件以及日志目录等非编译性的必要条件
	checkRequiredFolders()

	//3.初始化表单参数验证器，注册在容器
	register_validator.RegisterValidator()

	// 4.启动针对配置文件(confgi.yml、gorm_v2.yaml)变化的监听， 配置文件操作指针，初始化为全局变量
	variable.ConfigYml = configuration.CreateYamlFactory()
	variable.ConfigYml.ConfigFileChangeListen()
	// config>gorm_v2.yaml 启动文件变化监听事件
	variable.ConfigGormYml = variable.ConfigYml.Clone("gorm_v2")
	variable.ConfigGormYml.ConfigFileChangeListen()

	// 5.初始化全局日志句柄，并载入日志钩子处理函数
	variable.ZapLog = zap_factory.CreateZapFactory(sys_log_hook.ZapLogHandler)

	// 6.根据配置初始化 gorm mysql 全局 *gorm.Db
	if variable.ConfigGormYml.GetInt("Gormv2.Mysql.IsInitGolobalGormMysql") == 1 {
		if dbMysql, err := gorm_v2.GetOneMysqlClient(); err != nil {
			log.Fatal(errno.ErrorsGormInitFail + err.Error())
		} else {
			variable.GormDBMysql = dbMysql
		}
	}
	// 根据配置初始化 gorm sqlserver 全局 *gorm.Db
	if variable.ConfigGormYml.GetInt("Gormv2.Sqlserver.IsInitGolobalGormSqlserver") == 1 {
		if dbSqlserver, err := gorm_v2.GetOneSqlserverClient(); err != nil {
			log.Fatal(errno.ErrorsGormInitFail + err.Error())
		} else {
			variable.GormDBSqlserver = dbSqlserver

		}
	}
	// 根据配置初始化 gorm postgresql 全局 *gorm.Db
	if variable.ConfigGormYml.GetInt("Gormv2.PostgreSql.IsInitGolobalGormPostgreSql") == 1 {
		if dbPostgre, err := gorm_v2.GetOnePostgreSqlClient(); err != nil {
			log.Fatal(errno.ErrorsGormInitFail + err.Error())
		} else {
			variable.GormDBPostgreSql = dbPostgre
		}
	}
}
