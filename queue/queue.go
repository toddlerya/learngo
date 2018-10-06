package queue

// An FIFO queue
type Queue []interface{}

// Pushes the element into the queue
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int))
}

// Pops the element from head.
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

// Return if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
