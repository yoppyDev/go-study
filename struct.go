package main

import (
	"fmt"
)

type Vertex struct {
	x, y int
}

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func Say() {
	fmt.Println("Hello")
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func (v Vertex) Area() int {
// ポインタの値を変更すると、元の値も変更される
// func (v *Vertex) Area() int {
	v.x = v.x * 10
	v.y = v.y * 10
	return v.x * v.y
}

func do (i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v is int\n", v)
	case string:
		fmt.Printf("%v is string\n", v)
	default:
		fmt.Printf("%v is unknown\n", v)
	}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Area())
	fmt.Println(v.x)
	fmt.Println(v.y)

	var jon Human = &Person{"Jon"}
	jon.Say()

	do(10)
	do("10")
	do(nil)
}