package main

import "fmt"

func main() {
	// Arrays
	//var fruitArr [2] string

	// Assign values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	// declare and assign
	fruitArr := [2]string{"Apple", "Oranges"}

	//fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	//using slices
	fruitSlice := []string{"Apple", "Oranges", "Grapes", "Cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}
