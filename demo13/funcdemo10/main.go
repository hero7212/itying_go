package main

import "fmt"

func fn1(n int) {

	if n > 0 {
		fmt.Println(n)
		n--
		fn1(n)
	}
}

func fn2(n int) int {
	if n > 1 {
		return n + fn2(n-1)
	} else {
		return 1
	}
}

func fn3(n int) int {
	if n > 1 {
		return n * fn3(n-1)
	} else {
		return 1
	}
}

func main() {
	fn1(10)
	fmt.Println("----------------")
	fmt.Println(fn2(100))
	fmt.Println(fn3(5))
}
