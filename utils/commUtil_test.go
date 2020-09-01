package utils

import (
	"fmt"
	"testing"
	"time"
)

type TestStruct struct {
	id    int
	name  string
	value string
}

func TestIsEqual(t *testing.T) {

	var a, b float64
	a = 0.1
	b = 0.1
	fmt.Printf("a:[%f],b:[%f],a==b:[%v]\n", a, b, IsEqual(a, b))
	fmt.Printf("f:[%d]", int(time.Now().Month()))

	t1 := TestStruct{1, "name", "value"}

	fmt.Printf("old value[%s]\n", t1.value)
	ModifyTestStruct(&t1)
	fmt.Printf("new value[%s]\n", t1.value)

	bb := []string{"test"}
	aa := []int{1}
	ModifyTestSlice(bb, aa)
	fmt.Printf("a[%s]\n", bb)
	fmt.Printf("a[%d]\n", aa)

}

func ModifyTestStruct(a *TestStruct) {
	a.value = "new value"
}

func ModifyTestSlice(a []string, b []int) {
	a = append(a, "test0")
	a = append(a, "test1")
	b = append(b, 1)
	b = append(b, 2)
}
