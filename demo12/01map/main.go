package main

import "fmt"

func main() {
	// var userinfo = make(map[string]string)
	// userinfo["username"] = "张三"
	// userinfo["age"] = "20"
	// userinfo["sex"] = "男"
	// fmt.Println(userinfo)
	// fmt.Println(userinfo["username"])

	// var userinfo = map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// 	"sex":      "男",
	// }
	// fmt.Println(userinfo)

	userinfo := map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
	}
	fmt.Println(userinfo)
}
