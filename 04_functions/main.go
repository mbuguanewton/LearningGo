package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

// another function

func getSum(num1, num2 int) int {
	return num1 + num2
}


func main() {
	fmt.Println(greeting("Newton"))
	fmt.Println(getSum(4, 6))
}
