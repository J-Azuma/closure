package main

import "fmt"

func foo(n int) func(int) int {
	return func(x int) int { return x * n }
}

func main() {
	foo10 := foo(10)
	foo20 := foo(20)
	fmt.Println(foo10(1))
	fmt.Println(foo20(20))
}
