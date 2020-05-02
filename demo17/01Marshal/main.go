package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	name   string
	Sno    string
}

func main() {

	var s1 = Student{
		ID:     21,
		Gender: "男",
		name:   "李四", // 小写不能被转换，应该用Name
		Sno:    "s0001",
	}
	fmt.Printf("%#v \n", s1)

	jsonByte, _ := json.Marshal(s1)

	jsonStr := string(jsonByte)

	fmt.Printf("%v \n", jsonStr) // {"ID":21,"Gender":"男","name":"李四","Sno":"s0001"}
}
