package main

import (
	"fmt"
	"io/ioutil"
)

// 文件小，可以用这个
func main() {
	byteStr, err := ioutil.ReadFile("C:/Users/29149/Desktop/涛涛的面试.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteStr))
}
