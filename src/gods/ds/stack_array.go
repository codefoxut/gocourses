package ds

import (
	"fmt"
	"sync"
)

type ArrayNode struct {
	value int
}

func (n *ArrayNode) GetValue() int {
	return n.value
}

func (n *ArrayNode) String() {
	fmt.Println(n.value)
}

type StackSlice struct {
	src  []*ArrayNode
	size int
	lock sync.RWMutex
}

func (s *StackSlice) Push(val int) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	node := &ArrayNode{val}
	s.src = append(s.src, node)
	s.size++
}

func (s *StackSlice) Pop() (int, error) {
	if s.size == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	s.size--
	node := s.src[s.size]
	s.src = s.src[0:s.size]

	return node.GetValue(), nil
}

func (s *StackSlice) IsEmpty() bool {
	return s.size == 0
}

func (s *StackSlice) Size() int {
	return s.size
}

func (s *StackSlice) Top() NodeInterface {
	if !s.IsEmpty() {
		return nil
	}
	return s.src[0]
}

func (s *StackSlice) String() {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for _, v := range s.src {
		fmt.Print(v.GetValue())
	}
	fmt.Println()
}

func NewStackSlice() StackInterface {
	stack := &StackSlice{}
	return stack
}

func ExecuteStackSlice() {
	stack := NewStackSLL()
	val, err := stack.Pop()
	fmt.Print(val, err)
	stack.Push(100)
	stack.Push(500)
	stack.Push(400)
	stack.String()
	val, err = stack.Pop()
	fmt.Println(val, err)
	fmt.Println(stack.Top().GetValue())
}
