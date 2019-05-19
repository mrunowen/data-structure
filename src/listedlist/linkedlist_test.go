package linkedlist

import (
	"fmt"
	"testing"
)

var list LinkedLister

//	测试 New
func TestLinkedList(t *testing.T) {
	list = New()
	println("通过 New() 实例化: list")
	println("测试 GetSize(), 预计： 0")
	size := list.GetSize()
	fmt.Printf("测试 GetSize(), 实际： %d\n", size)
	if size != 0 {
		t.Fatal("GetSize() ====失败====")
	}

	isEmpty := list.IsEmpty()
	fmt.Println("测试 IsEmpty()， 预计： true")
	fmt.Println("测试 IsEmpty()， 实际： ", isEmpty)
	if !isEmpty {
		t.Fatal("IsEmpty() ====失败====")
	}

}

func TestAppend(t *testing.T) {
	node := *NewNode()
	node.Value = 2

	list.Append(node)
	size := list.GetSize()
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

func TestFindNode(t *testing.T) {
	node := NewNode()
	node.Value = 1

	fmt.Println("list find node, value is 0, the result should be 0")
	index := list.FindNode(*node)
	fmt.Println("node value is: ", node.Value)
	fmt.Println("============================the result is: ", index)

	if index != 0 {
		t.Fatal("func GetNode() Fatal")
	}

	fmt.Println("list find node, value is 2, the result should be 1")
	node.Value = 2
	index = list.FindNode(*node)
	fmt.Println("node value is: ", node.Value)
	fmt.Println("============================the result is: ", index)
	if index != 1 {
		t.Fatal("func GetNode() Fatal")
	}

}

func TestRemove(t *testing.T) {
	fmt.Println("\n\nnow list have two node")
	fmt.Println("value 1, index: 0")
	fmt.Println("value 2, index: 1")
	fmt.Println("now remove index 2, will exception...")

	err := list.Remove(2)
	if err == nil {
		t.Fatal("func Remove() Fatal")
	}
	fmt.Println("pass")

	fmt.Println("now Remove index 0, list Size will be 1, the value of node at index 0 will be 2...")
	err = list.Remove(0)
	if err != nil {
		t.Fatal(err)
	}
	size := list.GetSize()
	if size != 1 {
		t.Fatal("func Remove() Fatal")
	}
	node, err := list.GetNode(0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("now get value is: ", node.Value)
	if node.Value != 2 {
		t.Fatal("func Remove() value Fatal")
	}
}

func TestReverse(t *testing.T) {
	print("\n\nnow test func Reverse()...\n")
	print("now list have one node, value is 2\ninsert into index 0, now size should be 2\n")
	node := NewNode()
	node.Value = 1
	list.Insert(0, *node)
	size := list.GetSize()
	fmt.Println("list size is: ", size)
	if size != 2 {
		t.Fatal("list size error")
	}

	print("list index 0, value should be 1, index 1, vlaue should be 2\n")
	thisNode, err := list.GetNode(0)
	if err != nil {
		t.Fatal(err)
	}
	if thisNode.Value != 1 {
		t.Fatal("the value at index 0 error, is: ", thisNode.Value)
	}
	thisNode, err = list.GetNode(1)
	if err != nil {
		t.Fatal(err)
	}
	if thisNode.Value != 2 {
		t.Fatal("the value at index 1 error, is: ", thisNode.Value)
	}

	print("now reverse list\n")
	list.Reverse()
	print("reversed\n")
	thisNode, err = list.GetNode(0)
	if err != nil {
		t.Fatal(err)
	}
	if thisNode.Value != 2 {
		t.Fatal("the value at index 0 error, is: ", thisNode.Value)
	}
	print("the value at index 0 should be 2, is: ", thisNode.Value, "\n")

	thisNode, err = list.GetNode(1)
	if err != nil {
		t.Fatal(err)
	}
	if thisNode.Value != 1 {
		t.Fatal("the value at index 1 error, is: ", thisNode.Value)
	}
	print("the value at index 1 should be 1, is: ", thisNode.Value, "\n")

}

func TestSort(t *testing.T) {
	list.Sort()
	fmt.Println("sorted")
	thisNode, err := list.GetNode(0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("the value of list at index 0 is: ", thisNode.Value)

	thisNode, err = list.GetNode(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("the value of list at index 1 is: ", thisNode.Value)
}
