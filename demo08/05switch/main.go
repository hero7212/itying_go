package main

import "fmt"

func main() {

	var score = "A"
	switch score {
	case "A", "B", "C":
		fmt.Println("及格")
	case "D":
		fmt.Println("不及格")
	}
}
