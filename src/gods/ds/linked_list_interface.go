package ds


type NodeInterface interface {
	GetValue() int
	String()
}

type LinkedListInterface interface {
	Append(val int)
	InsertAtIndex(index, val int) error
	Head() NodeInterface
	RemoveAtIndex(index int)  (int, error)
	IsEmpty() bool
	Size() int
	String()

}