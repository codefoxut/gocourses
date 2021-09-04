package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	person := make(map[string]string)

	var name, address string

	fmt.Print("Please enter a name: ")
	fmt.Scanln(&name)

	fmt.Print("Please enter an address: ")
	fmt.Scanln(&address)

	person["name"] = name
	person["address"] = address

	barr, err := json.Marshal(person)
	if err != nil  {
		fmt.Println("erros is ", err)
	}  else {
		fmt.Println("json object is ",  string(barr))
	}
}