package main

import "fmt"

func sumOf(f func(int) int, n, m int) int {
	a := 0
	for ; n <= m; n++ {
		a += f(n)
	}
	return a
}

func main() {
	square := func(x int) int { return x * x }
	fmt.Println(square(10))
	a := func(x int) int { return x * x }(10)
	fmt.Println(a)
	fmt.Println(sumOf(square, 1, 100))
	fmt.Println(sumOf(func(x int) int { return x * x }, 1, 100))
}
