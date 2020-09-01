// @Title  commUtil.go
// @Description  包含常用的方法的工具类
// @Author  zhengbin  2020/7/23 19:17
// @Update  姓名（需要改）  2020/7/23 19:17
package utils

import (
	"bytes"
	"encoding/binary"
	"time"
)

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

//字节转换成整形
func BytesToInt32(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return x
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func Int32ToBytes(i int32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// MIN 为用户自定义的比较精度
func IsEqual(f1, f2 float64) bool {
	if f1 > f2 {
		return f1-f2 < 0.0000001
	} else {
		return f2-f1 < 0.0000001
	}
}

//获取当前月份，返回两位字符串，不足两位用0补齐
func GetCurrentMonth() string {
	d := time.Now().Month()
	if d > 9 {
		return string(d)
	} else {
		return "0" + string(d)
	}
}

// 为实现 int64 的切片排序，新增一个结构体，并实现排序接口所需方法
// 将 []int64 定义为 SortableInt64List 类型
type SortableInt64List []int64

// 实现 sort.Interface 接口的获取元素数量方法
func (m SortableInt64List) Len() int {
	return len(m)
}

// 实现 sort.Interface 接口的比较元素方法
func (m SortableInt64List) Less(i, j int) bool {
	return m[i] < m[j]
}

// 实现 sort.Interface 接口的交换元素方法
func (m SortableInt64List) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
