package main

import (
	"fmt"
)

func main()  {
	number := 12

	//single condition
	if number == 12 {
		fmt.Println("sama")
	}

	//with else condition
	if number > 15 {
		fmt.Println("nilai lebih besar dari 15")
	} else {
		fmt.Println("nilai kurang dari 15")
	}

	//more complex
	if number == 15 {
		fmt.Println("nilai sama 15")
	} else if number > 15 {
		fmt.Println("nilai besar dari 15")
	} else {
		fmt.Println("nilai kurang dari 15")
	}

	//switch case conditional
	var score = 33

	switch score {
	case 30,31,32,33:
		fmt.Println("score is 30")
	case 34,35:
		fmt.Println("score is 32")
		fallthrough
	default:
		{
			fmt.Println("score is not detected")
			fmt.Println("try again !!!")
		}
	}
}