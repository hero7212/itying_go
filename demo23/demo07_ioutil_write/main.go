package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "hello golang"
	err := ioutil.WriteFile("D:/workspace/private/learn/Golang/ITYing/GO_DEMO/demo23/demo07_ioutil_write/test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}
