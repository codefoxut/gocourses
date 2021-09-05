package main

import (
	"fmt"

	"gods/ds"
)

func main() {
	// ds.ExecuteLinkedList()
	// ds.ExecuteDoublyLinkedList()
	// ds.ExecuteStackSLL()
	// fmt.Println("slice execution")
	// ds.ExecuteStackSlice()
	// fmt.Println("dll execution")
	// ds.ExecuteStackDLL()

	// fmt.Println("//Queue//")
	// ds.ExecuteQueueSlice()
	// fmt.Println("//Queue  SLL//")
	// ds.ExecuteQueueSLL()
	// fmt.Println("//Queue DLL//")
	// ds.ExecuteQueueDLL()

	fmt.Println("//Dequeue//")
	ds.ExecuteDequeSLL()
	fmt.Println("//Dequeue-- DLL//")
	ds.ExecuteDequeDLL()

}
