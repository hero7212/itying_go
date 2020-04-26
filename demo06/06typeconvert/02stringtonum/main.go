package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		var i int = 20
		var f float64 = 12.456
		var t bool = true
		var b byte = 'a'

		str1 := fmt.Sprintf("%d", i)
		fmt.Printf("%v--类型：%T \n", str1, str1)

		str2 := fmt.Sprintf("%.2f", f) // 四舍五入 保留2位
		fmt.Printf("%v--类型：%T \n", str2, str2)

		str3 := fmt.Sprintf("%t", t)
		fmt.Printf("%v--类型：%T \n", str3, str3)

		str4 := fmt.Sprintf("%c", b)
		fmt.Printf("%v--类型：%T \n", str4, str4)
	*/

	// var i int = 20

	// str1 := strconv.FormatInt(int64(i), 10) // 第一个参数必须是int64类型
	// fmt.Printf("值：%v--类型：%T \n", str1, str1)

	// var f float32 = 20.23

	// str2 := strconv.FormatFloat(float64(f), 'f', 4, 64) // 第一个参数必须是int64类型
	// fmt.Printf("值：%v--类型：%T \n", str2, str2)

	// s3 := strconv.FormatBool(true)
	// fmt.Printf("值：%v--类型：%T \n", s3, s3)

	a := 'a'
	str4 := strconv.FormatUint(uint64(a), 10)
	fmt.Printf("值：%v--类型：%T \n", str4, str4) // 97 string
}
