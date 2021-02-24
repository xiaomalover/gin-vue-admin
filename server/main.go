package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.GvaVp = core.Viper()          // 初始化Viper
	global.GvaLog = core.Zap()           // 初始化zap日志库
	global.GvaDb = initialize.Gorm()     // gorm连接数据库
	initialize.MysqlTables(global.GvaDb) // 初始化表
	// 程序结束前关闭数据库链接
	db, _ := global.GvaDb.DB()
	defer db.Close()

	core.RunWindowsServer()
}
