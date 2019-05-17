package linkedlist

import (
	"testing"
	"fmt"
)

var list LinkedLister

//	测试 New
func TestLinkedList(t *testing.T) {
	list = New()
	println("通过 New() 实例化: list")
	println("测试 GetSize(), 预计： 0")
	size:= list.GetSize()
	fmt.Printf("测试 GetSize(), 实际： %d\n", size)
	if (size != 0) {
		t.Fatal("GetSize() ====失败====")
	}
	
	isEmpty:= list.IsEmpty()
	fmt.Println("测试 IsEmpty()， 预计： true")
	fmt.Println("测试 IsEmpty()， 实际： ", isEmpty)
	if !isEmpty {
		t.Fatal("IsEmpty() ====失败====")
	}

}

func TestAppend(t *testing.T) {
	node:= *NewNode()
	node.Value = 2

	list.Append(node)
	size:= list.GetSize()
	fmt.Println("空的 list Append 了一个结点 Value: 2")
	fmt.Println("list IsEmpty(): 应为 false, 实际：", list.IsEmpty())
	fmt.Println("list GetSize(): 应为 1， 实际：", size)
	if size != 1 {
		t.Fatal("GetSize() ====失败====")
	}

	node.Value = 1
	list.Insert(0, node)
	size = list.GetSize()
	fmt.Println("有一个结点的 list Insert(位置 0 )了一个结点 Value: 1")
	fmt.Println("list GetSize(): 应为 2， 实际： ", size)
	if size != 2 {
		t.Fatal("Insert() ====失败====")
	}

	node, err := list.GetNode(0)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("list GetNode(0)， 预期值 Value：1，实际值 Value：", node.Value)
	if node.Value != 1 {
		t.Fatal("GetNode(0) ====失败====")
	}

}