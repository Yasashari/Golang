package main

import "fmt"

func main() {
	name := "Code & String"

	runeName := []rune(name)
	fmt.Println(runeName)

	byteName := []byte(name)
	fmt.Println(byteName)

	nameGrade := map[string]int{
		"foo":     10,
		"bar":     9,
		"foobard": 8,
		"x":       0,
	}

	gradeX, ok := nameGrade["x"]
	fmt.Println(gradeX, ok)

	//Anonimous struct

	guardian := struct {
		firstName string
		LastName  string
	}{
		firstName: "Alex",
		LastName:  "Theron",
	}

	fmt.Printf("%+v\n", guardian)

	age := 10

	if even := isEven(age); even {
		fmt.Println("Age is even")
	}

	//range for loop

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range a {
		fmt.Println(i, v)
	}

	ages := map[string]int{
		"Alex": 10,
		"John": 20,
	}

	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue outerLoop
			}
			fmt.Println(i, j)
		}

	}

	//Goto statement

	P := 9

	if P == 10 {
		goto end
	}

	fmt.Println(P)

end:
	fmt.Println("End")

	s := sum(1, 2, 3, 4, 5, 6, 10)
	fmt.Println(s)

	nums := []int{1, 2, 3, 4, 5, 6, 10}
	ss := sum(nums...)
	fmt.Println(ss)

}

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func isEven(n int) bool {
	return n&1 == 0
}
