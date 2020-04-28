package main

import "fmt"

func main() {

	userinfo := map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
	}
	for k, v := range userinfo {
		fmt.Println(k, v)
	}
}
