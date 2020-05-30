package main

import (
	"fmt"
	"os"
)

func main() {
	// err := os.Remove("aaa.txt")

	// err := os.Remove("aaa")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	err := os.RemoveAll("dir3")

	if err != nil {
		fmt.Println(err)
	}
}
