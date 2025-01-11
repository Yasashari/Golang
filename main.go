// package main

// import (
// 	"errors"
// 	"fmt"
// )

// type student struct {
// 	firstName string
// 	lastName  string
// }

// func main() {
// 	// a := 10
// 	// increment(&a)
// 	// fmt.Println(a)

// 	s := student{
// 		firstName: "John",
// 		lastName:  "Smith",
// 	}

// 	previousLastName, err := updateLastName(&s, "")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(previousLastName)

// }

// // func increment(x *int) {
// // 	*x++
// // }

// func updateLastName(s *student, newLastName string) (*string, error) {

// 	if newLastName == "" {
// 		return nil, errors.New("empty new last name")
// 	}

// 	previous := s.lastName
// 	s.lastName = newLastName
// 	return &previous, nil
// }

// package main

// import (
// 	"fmt"
// )

//func main() {

// s := []int{1, 2}

// fmt.Println(s)

// somethingSlice(s)

// fmt.Println(s)

// 	myMap := map[int]string{
// 		0: "one",
// 		1: "two",
// 	}

// 	fmt.Println(myMap)

// 	somethingMap(myMap)

// 	fmt.Println(myMap)
// }

// func somethingSlice(s []int) {
// 	//s[0] = 100
// 	s = append(s, 100)
// }

// func somethingMap(m map[int]string) {
// 	delete(m, 0)
// }

// methods

// package main

// import (
// 	"fmt"
// )

// type Person struct {
// 	Name string
// 	Age  string
// }

// type Circle struct {
// 	Radius float64
// }

// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }

// func (c Circle) Area() float64 {
// 	return 3.14 * c.Radius * c.Radius
// }

// func (p Person) GetDetails() string {
// 	return fmt.Sprintf("Name: %s, Age: %s", p.Name, p.Age)
// }

// func (r *Rectangle) Area() float64 {
// 	if r == nil {
// 		return 0
// 	}
// 	return r.Width * r.Height
// }

// func (r *Rectangle) SetWidth(w float64) {
// 	r.Width = w
// }

// func updateWidth(r *Rectangle, w float64) {
// 	r.SetWidth(w)
// }

// func main() {
// 	p := Person{
// 		Name: "John",
// 		Age:  "30",
// 	}

// 	fmt.Println(p.GetDetails())

// 	rect := Rectangle{
// 		Width:  10,
// 		Height: 5,
// 	}

// 	fmt.Println("Area of Rectangle:", rect.Area())
// 	rect.SetWidth(20)
// 	fmt.Println("Area of Rectangle:", rect.Area())

// 	var rect2 *Rectangle

// 	fmt.Println("Area of Rectangle2:", rect2.Area())

// 	updateWidth(&rect, 20)

// 	fmt.Println("Area of Rectangle3:", rect.Area())
// }

// package main

// import "fmt"

// type Calculator struct {
// }

// func (c *Calculator) Add(a, b int) int {
// 	return a + b
// }

// func AddFunciton(a, b int) int {
// 	return a + b
// }

// func ArithmeticOperation(f func(int, int) int, a, b int) int {
// 	return f(a, b)
// }

// func main() {
// 	fmt.Println(AddFunciton(10, 15))
// 	fmt.Println(ArithmeticOperation(AddFunciton, 10, 15))
// 	c := Calculator{}
// 	fmt.Println(ArithmeticOperation(c.Add, 10, 15))
// }

// package main

// import "fmt"

// type Printer struct{}

// func (p *Printer) Print(m string) {
// 	fmt.Println(m)
// }

// func main() {
// 	printer := &Printer{}

// 	printFunction := printer.Print

// 	printFunction("Hello World!")
// }

// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) GetDetails() string {
// 	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
// }

// func main() {
// 	p := Person{
// 		Name: "John",
// 		Age:  30,
// 	}

// 	f1 := Person.GetDetails
// 	fmt.Println(f1(p))

// }

//package main

//import "fmt"

//type Age int

// type Human struct {
// 	Age int
// }

//type Student Human

// func sum(a Age) {
// 	fmt.Printf("Age: %d\n", a)
// }

// func (h Human) printAge() {
// 	fmt.Println(h.Age)
// }

//type Student Human

//func main() {
// var young Age = 10
// var old Age = 100
// fmt.Println(young + old)
// sum(young)

//var s = Student{Age: 10}
//fmt.Println(s.Age)
//var newYoung int = 10

//sum(newYoung)
//}

// iota

// package main

// import "fmt"

// type Size int

// const (
// 	ExtraSmall Size = iota
// 	Small
// 	Medium
// 	Large
// 	ExtraLarge
// )

// func main() {
// 	fmt.Println(ExtraSmall)
// 	fmt.Println(Small)
// 	fmt.Println(Medium)
// 	fmt.Println(Large)
// 	fmt.Println(ExtraLarge)
// }

// package main

// import "fmt"

// type Engine struct {
// 	HorsePower int
// }

// type GPS struct {
// 	Model string
// }

// type Car struct {
// 	Model string
// 	Engine
// 	GPS
// }

// func (e *Engine) Start() {
// 	fmt.Println("Engine started")
// }

// func (c *Car) Drive() {
// 	fmt.Printf("Driving my %s...\n", c.Model)
// }

// type Car struct {
// 	Model string
// 	Engine
// }

//func main() {
// myCar := Car{
// 	Model: "Mustang",
// 	Engine: Engine{
// 		HorsePower: 300,
// 	},
// }
// fmt.Println(myCar.Model, myCar.HorsePower)

// myCar.Start()
// myCar.Drive()

// 	myCar := Car{
// 		Model: "Mustang",
// 		Engine: Engine{
// 			HorsePower: 300,
// 		},
// 		GPS: GPS{
// 			Model: "Garmin",
// 		},
// 	}

// 	fmt.Println(myCar.Model, myCar.HorsePower, myCar.GPS.Model)
// }

package main

import "fmt"

type Shape interface {
	Area() float64
}

type Object interface {
	Name() string
	Shape
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Name() string {
	return "Rectangle"
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

func printObject(o Object) {
	fmt.Printf("Area:%f , Name: %s\n", o.Area(), o.Name())
}
func main() {
	r := Rectangle{
		Width:  10,
		Height: 5,
	}

	c := Circle{
		Radius: 5,
	}

	shapes := []Shape{r, c}

	for _, shape := range shapes {
		//fmt.Printf("%f\n", shape.Area())
		printArea(shape)

	}

	printObject(r)
}
