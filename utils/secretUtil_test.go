// @Title  secretUtil_test.go
// @Description  测试类
// @Author  zhengbin  2020/7/23 17:03
// @Update  姓名（需要改）  2020/7/23 17:03
package utils

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAESEncrypt(t *testing.T) {
	key := []byte("lifeisshortandartislong.")

	result, _ := AESEncrypt([]byte("admin"), key)
	fmt.Println(base64.StdEncoding.EncodeToString(result))
}
