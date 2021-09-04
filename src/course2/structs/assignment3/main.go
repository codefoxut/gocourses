package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	fmt.Println(filename, "this is the file.")
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error occured:", err)
	}
	io.Copy(os.Stdout, f)
}
