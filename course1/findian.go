package main

import (
	"fmt"
	"strings"
)

func main() {

	var user_input string

	fmt.Println("Please enter a string.")
	fmt.Scanln(&user_input)

	lower_input := strings.ToLower(user_input)
	len_input := len(user_input)

	switch {
	case lower_input[0] == 'i' && lower_input[len_input-1] == 'n' && strings.Contains(lower_input, "a"):
		fmt.Println("Fonud!")
	default:
		fmt.Println("Not Found!")
	}
}
