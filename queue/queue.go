package queue

// Queue int队列
type Queue []interface{}

// Push 一个整数到队列
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int))
}

// Pop 一个整数从队列里面
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

// IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
