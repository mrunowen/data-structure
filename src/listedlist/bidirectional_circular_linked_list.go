package linkedlist

import (
	"errors"
	"time"
)

const (
	NOT_FOUND = -1
	//	中间结点
	MID_NODE = iota
	//	头结点
	HEAD_NODE
	//	尾结点
	TAIL_NODE
)

var IndexOutOfRangeException = errors.New("index out of range")

type LinkedList struct {
	headNode   *Node
	tailNode   *Node
	CreateDate time.Time
}

type Node struct {
	//	前驱结点
	prev *Node
	//	后驱结点
	next *Node
	//	结点类型
	nodeType int
	//	值
	Value int
}

//	生成一个双向循环链表
func New() LinkedLister {
	headNode, tailNode := newHeadNode()
	linkedList := &LinkedList{
		headNode:   headNode,
		tailNode:   tailNode,
		CreateDate: time.Now().UTC(),
	}
	return linkedList
}

func newHeadNode() (headNode, tailNode *Node) {
	headNode = &Node{
		nodeType: HEAD_NODE,
	}
	tailNode = &Node{
		prev:     headNode,
		next:     headNode,
		nodeType: TAIL_NODE,
	}
	headNode.prev = tailNode
	headNode.next = tailNode
	return
}

//	生成一个结点
func NewNode() *Node {
	return &Node{nodeType: MID_NODE}
}

//	返回链表元素个数
func (n *LinkedList) GetSize() int {
	count := 0
	node := n.headNode.next
	for node.nodeType == MID_NODE {
		count++
		node = node.next
	}
	return count
}

//	链表是否为空
func (n *LinkedList) IsEmpty() bool {
	return n.headNode.next.nodeType == TAIL_NODE
}

//	获取指定索引的结点
func (n *LinkedList) GetNode(index int) (Node, error) {
	size := n.GetSize()
	if size == 0 || index > size-1 || index < 0 {
		return Node{}, IndexOutOfRangeException
	}

	node := n.headNode

	for i := 0; i <= index; i++ {
		node = node.next
	}

	return *node, nil
}

//	寻找特定值的结点
//	return: 第一个匹配到的结点的下标
//			若没有找到结点，返回 NOT_FOUND
func (n *LinkedList) FindNode(node Node) int {
	thisNode := n.headNode.next
	count := 0
	for thisNode.nodeType != TAIL_NODE {
		if thisNode.Value == node.Value {
			return count
		}
		count++
		thisNode = thisNode.next
	}
	return NOT_FOUND
}

//	向链表中指定索引处插入结点
func (n *LinkedList) Insert(index int, node Node) error {
	size := n.GetSize()
	if size == 0 || index > size-1 || index < 0 {
		return IndexOutOfRangeException
	}
	thisNode := n.headNode
	for i := 0; i <= index; i++ {
		thisNode = thisNode.next
	}
	prevNode := thisNode.prev
	prevNode.next = &node
	node.prev = prevNode
	node.next = thisNode
	thisNode.prev = &node

	return nil
}

//	向链表尾部添加结点
func (n *LinkedList) Append(node Node) {
	tailNode := n.tailNode

	prevNode := tailNode.prev
	prevNode.next = &node
	node.prev = prevNode
	node.next = tailNode
	tailNode.prev = &node
}

//	删除链表中指定位置的元素
func (n *LinkedList) Remove(index int) error {
	size := n.GetSize()
	if size == 0 || index > size-1 || index < 0 {
		return IndexOutOfRangeException
	}

	thisNode := n.headNode
	for i := 0; i <= index; i++ {
		thisNode = thisNode.next
	}
	prevNode := thisNode.prev
	nextNode := thisNode.next
	prevNode.next = nextNode
	nextNode.prev = prevNode

	thisNode.prev = nil
	thisNode.next = nil
	thisNode.nodeType = NOT_FOUND
	thisNode.Value = 0

	return nil
}

//	翻转整个链表
func (n *LinkedList) Reverse() {
	//	若链表除了头尾结点外没有结点或只有一个
	//	那么翻不翻转都无所谓了
	if n.GetSize() <= 1 {
		return
	}

	thisNode := n.tailNode.prev
	newHead, newTail := newHeadNode()
	n.headNode = newHead
	n.tailNode = newTail
	for thisNode.nodeType != HEAD_NODE {
		n.Append(*thisNode)
		thisNode = thisNode.prev
	}
}

//	链表升序排序
func (n *LinkedList) Sort() {
	//	未实现
}
