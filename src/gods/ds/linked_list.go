package ds

import (
	"fmt"
	"sync"
)

type LLNode struct {
	value    int
	nextNode *LLNode
}

func (n *LLNode) String() {
	fmt.Println(n.value)
}

func (n *LLNode) GetValue() int {
	return n.value
}

type LLError struct {
	msg string
}

func (ll *LLError) Error() string {
	return ll.msg
}

type LinkedList struct {
	head *LLNode
	size int
	lock sync.RWMutex
}

func (ll *LinkedList) Append(val int) {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := &LLNode{value: val, nextNode: ll.head}
	fmt.Println(node)
	ll.head = node
	ll.size++
}

func (ll *LinkedList) IterateList() []int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	result := make([]int, 0)

	for item := ll.head; item != nil; item = item.nextNode {
		result = append(result, item.value)
	}
	return result
}

func (ll *LinkedList) String() {

	for item := ll.head; item != nil; item = item.nextNode {
		fmt.Print(item.value, " ")
	}
	fmt.Println()
}

func (ll *LinkedList) GetLastNode() *LLNode {
	var lastNode *LLNode
	for item := ll.head; item != nil; item = item.nextNode {
		if item.nextNode == nil {
			lastNode = item
		}
	}
	return lastNode
}

// AddToEnd .. Can be treated as Append method
func (ll *LinkedList) AddToEnd(val int) {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	lastNode := ll.GetLastNode()
	newNode := &LLNode{value: val}
	if lastNode != nil {
		lastNode.nextNode = newNode
	} else {
		ll.head = newNode
	}
	ll.size++
}

func (ll *LinkedList) GetNodeAtIndex(index int) (NodeInterface, error) {
	return ll.getNodeByIndex(index)
}

func (ll *LinkedList) getNodeByIndex(index int) (*LLNode, error) {
	var err error
	var indexedNode *LLNode
	// find node at index or the last one
	counter := 0
	for item := ll.head; item != nil; item = item.nextNode {

		if counter == index {
			indexedNode = item
			break
		}
		counter++

	}
	if indexedNode == nil {
		err = &LLError{"index is nil."}
	}
	return indexedNode, err

}

func (ll *LinkedList) IndexOf(val int) int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	// find node at index or the last one
	counter := 0
	found := false
	for item := ll.head; item != nil; item = item.nextNode {
		if item.value == val {
			found = true
			break
		}
		counter++
	}
	if !found {
		return -1
	}
	return counter
}

func (ll *LinkedList) InsertAtIndex(index, val int) error {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := &LLNode{value: val}
	var indexedNode *LLNode
	// find node at index or the last one
	counter := 0
	for item := ll.head; item != nil && counter < index; item = item.nextNode {
		counter++
		indexedNode = item
	}

	if indexedNode == nil {
		ll.head = node
	} else {
		node.nextNode = indexedNode.nextNode
		indexedNode.nextNode = node
	}
	ll.size++
	return nil
}

func (ll *LinkedList) AddAfterValue(val, newNodeValue int) error {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	newNode := &LLNode{value: newNodeValue}
	var err error
	valuedNode, err := ll.FindNodeWithValue(val)
	if err == nil {
		newNode.nextNode = valuedNode.nextNode
		valuedNode.nextNode = newNode
		ll.size++
	} else {
		err = &LLError{"value not found in list"}
	}

	return err
}

func (ll *LinkedList) FindNodeWithValue(val int) (*LLNode, error) {
	var nodeWith *LLNode
	var err error
	found := false
	for item := ll.head; item != nil; item = item.nextNode {
		if item.value == val {
			nodeWith = item
			found = true
			break
		}
	}
	if !found {
		err = &LLError{"value not found in list"}
	}
	return nodeWith, err

}

func (ll *LinkedList) IsEmpty() bool {
	return ll.head == nil
}

func (ll *LinkedList) Head() NodeInterface {
	return ll.head
}

func (ll *LinkedList) Size() (counter int) {
	for item := ll.head; item != nil; item = item.nextNode {
		counter++
	}
	fmt.Println("ll.size", ll.size)
	return counter
}

func (ll *LinkedList) RemoveAtIndex(index int) (int, error) {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	if index < 0 || index > ll.size || ll.size == 0 {
		return 0, fmt.Errorf("index out of bounds")
	}
	if index == 0 {
		item := ll.head
		ll.head = item.nextNode
		return item.value, nil
	}
	indexedNode1, _ := ll.getNodeByIndex(index - 1)
	indexedNode := indexedNode1.nextNode
	indexedNode1.nextNode = indexedNode.nextNode
	indexedNode.nextNode = nil
	ll.size--
	return indexedNode.value, nil

}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func NewLinkedListInterface() LinkedListInterface {
	return &LinkedList{}
}

func ExecuteLinkedList() {
	linkedList := LinkedList{}
	fmt.Println(linkedList.head)
	linkedList.InsertAtIndex(19, 1000)
	linkedList.AddToEnd(12)
	linkedList.Append(10)
	// fmt.Println(linkedList)
	linkedList.Append(25)
	linkedList.AddToEnd(100)
	linkedList.InsertAtIndex(19, 1100)
	fmt.Println(linkedList.head.value)
	linkedList.String()
	fmt.Println(linkedList.FindNodeWithValue(12))
	fmt.Println(linkedList.FindNodeWithValue(120))
	fmt.Println("SIZE:", linkedList.Size())
	val, err := linkedList.RemoveAtIndex(1)
	fmt.Printf("remove at %v %s \n", val, err)
	linkedList.String()
	fmt.Println("index of 10", linkedList.IndexOf(10))

}

// func main() {
// 	ExecuteLinkedList()
// }
