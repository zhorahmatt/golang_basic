package main

import (
	"fmt"
)

func main()  {
	a := 20
	b := 4

	//add
	numberAdded := a + b
	fmt.Println(numberAdded)

	//substract
	numberMinus := a - b
	fmt.Println(numberMinus)

	//multiplies
	numberTimes := a * b
	fmt.Println(numberTimes)

	//divide
	numberDivide := a / b
	fmt.Println(numberDivide)

	//modules
	numberModulos := a % b
	fmt.Println(numberModulos)
}