package main

import "fmt"

type Address struct {
	Name  string
	Phone int
}

// 空接口和类型断言的细节
func main() {

	var userinfo = make(map[string]interface{})
	userinfo["username"] = "张三"
	userinfo["age"] = 20
	userinfo["hobby"] = []string{"睡觉", "吃饭"}

	fmt.Println(userinfo["age"])
	fmt.Println(userinfo["hobby"])

	// fmt.Println(userinfo["hobby"][0]) // type interface {} does not support indexing

	var address = Address{
		Name:  "李四",
		Phone: 13813818438,
	}
	fmt.Println(address.Name) // 李四

	userinfo["address"] = address

	fmt.Println(userinfo["address"])

	// var name = userinfo["address"].Name // type interface {} is interface with no methods
	// fmt.Println(name)

	hobby2, _ := userinfo["hobby"].([]string)

	fmt.Println(hobby2[0])

	address2, _ := userinfo["address"].(Address)
	fmt.Println(address2.Name, address2.Phone)

}
