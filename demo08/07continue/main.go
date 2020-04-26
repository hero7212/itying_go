package main

import "fmt"

func main() {
	count := 10
label1:
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if j == 3 {
				continue label1
			}
			fmt.Println("j--", j)
		}
		fmt.Println("i--", i)
	}
}
