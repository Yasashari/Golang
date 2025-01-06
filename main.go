package main

import (
	"fmt"
	"strconv"
)

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

	firstname := "John"
	Lastname := "Doe"
	fmt.Printf("%s %s\n", firstname, Lastname)

	//Right shift Operation
	var a uint8 = 20
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	result := a >> 2 // Remove the last to bits -> 101 or 5
	fmt.Println(strconv.FormatUint(uint64(result), 2))
	fmt.Println(result)

	//Left shift Operation
	a = a << 2 // Append two zero bits -> 10100
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	//Set a bit given position

	a = 5 | (1 << 1) // 101 -> 111
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	// Clear a bit given position
	a = a &^ (1 << 1) // 111 -> 101
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	// Toggle a bit given position
	a = a ^ (1 << 0) // 101 -> 100
	fmt.Println(strconv.FormatUint(uint64(a), 2))

}
