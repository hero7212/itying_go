package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1、打开文件
	file, err := os.Open("C:/Users/29149/Desktop/涛涛的面试.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	// 2、读取文件内容
	var strSlice []byte
	var tempSlice = make([]byte, 128)

	for {
		n, err := file.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取失败")
			return
		}
		// fmt.Printf("读取到了%v个字节/n", n)
		strSlice = append(strSlice, tempSlice[:n]...)
	}

	fmt.Println(string(strSlice))
}
