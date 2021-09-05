package ds

import (
	"fmt"
	"sync"
)

type DLNode struct {
	value        int
	nextNode     *DLNode
	previousNode *DLNode
}

func (n *DLNode) String() {
	fmt.Printf("%d", n.value)
}

func (n *DLNode) GetValue() int {
	return n.value
}

type DoublyLinkedList struct {
	head *DLNode
	size int
	lock sync.RWMutex
}

func (dl *DoublyLinkedList) Append(val int) {
	dl.lock.RLock()
	defer dl.lock.RUnlock()
	var nextNode, prevNode *DLNode
	node := DLNode{}
	node.value = val
	if dl.size > 0 {
		nextNode = dl.head
		nextNode.previousNode = &node
	}

	node.previousNode = prevNode
	node.nextNode = nextNode
	dl.head = &node
	dl.size++

}

func (dl *DoublyLinkedList) getNodeWithIndex(index int) (*DLNode, error) {
	if index < 0 || index > dl.size {
		return nil, fmt.Errorf("index out of bounds")
	}
	counter := 0
	var indexedNode *DLNode
	item := dl.head
	for {
		if counter == index {
			indexedNode = item
			break
		}
		counter++
		item = item.nextNode
	}
	return indexedNode, nil
}

func (dl *DoublyLinkedList) InsertAtIndex(index, val int) error {
	dl.lock.RLock()
	defer dl.lock.RUnlock()
	if dl.size == 0 || index == 0 {
		dl.Append(val)
		return nil
	}
	indexedNode, err := dl.getNodeWithIndex(index)
	if err != nil {
		return err
	}
	node := &DLNode{val, nil, nil}
	preIndexedNode := indexedNode.previousNode
	if preIndexedNode != nil {
		preIndexedNode.nextNode = node
	}

	node.previousNode = preIndexedNode
	node.nextNode = indexedNode
	indexedNode.previousNode = node
	dl.size++
	return nil

}

func (dl *DoublyLinkedList) RemoveAtIndex(index int) (int, error) {
	dl.lock.RLock()
	defer dl.lock.RUnlock()

	item, err := dl.getNodeWithIndex(index)
	if err != nil {
		return 0, err
	}

	if item.previousNode != nil {
		item.previousNode.nextNode = item.nextNode
	}
	if item.nextNode != nil {
		item.nextNode.previousNode = item.previousNode
	}
	if item == dl.Head() {
		dl.head = dl.head.nextNode
	}
	item.previousNode = nil
	item.nextNode = nil
	dl.size--

	return item.value, nil

}
func (dl *DoublyLinkedList) IndexOf(val int) int {
	dl.lock.RLock()
	defer dl.lock.RUnlock()
	counter := -1
	found := false
	for item := dl.head; item != nil; item = item.nextNode {
		counter++
		if item.value == val {
			found = true
			break
		}
	}
	if !found {
		return -1
	}
	return counter

}
func (dl *DoublyLinkedList) IsEmpty() bool {
	dl.lock.RLock()
	defer dl.lock.RUnlock()
	return dl.Size() == 0

}
func (dl *DoublyLinkedList) Size() int {
	dl.lock.RLock()
	defer dl.lock.RUnlock()
	return dl.size
}
func (dl *DoublyLinkedList) String() {
	dl.lock.RLock()
	defer dl.lock.RUnlock()

	for item := dl.head; item != nil; item = item.nextNode {
		fmt.Print(item.value, " ")
	}
	fmt.Println()

}

func (dl *DoublyLinkedList) Head() NodeInterface {
	dl.lock.RLock()
	defer dl.lock.RUnlock()
	return dl.head
}

func (dl *DoublyLinkedList) GetNodesBetweenValues(value1, value2 int) []*DLNode {
	dl.lock.RLock()
	defer dl.lock.RUnlock()
	result := make([]*DLNode, 0)

	for item := dl.head; item != nil; item = item.nextNode {
		if value1 < item.value && item.value < value2 {
			result = append(result, item)
		}
	}
	return result
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func NewDoublyLinkedListWithInterface() LinkedListInterface {
	return &DoublyLinkedList{}
}

func ExecuteDoublyLinkedList() {
	dll := NewDoublyLinkedList()
	x := dll.IndexOf(36)
	fmt.Println(dll.IsEmpty())
	fmt.Println("index fof 36", x)
	dll.Append(10)
	dll.Append(12)
	dll.Append(14)
	dll.String()
	dll.InsertAtIndex(2, 25)
	dll.InsertAtIndex(1, 29)
	dll.InsertAtIndex(29, 1)
	dll.String()
	item, err := dll.RemoveAtIndex(0)
	fmt.Println(item, err)
	item, err = dll.RemoveAtIndex(3)
	fmt.Println(item, err)
	dll.String()
	item, err = dll.RemoveAtIndex(6)
	fmt.Println(item, err)
	x = dll.IndexOf(10)
	fmt.Println("index of 10", x)
	fmt.Println(dll.Head(), dll.Size())
	fmt.Println(dll.GetNodesBetweenValues(10, 20))

}

// func main() {
// 	ExecuteDoublyLinkedList()
// }
