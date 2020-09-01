// @Title  server.go
// @Description  服务器启动入口
// @Author  zhengbin  2020/8/5 17:08
// @Update  姓名（需要改）  2020/8/5 17:08
package main

import (
	log "github.com/cihub/seelog"
)

// 初始化
func init() {
	// 初始化日志配置
	SetupLogger()
}

func main() {
	log.Info("Hello, Raccoon!")
}

// 设置日志功能配置
func SetupLogger() {
	logger, err := log.LoggerFromConfigAsFile("config/seelog.xml")
	if err != nil {
		return
	}

	log.ReplaceLogger(logger)
}
