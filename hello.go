package main

import "fmt"

type X interface{}
type A struct{}
type B struct{}

func main() {
	fmt.Println("Hello")
	var x X
	x = A{}
	b := x.(B) // will panic, good
	fmt.Println(b)
}
