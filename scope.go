package main

import "fmt"

var x int = 10

func foo() {
	fmt.Println(x)
}

func foo1() {
	x := 100
	fmt.Println("local var x = ", x)
	foo()
}

func main() {
	foo()
	foo1()
}
