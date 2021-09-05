package ds

import (
	"fmt"
)

type QueueLinkedList struct {
	src  LinkedListInterface
	size int
}

func (q *QueueLinkedList) Size() int {
	return q.size
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueLinkedList) Enqueue(val int) {
	q.src.Append(val)
	q.size++
}

func (q *QueueLinkedList) First() (NodeInterface, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}
	return q.src.GetNodeAtIndex(q.size - 1)
}

func (q *QueueLinkedList) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	q.size--
	return q.src.RemoveAtIndex(q.size)
}

func (q *QueueLinkedList) String() {
	q.src.String()
}

func NewQueueSLL() *QueueLinkedList {
	q := &QueueLinkedList{}
	q.src = &LinkedList{}
	return q
}

func NewQueueDLL() *QueueLinkedList {
	q := &QueueLinkedList{}
	q.src = &DoublyLinkedList{}
	return q
}

func ExecuteQueueSLL() {
	qs := NewQueueSLL()

	fmt.Println(qs.IsEmpty())
	_, err := qs.First()
	fmt.Printf("%s \n", err)
	i, err := qs.Dequeue()
	fmt.Printf("%s %d\n", err, i)
	qs.Enqueue(10)
	qs.Enqueue(25)
	qs.Enqueue(35)
	qs.String()

	item, _ := qs.First()
	fmt.Println("first", item)

	i, _ = qs.Dequeue()
	fmt.Println("first dequeue", i)
	qs.String()
	i, _ = qs.Dequeue()
	fmt.Println("first dequeue", i)
	qs.String()
}

func ExecuteQueueDLL() {
	qs := NewQueueDLL()

	fmt.Println(qs.IsEmpty())
	_, err := qs.First()
	fmt.Printf("%s \n", err)
	i, err := qs.Dequeue()
	fmt.Printf("%s %d\n", err, i)
	qs.Enqueue(10)
	qs.Enqueue(25)
	qs.Enqueue(35)
	qs.String()

	item, _ := qs.First()
	fmt.Println("first", item)

	i, _ = qs.Dequeue()
	fmt.Println("first dequeue", i)
	qs.String()
	i, _ = qs.Dequeue()
	fmt.Println("first dequeue", i)
	qs.String()
}
