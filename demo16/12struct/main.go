package main

import "fmt"

type User struct {
	Username string
	Password string
	Address  // 嵌套匿名结构体
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

	fmt.Printf("%#v \n", u) // main.User{Username:"朱莉", Password:"1234567", Address:main.Address{Name:"张敏", Phone:"13813818438", City:"南京"}}

	u.City = "上海" // 当访问结构体成员的时候，会先在结构体中找，找不到，就往下从嵌套结构体中找

	fmt.Printf("%#v \n", u) // main.User{Username:"朱莉", Password:"1234567", Address:main.Address{Name:"张敏", Phone:"13813818438", City:"上海"}}

	fmt.Println(u.Address.Phone) // 13813818438
	fmt.Println(u.Phone)         // 13813818438
}
