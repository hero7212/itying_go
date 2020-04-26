package main

import "fmt"

func main() {

	var arr = [3][2]string{
		{"背景", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}

	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// fmt.Println(arr[0][1])
}
