package linkedlist

import (
	"errors"
)

const (
	NOT_FOUND = iota
	//	头结点
	HEAD_NODE
	//	尾结点
	TAIL_NODE
	//	中间结点
	MID_NODE
)

type Node struct {
	//	前驱结点
	prev *Node
	//	后驱结点
	next *Node
	//	结点类型
	Type int
	//	值
	Value string
}

func New() LinkedLister {
	return newBidirectionalCircularLinkedList()
}

func newBidirectionalCircularLinkedList() LinkedLister {
	headNode := &Node {
		Type: HEAD_NODE,
	}
	tailNode := &Node {
		prev: headNode,
		next: headNode,
		Type: TAIL_NODE,
	}
	headNode.prev = tailNode
	headNode.next = tailNode
	return headNode
}

//	返回链表元素个数
func (n *Node) GetSize() int {
	return 0
}
//	链表是否为空
func (n *Node) IsEmpty() bool {
	return false
}
//	获取指定索引的结点
func (n *Node) GetNode(index int) error {
	return errors.New("")
}
//	寻找特定值的结点
//	return: 第一个匹配到的结点的下标
//			若没有找到结点，返回 NOT_FOUND
func (n *Node) FindNode(node LinkedLister) int {
	return 0
}
//	向链表中指定索引处插入结点
func (n *Node) Insert(index int, node LinkedLister) error {
	return errors.New("")	
}
//	删除链表中指定位置的元素
func (n *Node) Remove(index int) error {
	return errors.New("")	
}
//	翻转整个链表
func (n *Node) Reverse() {
	
}
//	链表升序排序
func (n *Node) Sort() {
	
}
