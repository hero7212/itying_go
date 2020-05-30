package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(srcFileName string, dstFileName string) (err error) {
	sFile, err1 := os.Open(srcFileName)
	defer sFile.Close()
	dFile, err2 := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)
	defer dFile.Close()
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	var tempSlice = make([]byte, 1280)
	for {
		// 读取
		n1, err := sFile.Read(tempSlice)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		// 写入
		if _, err := dFile.Write(tempSlice[:n1]); err != nil {
			return err
		}

	}
	return nil
}

func main() {
	src := "E:/图片/nude/李晴/17.jpg"
	dst := "D:/workspace/private/learn/Golang/ITYing/GO_DEMO/demo23/demo09_copy_file/2.jpg"
	err := copyFile(src, dst)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("复制文件成功")

}
