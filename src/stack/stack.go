/*
	栈数据结构
*/
package stack

//	栈接口
type Stack interface {
	//	获取栈的长度
	GetLen() int
	//	进栈
	Push(elem Elem)
	//	出栈
	Pop() *Elem
}

type MyStack struct {
	data []Elem
}

//	栈元素
type Elem struct {
	Value int
}

func (m *MyStack) GetLen() int {
	return len(m.data)
}

func (m *MyStack) Push(elem Elem) {
	m.data = append(m.data, elem)
}

func (m *MyStack) Pop() *Elem {
	len := m.GetLen()
	if len < 1 {
		return nil
	}
	popElem := m.data[len - 1 : len][0]
	m.data = m.data[: len - 1]
	return &popElem
}

