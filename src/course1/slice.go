package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	int_slice := make([]int, 0, 3)

	for {
		var temp_str string
		fmt.Println("Please enter an integer")
		fmt.Scan(&temp_str)
		if temp_str == "X" {
			break
		}
		val, err := strconv.ParseInt(temp_str, 10, 64)
		if err != nil {
			fmt.Println("bad data is inserted.")

		} else {
			int_slice = append(int_slice, int(val))

			sort.Ints(int_slice)
			fmt.Println("Slice : ", int_slice)
		}

	}
}
