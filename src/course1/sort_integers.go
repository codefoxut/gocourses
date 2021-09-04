package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main2() {

	// read user input.
	var temp_str string
	fmt.Print("Please enter series of integer: example: 12 23 24 81 43 21 84 1999 1 31 91: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		temp_str = scanner.Text()
	}

	// create list of integers.
	temp_list := strings.Split(temp_str, " ")
	temp_integers := make([]int, len(temp_list))
	for i := 0; i < len(temp_list); i++ {
		temp_integers[i], _ = strconv.Atoi(temp_list[i])
	}

	// divide array in 4 arrays.
	// var array1, array2, array3, array4  []int
	small_array_size := (len(temp_integers) / 4) + 1
	array1 := temp_integers[0:small_array_size]
	array2 := temp_integers[small_array_size:(2 * small_array_size)]
	array3 := temp_integers[(2 * small_array_size):(3 * small_array_size)]
	array4 := temp_integers[(3 * small_array_size):len(temp_integers)]

	fmt.Println("divided arrays are", array1, array2, array3, array4)

	// sort all small arrays with goroutines.
	cn := make(chan []int)
	go BubbleSortChannel(cn, array1)
	go BubbleSortChannel(cn, array2)
	go BubbleSortChannel(cn, array3)
	go BubbleSortChannel(cn, array4)

	final_seq := make([]int, 0, len(temp_integers))

	//merge sorted arrays
	final_seq = MergeArrays(<-cn, final_seq)
	final_seq = MergeArrays(<-cn, final_seq)
	final_seq = MergeArrays(<-cn, final_seq)
	final_seq = MergeArrays(<-cn, final_seq)
	fmt.Println("final response", final_seq)
	close(cn)

}

// MergeArrays.. to create sorted array.
func MergeArrays(array1 []int, final_resp []int) []int {
	// fmt.Println("merging", array1, final_resp)
	final_len := len(final_resp)
	data_len := len(array1)
	x := 0
	y := 0
	response := make([]int, 0, final_len+data_len)
	for x < data_len && y < final_len {
		if array1[x] < final_resp[y] {
			response = append(response, array1[x])
			x++
		} else {
			response = append(response, final_resp[y])
			y++
		}
	}

	for x < data_len {
		response = append(response, array1[x])
		x++
	}

	for y < final_len {
		response = append(response, final_resp[y])
		y++
	}

	// fmt.Println("response merging", response)
	return response
}

//BubbleSort...
func BubbleSortChannel(cn chan []int, int_seq []int) {
	fmt.Println("sorting array ", int_seq)
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
	cn <- int_seq
	fmt.Println("sorted sequence is ", int_seq)

}

func Swap(int_seq []int, index int) {
	int_seq[index], int_seq[index+1] = int_seq[index+1], int_seq[index]
}

func swap(intSlice []int, pos int) {
	// swaps curr pos int with adj
	temp := intSlice[pos]
	intSlice[pos] = intSlice[pos+1]
	intSlice[pos+1] = temp
}

func bubblesort(subArray []int, wg *sync.WaitGroup) {
	for i := len(subArray) - 1; i >= 0; i-- {
		var isSwapped bool = false
		for j := 0; j < i; j++ {
			if subArray[j] > subArray[j+1] {
				isSwapped = true
				swap(subArray, j)
			}
		}
		// sorting complete
		if !isSwapped {
			break
		}
	}
	// print sorted subarray
	fmt.Println(subArray)
	wg.Done()
}

func finalSort(intSlice []int) {
	for i := len(intSlice) - 1; i >= 0; i-- {
		var isSwapped bool = false
		for j := 0; j < i; j++ {
			if intSlice[j] > intSlice[j+1] {
				isSwapped = true
				swap(intSlice, j)
			}
		}
		if !isSwapped {
			return
		}
	}
}

func main() {
	numArray := make([]int, 0, 3)
	var userInput string
	fmt.Println("Enter an integer to continue, enter 'x' to exit program")
	fmt.Scan(&userInput)
	// Take input of array of integers
	for {
		if strings.ToLower(userInput) == "x" {
			fmt.Println("Sorted Array")
			var wg sync.WaitGroup
			wg.Add(4)
			// Sort using 4 separate goroutines
			quarterIndex := len(numArray) / 4
			halfIndex := len(numArray) / 2
			go bubblesort(numArray[0:quarterIndex], &wg)
			go bubblesort(numArray[quarterIndex:halfIndex], &wg)
			go bubblesort(numArray[halfIndex:halfIndex+quarterIndex], &wg)
			go bubblesort(numArray[halfIndex+quarterIndex:], &wg)
			wg.Wait()
			finalSort(numArray)
			// print entire sorted array
			fmt.Println(numArray)
			return
		}

		// if not yet at end, append user input
		curr, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a valid number or the letter 'x' to continue")
		} else {
			numArray = append(numArray, curr)
		}

		fmt.Println("Enter an integer to continue, enter 'x' to exit the program")
		fmt.Scan(&userInput)
	}
}
