package main

import (
	"errors"
	"fmt"
)

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func myFn() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("给管理员发送邮件：", err)
		}
	}()
	err := readFile("main.go")
	if err != nil {
		panic(err)
	}
}

func main() {
	myFn()
	fmt.Println("继续执行")
}
