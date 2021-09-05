package main


type NodeInterface interface {
	String()
}

type LinkedListInterface interface {
	Append(val int)
	Insert(index, val int)
	Head() NodeInterface
}