package main

import (
	"fmt"
	"strings"
)

func main() {
	// var str1 string = "a"
	// var str2 = "b"
	// str3 := "c"

	// fmt.Printf("%v--%T\n", str1, str1)
	// fmt.Printf("%v--%T\n", str2, str2)
	// fmt.Printf("%v--%T\n", str3, str3)

	// str1 := "this \r is str1"
	// fmt.Println(str1)

	// str2 := "D:\\Go" // 输出 \
	// fmt.Println(str2)

	// str3 := "D:\"Go" // 输出 ""
	// fmt.Println(str3)

	// str1 := "你好"
	// fmt.Println(len(str1))	// 6  一个汉字3个字节

	// str1 := "你好"
	// str2 := "goland"
	// fmt.Println(str1 + str2)

	// str1 := "你好"
	// str2 := "goland"
	// str3 := fmt.Sprintf("%v-%v", str1, str2)
	// fmt.Println(str3)

	// str1 := "你好啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊阿" +
	// 	"住啦啦啦啦啦啦啦" +
	// 	"aaaa"
	// fmt.Println(str1)

	// str1 := "aaa-bbb-ccc"
	// arr := strings.Split(str1, "-")
	// fmt.Println(arr)
	// str2 := strings.Join(arr, "*")
	// fmt.Println(str2)

	// arr := []string{"a", "b"}
	// fmt.Println(arr)
	// str3 := strings.Join(arr, "-")
	// fmt.Println(str3)

	// 包含
	// str1 := "this is str"
	// str2 := "this"
	// flag := strings.Contains(str1, str2)
	// fmt.Println(flag)

	// 前缀
	// str1 := "this is str"
	// str2 := "this"
	// flag := strings.HasPrefix(str1, str2)
	// fmt.Println(flag)

	// 出现的位置 正着数 找不到返回-1
	// str1 := "this is str"
	// str2 := "str"
	// num := strings.Index(str1, str2)
	// fmt.Println(num)

	// 出现的位置 倒着数 找不到返回-1
	str1 := "this is str"
	str2 := "is"
	num := strings.LastIndex(str1, str2)
	fmt.Println(num)
}
