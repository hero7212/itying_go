package main

import (
	"fmt"
	// "github.com/shopspring/decimal"
)

func main() {

	// var a float32 = 3.12
	// fmt.Printf("值：%v--%f, 类型%T\n", a, a, a)
	// fmt.Println(unsafe.Sizeof(a))

	// var a float64 = 3.12
	// fmt.Printf("值：%v--%f, 类型%T\n", a, a, a)
	// fmt.Println(unsafe.Sizeof(a))

	// var c float64 = 3.141592654432
	// fmt.Printf("值：%v--%f--%.2f \n", c, c, c)

	// f1 := 3.1313144

	// fmt.Printf("%f--%T", f1, f1)

	// f2 := 3.14e2 // 3.14*10的2次方
	// fmt.Printf("%f--%T", f2, f2)

	// f4 := 1129.6

	// fmt.Println(f4 * 100) //精度丢失

	// m1 := 8.2
	// m2 := 3.8
	// fmt.Println(m1 - m2) // 精度丢失

	// d1 := decimal.NewFromFloat(m1).Add(decimal.NewFromFloat(m2))
	// fmt.Println(d1)

	// a := 10
	// b := float64(a)
	// fmt.Printf("a的类型是%T,b的类型是%T", a, b)

	// var a1 float32 = 23.4
	// a2 := float64(a1)
	// fmt.Printf("a1的类型是%T,a2的类型是%T", a1, a2)

	var c1 float32 = 23.95
	c2 := int(c1)
	fmt.Printf("c2的值是%v,c2的类型是%T", c2, c2)
}
