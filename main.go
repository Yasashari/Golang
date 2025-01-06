package main

import "fmt"

func main() {
	var a []int
	fmt.Println(a)

	a = []int{}
	fmt.Println(a)

	b := make([]int, 5)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	c := make([]int, 5, 10)
	fmt.Println(len(c))
	fmt.Println(cap(c))

	e := []int{1, 2, 3}
	f := []int{4, 5, 6}
	g := append(e, f...)

	fmt.Println(g)

}
