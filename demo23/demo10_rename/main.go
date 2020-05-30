package main

import (
	"fmt"
	"os"
)

func main() {

	err := os.Rename("./1.txt", "./2.txt")

	if err != nil {
		fmt.Println(err)
	}
}
