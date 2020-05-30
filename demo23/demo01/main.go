package main

import (
	"fmt"
	"os"
)

func main() {
	// 只读打开main.go文件
	file, err := os.Open("./main.go")
	defer file.Close() // 必须得关闭文件流
	if err != nil {
		fmt.Println(err)
		return
	}
	// 操作文件
	fmt.Println(file) // &{0xc00007e780}
}
