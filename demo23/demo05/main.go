package main

import (
	"fmt"
	"os"
)

/*
	1

	2

	3
*/
func main() {
	file, err := os.OpenFile("D:/workspace/private/learn/Golang/ITYing/GO_DEMO/demo23/demo05/test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	// 写入文件
	// for i := 0; i < 10; i++ {
	// 	file.WriteString("直接写入的字符串数据" + strconv.Itoa(i) + "\r\n") // \r\n才是换行
	// }

	var str = "直接写入的字符串数据"
	file.Write([]byte(str))
}
