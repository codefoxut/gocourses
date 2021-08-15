package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (cow Cow) Speak() {
	fmt.Println("moo")

}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (acow Cow) Eat() {
	fmt.Println("grass")
}

type Snake struct {
	name string
}

func (snake Snake) Speak() {
	fmt.Println("hsss")

}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

type Bird struct {
	name string
}

func (bird Bird) Speak() {
	fmt.Println("peep")

}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Eat() {
	fmt.Println("worm")
}

func CreateAnimal(name, kind string) Animal {
	var animal Animal
	switch kind {
	case "cow":
		cow := Cow{name}
		animal = cow
	case "bird":
		bird := Bird{name}
		animal = bird
	case "snake":
		snake := Snake{name}
		animal = snake
	}
	fmt.Println("Creating it!")
	return animal

}

func QueryAnimal(animal Animal, action string) {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func main() {
	var animal Animal

	zoo := make(map[string]Animal)

	fmt.Println("Please enter your request. Example:> newanimal naagraj snake")
	for {
		fmt.Print(">")
		var request string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			request = scanner.Text()
		}

		request_data := strings.Split(request, " ")
		if len(request_data) != 3 {
			log.Fatal("Data is not right.")
		}

		switch request_data[0] {
		case "newanimal":
			zoo[request_data[1]] = CreateAnimal(request_data[1], request_data[2])
		case "query":
			animal = zoo[request_data[1]]
			QueryAnimal(animal, request_data[2])
		default:
			break
		}
	}
}
