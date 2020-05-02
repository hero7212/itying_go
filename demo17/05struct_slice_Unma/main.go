package main

import (
	"encoding/json"
	"fmt"
)

// Student
type Student struct {
	Id     int
	Gender string
	Name   string
}

// Class
type Class struct {
	Title    string
	Students []Student
}

func main() {

	str := `{"Title":"001班","Students":[{"Id":0,"Gender":"男","Name":"stu_0"},{"Id":1,"Gender":"男","Name":"stu_1"},{"Id":2,"Gender":"男","Name":"stu_2"},{"Id":3,"Gender":"男","Name":"stu_3"},{"Id":4,"Gender":"男","Name":"stu_4"},{"Id":5,"Gender":"男","Name":"stu_5"},{"Id":6,"Gender":"男","Name":"stu_6"},{"Id":7,"Gender":"男","Name":"stu_7"},{"Id":8,"Gender":"男","Name":"stu_8"},{"Id":9,"Gender":"男","Name":"stu_9"}]}`

	var c = &Class{}
	err := json.Unmarshal([]byte(str), c)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v \n", c)

		fmt.Printf("%#v \n", c.Title)
	}
}
