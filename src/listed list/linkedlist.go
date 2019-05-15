/*
	链表接口
*/

package linkedlist

type LinkedLister interface {
	//	返回链表元素个数
	GetSize() int
	//	链表是否为空
	IsEmpty() bool
	//	获取指定索引的结点
	GetNode(index int) (Node, error)
	//	寻找特定值的结点
	//	return: 第一个匹配到的结点的下标
	//			若没有找到结点，返回 NOT_FOUND
	FindNode(node Node) int
	//	向链表中指定索引处插入结点
	Insert(index int, node Node) error
	//	删除链表中指定位置的元素
	Remove(index int) error
	//	翻转整个链表
	Reverse()
	//	链表升序排序
	Sort()
}
