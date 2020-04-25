package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "123456"
	fmt.Printf("%v--%T \n", str, str)

	// num, _ := strconv.ParseInt(str, 10, 64) // 转换不了 就 返回 0
	// fmt.Printf("%v--%T \n", num, num)

	// num, _ := strconv.ParseFloat(str, 64) // 转换不了 就 返回 0
	// fmt.Printf("%v--%T \n", num, num)

	b, _ := strconv.ParseBool("xxxx") // 转换不了 就 返回 false
	fmt.Printf("%v--%T \n", b, b)
}
