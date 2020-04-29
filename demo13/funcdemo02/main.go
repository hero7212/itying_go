package main

import "fmt"

func sumFn(x int, y int) int {
	return x + y
}

func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func calc1(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	sum1 := sumFn(12, 3)
	fmt.Println(sum1)

	a, b := calc(10, 2)
	fmt.Println(a, b)

	a, b = calc1(10, 2)
	fmt.Println(a, b)
}
