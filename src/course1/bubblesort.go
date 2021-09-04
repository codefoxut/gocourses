package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please provide a sequence of integers(max  10,  separate them with a  space) !")
	var seq string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		seq = scanner.Text()
	}

	str_numbers := strings.Split(seq, " ")
	int_numbers := make([]int, len(str_numbers))
	for i := 0; i < len(str_numbers); i++ {
		int_numbers[i], _ = strconv.Atoi(str_numbers[i])
	}

	BubbleSort(int_numbers)

}

//BubbleSort...
func BubbleSort(int_seq []int) {
	last_swap_index := len(int_seq) - 1 // a variable to use for already sorted sequence.
	start := 0
	for i := 0; i < len(int_seq); i++ {
		start = 0
		for j := 0; j < last_swap_index; j++ {
			// compare the number and swap
			if int_seq[j] > int_seq[j+1] {
				Swap(int_seq, j)
				start = j
				// fmt.Println("swap index", j)
			}
			// fmt.Println(int_seq)

		}
		// fmt.Println(start, largest_num_index)
		last_swap_index = start
	}

	fmt.Println("sorted sequence is ", int_seq)

}

func Swap_old(int_seq []int, index int) {
	int_seq[index], int_seq[index+1] = int_seq[index+1], int_seq[index]
}
