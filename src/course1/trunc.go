package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main3() {
	lst := make([]int, 0, 3)
	for {
		fmt.Println("Enter an integer: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "X" {
			break
		}
		if i, err := strconv.Atoi(text); err == nil {
			lst = append(lst, i)
			sort.Ints(lst)
			fmt.Println(lst)
		} else {
			fmt.Println(err)
		}
	}
}

func main() {
	makejson()
	main3()
	main2()
	trunc1()
	trunc2()
}

func main2() {
	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)

}

func trunc1() {
	var user_input string

	fmt.Println("Please provide a floating number.")
	fmt.Scanln(&user_input)

	f, err := strconv.ParseFloat(user_input, 64)

	if err != nil {
		fmt.Printf("Error in  conversion %s", err)
	}

	int_input := int64(f)

	fmt.Println(int_input)

}

func trunc2() {
	var user_input, f float64

	fmt.Println("Please provide a floating number.")
	_, err := fmt.Scanln(&user_input)

	// f, err := strconv.ParseFloat(user_input, 64)

	if err != nil {
		fmt.Printf("Error in  conversion %s", err)
	}
	f = user_input

	int_input := int64(f)

	fmt.Println(int_input)
}

func makejson() {

	var inputName, inputAddr string
	m := make(map[string]string)

	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Printf("Please enter your first and last name: ")
	inputName, _ = consoleReader.ReadString('\n')

	fmt.Printf("Please enter your address: ")
	inputAddr, _ = consoleReader.ReadString('\n')

	m["name"] = inputName
	m["address"] = inputAddr

	b, _ := json.Marshal(m)

	fmt.Println("\nThe 'map' data structure created from the user input: \n", m)
	fmt.Println("\nThe JSON byte array: \n", b)
	fmt.Printf("\nThe JSON byte array output as a string:\n %s\n", b)

}
