package main

import "fmt"

type Person struct {
	string
	int
}

func main() {

	p := Person{
		"张三",
		20,
	}
	fmt.Println(p)
}
