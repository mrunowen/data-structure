/*
	树 数据结构
*/

package tree

type Treer interface {
	//	销毁树
	Destroy()
	//	是否为叶子节点
	IsLeaf() bool
	//	计算树的深度
	GetDepth() int
	//	计算树的结点总个数
	GetCount() int
	//	前序遍历
	PreOrder()
	//	中序遍历
	InOrder()
	//	后序遍历
	PostOrder()
	//	层序遍历
	LevelOrder()
}