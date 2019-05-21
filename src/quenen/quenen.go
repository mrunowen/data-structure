/*
	队列数据结构

*/

package quenen

type Quenener interface {
	//	入队
	Enquenen(elem Elem)
	//	出队
	Dequenen() *Elem
	//	队列元素数量
	GetLen() int
}

type Elem struct {
	Value int
}

type Quenen struct {
	data []Elem
}

func New() Quenener {
	return &Quenen{}
}

func (q *Quenen) Enquenen(elem Elem) {
	q.data = append(q.data, elem)
}

func (q *Quenen) Dequenen() *Elem {
	length := q.GetLen()
	if length == 0 {
		return nil
	}
	resultElem := q.data[length-1:]
	q.data = q.data[:length-1]
	return &resultElem[0]
}

func (q *Quenen) GetLen() int {
	return len(q.data)
}
