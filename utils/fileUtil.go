// @Title  fileUtil.go
// @Description  文件操作工具类
// @Author  zhengbin  2020/7/22 16:00
// @Update  姓名（需要改）  ${DATE} ${TIME}
package utils

import (
	"bufio"
	log "github.com/cihub/seelog"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
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

//
// @title    	GetAllFileFromPathWithoutSubDir
// @description 获取文件夹下的所有文件，不处理子文件夹
// @auth      	zhengbin  2020/8/25 13:00
// @param 		pathname        string         "文件夹路径"
// @return 	  	fileArr			[]string		"文件名数组"
// @return 		err				error			"错误信息"
func GetAllFileFromPathWithoutSubDir(pathname string) (fileArr []string, err error) {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			// 如果是目录，则跳过
			continue
		} else {
			// 如果不是目录，则加入到返回列表
			fileArr = append(fileArr, fi.Name())
		}
	}
	return fileArr, err
}

//// 判断所给路径是否为文件夹
//func IsDir(path string) bool {
//	s, err := os.Stat(path)
//	if err != nil {
//		return false
//	}
//	return s.IsDir()
//}
//
//// 判断所给路径是否为文件
//func IsFile(path string) bool {
//	return !IsDir(path)
//}

// @title    	ReadLastLine
// @description 读取指定文件的最后一行
// @auth      	zhengbin  2020/8/26 15:04
// @param 		filePath        string		"文件路径"
// @param 	  	buffSize		int			"长度，该参数非常重要，必须保证这个数值大于最后一行的长度"
// @param 		exceptEmpty		bool		"是否排除空行"
// @return
// @return 		err				error		"错误信息"
func ReadLastLineFromFile(filePath string, buffSize int, exceptEmpty bool) (lastLine string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("fileUtil - ReadLastLine - 文件打开错误：%v", err)
		panic(err)
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		log.Errorf("fileUtil - ReadLastLine - 文件属性获取错误：%v", err)
	}
	// 读取的开始位置
	var startPosition int64
	if int64(buffSize) > fi.Size() {
		startPosition = 0
	} else {
		startPosition = fi.Size() - int64(buffSize)
	}
	//fmt.Println("startPosition======>", startPosition)
	// 将文件定位到开始读取位置
	file.Seek(startPosition, os.SEEK_SET)

	var line []byte
	reader := bufio.NewReader(file)
	for {
		tempLine, err := reader.ReadBytes('\n')
		//fmt.Println("::", line)
		if err == io.EOF {
			log.Debug("fileUtil - ReadLastLine - 已达到文档末尾！")
			break
		} else if err != nil {
			log.Errorf("fileUtil - ReadLastLine - 文件读取错误：%v", err)
			panic(err)
		}
		if exceptEmpty {
			if len(tempLine) != 0 && string(tempLine) != "\n" {
				line = tempLine
			}
		} else {
			line = tempLine
		}
	}
	//fmt.Println("---==>>>", strings.TrimSpace(string(line)))
	return strings.TrimSpace(string(line)), err
}
