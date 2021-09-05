package ds

import (
	"fmt"
	"sync"
)

type QueueSlice struct {
	src  []*ArrayNode
	size int
	lock sync.RWMutex
}

func (q *QueueSlice) Enqueue(val int) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	node := &ArrayNode{val}
	q.src = append(q.src, node)
	q.size++
}

func (q *QueueSlice) Dequeue() (int, error) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	node := q.src[0]
	q.src = q.src[1:len(q.src)]
	q.size--
	return node.GetValue(), nil
}

func (q *QueueSlice) First() (NodeInterface, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}
	return q.src[0], nil
}

func (q *QueueSlice) Size() int {
	return q.size
}

func (q *QueueSlice) IsEmpty() bool {
	return q.Size() == 0
}
func (q *QueueSlice) String() {
	for _, v := range q.src {
		fmt.Print(v.GetValue(), " ")
	}
	fmt.Println()

}

func NewQueueSlice() QueueInterface {
	queue := &QueueSlice{}
	queue.src = make([]*ArrayNode, 0)
	return queue
}

func ExecuteQueueSlice() {
	qs := NewQueueSlice()

	fmt.Println(qs.IsEmpty())
	_, err := qs.First()
	fmt.Printf("%s \n", err)
	i, err  := qs.Dequeue()
	fmt.Printf("%s %d\n", err, i)
	qs.Enqueue(10)
	qs.Enqueue(25)
	qs.Enqueue(35)
	qs.String()

	item, _ := qs.First()
	item.String()

	i, _ = qs.Dequeue()
	fmt.Println("first dequeue", i)
	qs.String()
	i, _ = qs.Dequeue()
	fmt.Println("first dequeue", i)
	qs.String()
}
