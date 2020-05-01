package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return 5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(y int) {
		// x=0
		// fmt.Println(y)
		y++
	}(x)

	// fmt.Println("a=", x)
	return 5
}

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
}
