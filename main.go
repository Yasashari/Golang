package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var myString string
	fmt.Println(myString)

	myString = "Hello World!"

	var myInt int

	fmt.Println(myInt)

	myInt = 42

	fmt.Println(myInt)

	var myBool bool
	fmt.Println(myBool)

	myBool = true
	fmt.Println(myBool)

	var myByte byte = 'A'
	fmt.Println(myByte)

	var myRune rune = 'B'
	fmt.Println(myRune)

	var myComplex complex128
	myComplex = 1 + 2i
	fmt.Println(myComplex)
}
