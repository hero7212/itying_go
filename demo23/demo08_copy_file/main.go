package main

import (
	"fmt"
	"io/ioutil"
)

func copy(srcFileName string, dstFileName string) (err error) {
	byteStr, err1 := ioutil.ReadFile(srcFileName)
	if err1 != nil {
		return err1
	}

	err2 := ioutil.WriteFile(dstFileName, byteStr, 0666)
	if err2 != nil {
		return err2
	}
	return nil
}

func main() {
	src := "E:/图片/nude/李晴/23_副本.jpg"
	dst := "D:/workspace/private/learn/Golang/ITYing/GO_DEMO/demo23/demo08_copy_file/2.jpg"
	err := copy(src, dst)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("复制文件成功")

}
