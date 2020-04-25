package main

import "fmt"

func main() {
	var str = "你好 golang"

	for k, v := range str {
		fmt.Printf("key=%v val=%c \n", k, v)
	}
}
