// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2020/9/1 15:25
// @Update  姓名（需要改）  2020/9/1 15:25
package main

import (
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my_bolt.db", 0600, nil)
	//db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
