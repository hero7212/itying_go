package main

import "fmt"

func fn1(x int) {
	x = 10
}

func fn2(x *int) {
	*x = 40
}

func main() {

	var a = 5
	fn1(5)
	fmt.Println(a)

	fn2(&a)
	fmt.Println(a)

}
