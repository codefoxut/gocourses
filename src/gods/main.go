package main

import (
	"fmt"

	"gods/ds"
)

func main() {
	// ds.ExecuteLinkedList()
	ds.ExecuteDoublyLinkedList()
	ds.ExecuteStackSLL()
	fmt.Println("slice execution")
	ds.ExecuteStackSlice()
	fmt.Println("dll execution")
	ds.ExecuteStackDLL()

}
