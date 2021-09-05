package ds

import (
	"fmt"
	"sync"
)

// singly linked list execution

type StackLinkedList struct {
	src  LinkedListInterface
	lock sync.RWMutex
}

func (s *StackLinkedList) Push(val int) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	s.src.Append(val)
}

func (s *StackLinkedList) Pop() (int, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	value, err := s.src.RemoveAtIndex(0)
	return value, err
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.src.IsEmpty()
}

func (s *StackLinkedList) Size() int {
	return s.src.Size()
}

func (s *StackLinkedList) Top() NodeInterface {
	return s.src.Head()
}

func (s *StackLinkedList) String() {
	s.lock.RLock()
	defer s.lock.RUnlock()
	s.src.String()
}

func NewStackSLL() StackInterface {
	stack := &StackLinkedList{}
	stack.src = NewLinkedList()
	return stack
}

func NewStackDLL() StackInterface {
	stack := &StackLinkedList{}
	stack.src = NewDoublyLinkedListWithInterface()
	return stack
}

func ExecuteStackSLL() {
	stack := NewStackSLL()
	val, err := stack.Pop()
	fmt.Println(val, err)
	stack.Push(100)
	stack.Push(500)
	stack.Push(400)
	stack.String()
	val, err = stack.Pop()
	fmt.Println(val, err)
	fmt.Println(stack.Top().GetValue())
}

func ExecuteStackDLL() {
	stack := NewStackDLL()
	val, err := stack.Pop()
	fmt.Println(val, err)
	stack.Push(100)
	stack.Push(500)
	stack.Push(400)
	stack.String()
	val, err = stack.Pop()
	fmt.Println(val, err)
	fmt.Println(stack.Top().GetValue())
}

// func main() {
// 	ExecuteStackSLL()
// }
