package ds

import "fmt"

type DEQueue interface {
	AddFirst(val int)
	AddLast(val int)
	DeleteFirst() (int, error)
	DeleteLast() (int, error)
	First() (NodeInterface, error)
	Last() (NodeInterface, error)
	IsEmpty() bool
	Size() int
	String()
}

type DEQueueLinkedList struct {
	src  LinkedListInterface
	size int
}

func (deq *DEQueueLinkedList) AddFirst(val int) {
	deq.src.Append(val)
	deq.size++
}

func (deq *DEQueueLinkedList) AddLast(val int) {
	deq.src.InsertAtIndex(deq.size, val)
	deq.size++
}

func (deq *DEQueueLinkedList) DeleteFirst() (int, error) {
	if deq.IsEmpty() {
		return 0, fmt.Errorf("deque is empty")
	}
	deq.size--
	return deq.src.RemoveAtIndex(deq.Size())
}

func (deq *DEQueueLinkedList) DeleteLast() (int, error) {
	if deq.IsEmpty() {
		return 0, fmt.Errorf("deque is empty")
	}
	deq.size--
	return deq.src.RemoveAtIndex(0)

}

func (deq *DEQueueLinkedList) First() (NodeInterface, error) {
	if deq.IsEmpty() {
		return nil, fmt.Errorf("deque is empty")
	}
	return deq.src.GetNodeAtIndex(deq.Size() - 1)
}

func (deq *DEQueueLinkedList) Last() (NodeInterface, error) {
	if deq.IsEmpty() {
		return nil, fmt.Errorf("deque is empty")
	}
	return deq.src.GetNodeAtIndex(0)

}
func (deq *DEQueueLinkedList) IsEmpty() bool {
	return deq.Size() == 0

}
func (deq *DEQueueLinkedList) Size() int {
	return deq.size
}
func (deqq *DEQueueLinkedList) String() {
	deqq.src.String()
}

func NewDEQueueSLL() *DEQueueLinkedList {
	q := &DEQueueLinkedList{}
	q.src = &LinkedList{}
	return q
}

func NewDEQueueDLL() *DEQueueLinkedList {
	q := &DEQueueLinkedList{}
	q.src = &DoublyLinkedList{}
	return q
}

func ExecuteDequeSLL() {
	qs := NewDEQueueSLL()

	fmt.Println(qs.IsEmpty())
	_, err := qs.First()
	fmt.Printf("%s \n", err)
	i, err := qs.DeleteFirst()
	fmt.Printf("%s %d\n", err, i)
	i, err = qs.DeleteLast()
	fmt.Printf("%s %d\n", err, i)
	qs.AddFirst(10)
	qs.AddFirst(25)
	qs.AddLast(35)
	qs.String()

	item, _ := qs.First()
	fmt.Println("first", item)
	item, _ = qs.Last()
	fmt.Println("last", item)

	i, _ = qs.DeleteLast()
	fmt.Println("first dequeue", i)
	qs.String()
	i, _ = qs.DeleteFirst()
	fmt.Println("first dequeue", i)
	qs.String()
}

func ExecuteDequeDLL() {
	qs := NewDEQueueDLL()

	fmt.Println(qs.IsEmpty())
	_, err := qs.First()
	fmt.Printf("%s \n", err)
	i, err := qs.DeleteFirst()
	fmt.Printf("%s %d\n", err, i)
	i, err = qs.DeleteLast()
	fmt.Printf("%s %d\n", err, i)
	qs.AddFirst(10)
	qs.AddFirst(25)
	qs.AddLast(35)
	qs.String()

	item, _ := qs.First()
	fmt.Println("first", item)
	item, _ = qs.Last()
	fmt.Println("last", item)

	i, _ = qs.DeleteLast()
	fmt.Println("last delete", i)
	qs.String()
	i, _ = qs.DeleteFirst()
	fmt.Println("first delete", i)
	qs.String()
}
