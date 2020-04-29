package main

import "fmt"

func main() {

	// var userinfo = []string{"张三", "李四"}

	var userinfo1 = make(map[string]string)

	userinfo1["username"] = "张三"
	userinfo1["age"] = "20"
	userinfo2 := userinfo1
	fmt.Println(userinfo1)
	fmt.Println(userinfo2)
	userinfo2["username"] = "李四"
	fmt.Println(userinfo1)
	fmt.Println(userinfo2)
}
