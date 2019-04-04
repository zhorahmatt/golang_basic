package main

import (
	"fmt"
)

func main()  {
	//in go, there is only on syntax for looping.
	//just using for syntax to handle all looping condition
	
	//normal loop
	for i := 0; i < 5; i++ {
		fmt.Println("Angka ",i)
	}

	//looping with condition like while
	var number = 1;

	for number < 8 {
		fmt.Println("Angka ", number)
		number++
	}

	//looping without argumen
	var duck = 1

	for {
		fmt.Println("duck count :", duck)

		duck++

		if duck == 10 {
			break
		}
	}

	//label on looping
	outerloop:
	for iIndex := 0; iIndex < 5; iIndex++ {
		for jIndex := 0; jIndex < 5; jIndex++ {
			if iIndex == 3 {
				break outerloop
			}

			fmt.Print("matriks [", iIndex, "][", jIndex, "]", "\n")
		}
	}
}