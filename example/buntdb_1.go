// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2020/9/1 16:04
// @Update  姓名（需要改）  2020/9/1 16:04
package main

import (
	"log"

	"github.com/tidwall/buntdb"
)

func main() {
	// Open the data.db file. It will be created if it doesn't exist.
	db, err := buntdb.Open("my_buntdb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
