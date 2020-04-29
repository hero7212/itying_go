package main

import "fmt"

func sumFn(x int, y int) int {
	sum := x + y
	return sum
}

func subFn(x, y int) int {
	sub := x - y
	return sub
}

func sumFn1(x ...int) int {
	var sum = 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func sumFn2(x int, y ...int) int {
	var sum = x
	for _, v := range y {
		sum += v
	}
	return sum
}

func main() {
	// sum1 := sumFn(12, 3)
	// fmt.Println(sum1)

	// sum2 := sumFn(15, 5)
	// fmt.Println(sum2)

	// a, b := 20, 2
	// sub1 := subFn(a, b)
	// fmt.Println(sub1)

	sum1 := sumFn1(1, 2, 3, 4, 5)
	fmt.Println(sum1)

	sum2 := sumFn2(100, 1, 2, 3, 4, 5)
	fmt.Println(sum2)
}
