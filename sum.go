package main

import "fmt"

func sumOfInteger(n, m int) int {
	a := 0
	for ; n <= m; n++ {
		a += n
	}
	return a
}

func Shiguma(x int) int {
	return x * (x + 1) / 2
}

func square(x int) int {
	return x * x
}

func cube(x int) int {
	return x * x * x
}

func sumOfSquare(n, m int) int {
	a := 0
	for ; n <= m; n++ {
		a += square(n)
	}
	return a
}

func sumOfCube(n, m int) int {
	a := 0
	for ; n <= m; n++ {
		a += cube(n)
	}
	return a
}

func sumOf(f func(int) int, n, m int) int {
	a := 0
	for ; n <= m; n++ {
		a += f(n)
	}
	return a
}

func main() {
	fmt.Println(sumOfInteger(1, 100))
	fmt.Println(Shiguma(100))
	fmt.Println(sumOfSquare(1, 100))
	fmt.Println(sumOfCube(1, 100))
	fmt.Println(sumOf(square, 1, 100))
	fmt.Println(sumOf(cube, 1, 100))
}
