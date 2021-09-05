package ds


type NodeInterface interface {
	GetValue() int
	String()
}

type LinkedListInterface interface {
	Append(val int)
	InsertAtIndex(index, val int) error
	Head() NodeInterface
	GetNodeAtIndex(index int) (NodeInterface, error)
	RemoveAtIndex(index int)  (int, error)
	IsEmpty() bool
	Size() int
	String()

}