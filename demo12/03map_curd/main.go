package main

import "fmt"

func main() {

	// var userinfo = map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// }
	// userinfo["sex"] = "男"
	// fmt.Println(userinfo)

	// v, ok := userinfo["age"]
	// fmt.Println(v, ok)

	// v, ok = userinfo["xxx"]
	// fmt.Println(v, ok) // 空 和 false

	var userinfo = map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
		"height":   "180cm",
	}
	fmt.Println(userinfo)

	delete(userinfo, "sex")
	fmt.Println(userinfo)
}
