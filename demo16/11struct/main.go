package main

import "fmt"

type User struct {
	Username string
	Password string
	Sex      string
	Age      int
	Address  Address // 嵌套Address结构体
}

type Address struct {
	Name  string
	Phone string
	City  string
}

func main() {

	var u User

	u.Username = "朱莉"
	u.Password = "1234567"
	u.Address.Name = "张敏"
	u.Address.Phone = "13813818438"
	u.Address.City = "南京"

	fmt.Printf("%#v", u)
}
