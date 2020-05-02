package main

import (
	"fmt"
	"itying/calc"    // 前面加个 _  表示不编译
	T "itying/tools" // T 是给包起的别名
)

// init 函数优先于 main函数	  	init是go自带的函数
func init() {
	fmt.Println("main init...")
}

func main() {
	sum := calc.Add(10, 2)
	fmt.Println(sum)

	fmt.Println(calc.Age)
	b := T.Mul(3, 5)
	fmt.Println(b)
}
