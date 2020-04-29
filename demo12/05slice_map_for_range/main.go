package main

import "fmt"

func main() {

	// var userinfo = []string{"张三", "李四"}

	var userinfo = make([]map[string]string, 3, 3)

	// fmt.Println(userinfo[0]) // map[]

	// fmt.Println(userinfo[0] == nil) // true

	if userinfo[0] == nil {
		userinfo[0] = make(map[string]string)
		userinfo[0]["username"] = "张三"
		userinfo[0]["age"] = "20"
		userinfo[0]["height"] = "180cm"
	}

	if userinfo[1] == nil {
		userinfo[0] = make(map[string]string)
		userinfo[0]["username"] = "李四"
		userinfo[0]["age"] = "22"
		userinfo[0]["height"] = "180cm"
	}

	if userinfo[1] == nil {
		userinfo[0] = make(map[string]string)
		userinfo[0]["username"] = "王五"
		userinfo[0]["age"] = "22"
		userinfo[0]["height"] = "180cm"
	}

	for _, v := range userinfo {
		// fmt.Println(v)
		for key, value := range v {
			fmt.Printf("key是%v -- value是%v\n", key, value)
		}
	}
}
