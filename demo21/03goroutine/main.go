package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 计算机上面cpu个数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNums", cpuNum)

	// 可以自己设置使用多个cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
