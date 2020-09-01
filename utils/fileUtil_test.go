// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2020/8/26 15:37
// @Update  姓名（需要改）  2020/8/26 15:37
package utils

import (
	"fmt"
	"testing"
)

func TestReadLastLineFromFile(t *testing.T) {
	filePath := "C:/tmp/sparrow_redo/Index.redo"
	lastLine, _ := ReadLastLineFromFile(filePath, 100, true)
	fmt.Println("最后一行==>", lastLine)
}
