package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

type calcType = func(int, int) int

func calc(x, y int, cb calcType) int {

	return cb(x, y)
}

func main() {
	sum := calc(10, 5, add)
	fmt.Println(sum)

	j := calc(3, 4, func(x, y int) int {
		return x * y
	})
	fmt.Println(j)
}
