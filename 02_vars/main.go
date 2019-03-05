package main

import "fmt"

func main() {
	// main types
	// string
	// bool
	// int int8 int16 int 32 int 64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte
	// rune
	// float32 float64
	// complex64 complex128

	// using var
	//var name = "Newton"

	var age = 23

	const isCool = false
	var size = 2.3

	// short hand method [ should always be in functions ]
	//name := "Newton\n"
	//email := "mymbugua@gmail.com\n"

	name, email := "Newton", "mymbugua@gmail.com"

	fmt.Println(name, email, age, isCool, size)

	fmt.Printf("%T\n", isCool)
}
