package animal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)

}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func main() {
	var animal Animal
	cow := Animal{"Grass", "walk", "moo"}
	snake := Animal{"mice", "slither", "hsss"}
	bird := Animal{"worm", "fly", "peep"}

	zoo := map[string]Animal{
		"cow":   cow,
		"snake": snake,
		"bird":  bird,
	}

	fmt.Println("Please enter your request. Example:> cow eat.")
	for {
		fmt.Print(">")
		var request string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			request = scanner.Text()
		}

		request_data := strings.Split(request, " ")
		if len(request_data) != 2 {
			log.Fatal("Data is not right.")
		}
		animal = zoo[request_data[0]]

		switch request_data[1] {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			break
		}
	}
}
