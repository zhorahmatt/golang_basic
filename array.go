package main

import (
	"fmt"
)

func main()  {
	var names [4]string

	names[0] = "Gempi"
	names[1] = "Yanto"
	names[2] = "Borik"
	names[3] = "Totok"
	//names[4] = "Borris" // it will show error coz out of index

	fmt.Println(names[0], names[1], names[2], names[3])

	//inisialisasi array dengan nilai sebelumnya
	//horizontal way
	var fruits = [4]string{"apple", "pineapple", "durian"} //it can assign less than total index

	//vertical way to declare array
	// var fruits = [4]string{
	// 	"apple",
	// 	"pineapple",
	// 	"durian"
	// }

	fmt.Println(fruits)
	fmt.Println(len(fruits))

	var numbers = [...]int{2,3,4,5,6} //inisialisasi tanpa jumlah index sebelumnya
	
}