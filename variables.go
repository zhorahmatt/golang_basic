package main

import (
	"fmt"
)

func main()  {
	//static variable declaration
	var numbers int
	numbers = 10

	var floating_numbers float64
	floating_numbers = 100.3

	fmt.Println(numbers)
	fmt.Println(floating_numbers)

	fmt.Printf("Data type of numbers is %T \n", numbers)
	fmt.Printf("Data type of floating numbers is %T \n", floating_numbers)


	//dynamic declaration variable
	cars := "Toyota"
	fmt.Println(cars)
}