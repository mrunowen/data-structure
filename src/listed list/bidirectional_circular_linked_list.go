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

type LinkedList struct {
	headNode *Node
}

type Node struct {
	//	前驱结点
	prev *Node
	//	后驱结点
	next *Node
	//	结点类型
	nodeType int
	//	值
	Value string
}

//	生成一个双向循环链表
func New() LinkedLister {
	headNode := &Node {
		nodeType: HEAD_NODE,
	}
	tailNode := &Node {
		prev: headNode,
		next: headNode,
		nodeType: TAIL_NODE,
	}
	headNode.prev = tailNode
	headNode.next = tailNode
	linkedList := &LinkedList {
		headNode: headNode,
	}
	return linkedList
}

//	生成一个结点
func NewNode() *Node {
	return &Node { nodeType: MID_NODE }
}

//	返回链表元素个数
func (n *LinkedList) GetSize() int {
	count := 0
	node:= n.headNode.next
	for node.nodeType == MID_NODE {
		count++
	}
	return count
}
//	链表是否为空
func (n *LinkedList) IsEmpty() bool {
	return n.headNode.next.nodeType == TAIL_NODE
}
//	获取指定索引的结点
func (n *LinkedList) GetNode(index int) error {
	return errors.New("")
}
//	寻找特定值的结点
//	return: 第一个匹配到的结点的下标
//			若没有找到结点，返回 NOT_FOUND
func (n *LinkedList) FindNode(node Node) int {
	return 0
}
//	向链表中指定索引处插入结点
func (n *LinkedList) Insert(index int, node Node) error {
	return errors.New("")	
}
//	删除链表中指定位置的元素
func (n *LinkedList) Remove(index int) error {
	return errors.New("")	
}
//	翻转整个链表
func (n *LinkedList) Reverse() {
	
}
//	链表升序排序
func (n *LinkedList) Sort() {
	
}
