package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	1

	2

	3
*/
func main() {
	file, err := os.OpenFile("D:/workspace/private/learn/Golang/ITYing/GO_DEMO/demo23/demo06_write/test.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(file)

	// writer.WriteString("你好golang")

	for i := 0; i < 10; i++ {
		writer.WriteString("你好golang" + strconv.Itoa(i) + "\r\n") // \r\n才是换行
	}

	writer.Flush() // 把缓存的内容也写入， 这个不能省略
}
