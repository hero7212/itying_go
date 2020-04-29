package main

import "fmt"

func main() {

	// var userinfo = make(map[string]string)
	// userinfo["username"] = "张三"
	// userinfo["hobby"] = "睡觉"

	var userinfo = make(map[string][]string)
	userinfo["hobby"] = []string{
		"吃饭",
		"睡觉",
		"敲代码",
	}

	userinfo["work"] = []string{
		"php",
		"golang",
		"前端",
	}
	fmt.Println(userinfo)

	for _, v := range userinfo {
		for key, value := range v {
			fmt.Println(key, value)
		}
	}
}
