package linkedlist

import (
	"testing"
	"fmt"
)

//	测试 New
func TestLinkedList(t *testing.T) {
	list:= New()
	println("通过 New() 实例化: list")
	println("测试 GetSize(), 预计： 0")
	size:= list.GetSize()
	fmt.Printf("测试 GetSize(), 实际： %d\n", size)
	if (size != 0) {
		t.Error(" GetSize() 方式失败")
	}
}