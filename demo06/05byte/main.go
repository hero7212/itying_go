package main

import (
	"fmt"
)

func main() {
	// 字符，
	// var a = 'a'
	// fmt.Printf("值是%v, 类型%T", a, a)	// 97 码值

	// var a = 'a'
	// fmt.Printf("值是%c, 类型%T", a, a) // a

	// str := "this"
	// fmt.Printf("值是%v, 原样输出%c, 类型%T", str[2], str[2], str[2]) // 105 i unit8

	// str := "你好go"
	// fmt.Println(unsafe.Sizeof(str)) // 字符串永远是16
	// fmt.Println(len(str))

	// var a = '国'
	// var s = fmt.Sprintf("%c", a)     // %c 转成 字符串
	// fmt.Printf("值是%v, 类型%T\n", a, a) // 汉字是utf8编码 22269  int32
	// fmt.Printf("值是%v, 类型%T\n", s, s) // 国 string

	// s := "你好 golang"
	// for i, v := range s {
	// 	fmt.Printf("下标是%v，%v - %c \n", i, v, v)
	// }

	// s1 := "big"
	// byteStr := []byte(s1)
	// byteStr[0] = 'p'
	// fmt.Println(string(byteStr))

	s1 := "你好golang"
	byteStr := []rune(s1)
	byteStr[0] = '大'
	fmt.Println(string(byteStr))
}
