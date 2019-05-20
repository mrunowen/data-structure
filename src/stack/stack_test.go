/*
	测试栈结构
*/

package stack

import (
	"fmt"
	"testing"
)

var stack Stacker

func TestStack(t *testing.T) {
	stack = New()

	fmt.Println("new stack")
	len := stack.GetLen()
	fmt.Println("stack length should be 0, is: ", len)
	if len != 0 {
		t.Fatal("func GetLen() error")
	}

	fmt.Println("stack Push element, value 0")
	e := &Elem{Value: 0}
	stack.Push(*e)
	len = stack.GetLen()
	fmt.Println("now stack length should be 1, is: ", len)
	if len != 1 {
		t.Fatal("func GetLen() error")
	}

	fmt.Println("stack Push element, value 1")
	e.Value = 1
	stack.Push(*e)
	len = stack.GetLen()
	fmt.Println("now stack length should be 2, is: ", len)
	if len != 2 {
		t.Fatal("func GetLen() error")
	}

	popElem := stack.Pop()
	fmt.Println("stack pop element, value should be 1, is: ", popElem.Value)
	if popElem.Value != 1 {
		t.Fatal("func Pop() error")
	}
	len = stack.GetLen()
	fmt.Println("new stack length should be 1, is: ", len)
	if len != 1 {
		t.Fatal("func GetLen() error")
	}

	popElem = stack.Pop()
	fmt.Println("stack pop element, value should be 0, is: ", popElem.Value)
	if popElem.Value != 0 {
		t.Fatal("func Pop() error")
	}
	len = stack.GetLen()
	fmt.Println("new stack length should be 0, is: ", len)
	if len != 0 {
		t.Fatal("func GetLen() error")
	}

	popElem = stack.Pop()
	fmt.Println("stack pop element, popElem should be nil, is: ", popElem == nil)
	if popElem != nil {
		t.Fatal("func Pop() error")
	}

}
