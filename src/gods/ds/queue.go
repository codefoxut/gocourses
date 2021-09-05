package ds


type QueueInterface interface {
	Enqueue(val int)
	Dequeue() (int, error)
	First() (NodeInterface, error)
	IsEmpty() bool
	Size() int
	String() 
}