package ds

type StackInterface interface {
	Push(val int)
	Pop() (int, error)
	IsEmpty() bool
	Size() int
	Top() NodeInterface
	String()
}
