package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Please enter values for acceleration, initial velocity, and initial displacement.")
	var acceleration, velocity, displacement, duration float64
	fmt.Print("Please enter value for acceleration: ")
	_, err := fmt.Scanln(&acceleration)

	if err != nil {
		fmt.Printf("Error in acceleration conversion %s", err)
	}
	fmt.Print("Please enter value for initial velocity: ")
	_, err = fmt.Scanln(&velocity)

	if err != nil {
		fmt.Printf("Error in velocity conversion %s", err)
	}
	fmt.Print("Please enter value for initial displacement: ")
	_, err = fmt.Scanln(&displacement)

	if err != nil {
		fmt.Printf("Error in displacement conversion %s", err)
	}

	fmt.Print("Please enter a value for time: ")
	_, err = fmt.Scanln(&duration)

	if err != nil {
		fmt.Printf("Error in time conversion %s", err)
	}

	fn := GenDisplaceFn(acceleration, velocity, displacement)
	fmt.Println(fn(duration))

}

func GenDisplaceFn(acceleration, velocity, displacement float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*acceleration*math.Pow(t, 2) + velocity*t + displacement
	}
	return fn

}
