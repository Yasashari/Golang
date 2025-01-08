// package main

// import "fmt"

// // A funciton that returns a function
// func createAdder(x int) func(int) int {
// 	return func(y int) int {
// 		return x + y
// 	}
// }

// func main() {
// 	addFive := createAdder(5)
// 	fmt.Println(addFive(10))
// }

// package main

// import "fmt"

// func double(x int) int {
// 	return x * 2
// }

// func triple(x int) int {
// 	return x * 3
// }

// func apply(f func(int) int, x int) int {
// 	return f(x)
// }

// func main() {
// 	result := apply(double, 5)
// 	fmt.Println(result)

// 	resultTriple := apply(triple, 5)
// 	fmt.Println(resultTriple)

// 	fiveTime := func(x int) int {
// 		return x * 5
// 	}

// 	resultFiveTime := apply(fiveTime, 5)
// 	fmt.Println(resultFiveTime)
// }

// package main

// import "fmt"

// type multiplier func(int) int

// func multiplyBy(m int) multiplier {
// 	return func(i int) int {
// 		return i * m
// 	}
// }

// func main() {
// 	multiplyByTwo := multiplyBy(2)
// 	mutlyByThree := multiplyBy(3)

// 	result1 := multiplyByTwo(5)
// 	result2 := mutlyByThree(3)

// 	fmt.Println(result1, result2)
// }

// package main

// import "fmt"

// type operration func(int, int) int

// func arithmeticOperation(op string) operration {
// 	switch op {
// 	case "add":
// 		return func(i1, i2 int) int {
// 			return i1 + i2
// 		}

// 	case "subtract":
// 		return func(i1, i2 int) int {
// 			return i1 - i2
// 		}

// 	default:
// 		return func(i1, i2 int) int {
// 			return i1 * i2
// 		}
// 	}
// }

// func main() {
// 	var perform operration
// 	perform = arithmeticOperation("add")
// 	result := perform(2, 3)
// 	fmt.Println(result)
// }

// package main

// import "fmt"

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5}

// 	squared := sliceOperation(numbers, func(i int) int {
// 		return i * i

// 	})

// 	fmt.Println(squared)
// }

// func sliceOperation(s []int, op func(int) int) []int {
// 	result := make([]int, len(s))
// 	for i, v := range s {
// 		result[i] = op(v)
// 	}
// 	return result
// }

//closures
// package main

// import "fmt"

// func counter() func() int {
// 	count := 0

// 	return func() int {
// 		count++
// 		return count
// 	}
// }

// func main() {
// 	increment := counter()
// 	fmt.Println(increment())
// 	fmt.Println(increment())
// }

//anonymous functions

// package main

// import "fmt"

// func main() {

// 	func() {
// 		fmt.Println("Anonymous function")
// 	}()

// 	func(msg string) {
// 		fmt.Println("Message", msg)
// 	}("Hello World")
// }

//defer funcitons

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	message := "Hello,"
// 	greetingFn := func(name string) {
// 		fmt.Println(message, name)

// 	}

// 	defer greetingFn("Alice")
// 	defer greetingFn("Bob")

// 	message = "Hi,"
// }

// Pointers

// var pointerToInt *int
// var pointerToString *string
// var pointerToBool *bool

// age := 10
//pointerToInt = &age

//ageValue := *pointerToInt

package main

import "fmt"

func main() {

	var intPtr *int

	age := 10
	intPtr = &age
	fmt.Println(intPtr)
	fmt.Println(*intPtr)

	name := "Yasas"
	namePtr := &name
	fmt.Println(namePtr)
	fmt.Println(*namePtr)

}
