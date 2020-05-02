package main

import (
	"encoding/json"
	"fmt"
)

// 结构体标签
type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Sno    string `json:"sno"`
}

func main() {
	var s1 = Student{
		ID:     21,
		Gender: "男",
		Name:   "李四",
		Sno:    "s0001",
	}
	fmt.Printf("%#v \n", s1)

	jsonByte, _ := json.Marshal(s1)
	jsonStr := string(jsonByte)
	fmt.Printf("%v \n", jsonStr) // {"id":21,"gender":"男","name":"李四","sno":"s0001"}
}
