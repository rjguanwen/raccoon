// @Title  fileUtil.go
// @Description  文件操作工具类
// @Author  zhengbin  2020/7/22 16:00
// @Update  姓名（需要改）  ${DATE} ${TIME}
package utils

import (
	log "github.com/cihub/seelog"
	"os"
	"path"
)

// @title    AbsPath
// @description   获取绝对路径，与入口程序所在目录相关，传入的文件相对路径参数为相对入口程序位置的路径
// @auth      zhengbin  2020/7/22 16:00
// @param     relFilepath        string         "文件相对路径"
// @return    absFilepath        string         "文件绝对路径"
func AbsPath(relFilepath string) string {
	// 获取绝对路径
	rootDir, err := os.Getwd()
	log.Info("rootDir:", rootDir)
	if err != nil {
		log.Error("获取文件绝对路径失败：", err)
	}
	absPath := path.Join(rootDir, relFilepath)
	log.Info("absPath:", absPath)
	return absPath
}

// @title    IsFileExist
// @description   判断文件是否存在
// @auth      zhengbin  2020/7/22 16:00
// @param     path        string         "文件路径"
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
