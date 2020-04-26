package main

import "fmt"

func main() {

	var strSlice = []string{"php", "java", "nodejs", "golang"}
	l := len(strSlice)
	for i := 0; i < l; i++ {
		fmt.Println(strSlice[i])
	}

	for k, v := range strSlice {
		fmt.Println(k, v)
	}
}
