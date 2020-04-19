package main

import (
	"fmt"
)

func main() {

	// var num int32 = 15

	// fmt.Printf("num=%v 类型：%T\n", num, num)
	// fmt.Println(unsafe.Sizeof(num)) // sizeof可以查看在内存中的占用空间

	// var n1 uint8 = 1
	// fmt.Printf("num=%v 类型：%T\n", n1, n1)

	// // int类型转换
	// var a1 int32 = 10
	// var a2 int64 = 21

	// fmt.Println(int64(a1) + a2) // 把a1转换为64位

	// var num int = 30
	// fmt.Printf("num=%v 类型：%T\n", num, num)
	// fmt.Println(unsafe.Sizeof(num)) // sizeof可以查看在内存中的占用空间

	// %d表示十进制 %b表示二进制 o%表示八进制 %x表示十六进制
	num := 9
	fmt.Printf("num=%v 类型：%T\n", num, num)
	fmt.Printf("num=%d 类型：%T\n", num, num)
	fmt.Printf("num=%b 类型：%T\n", num, num)
	fmt.Printf("num=%o 类型：%T\n", num, num)
	fmt.Printf("num=%x 类型：%T\n", num, num)
}
