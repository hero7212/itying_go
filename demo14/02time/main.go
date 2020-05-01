package main

import (
	"fmt"
	"time"
)

/*
2006 年
01 月
02 日
03 时 12小时制	写成15 表示 24小时制
04 分
05 秒

*/
func main() {
	timeObj := time.Now()
	str := timeObj.Format("2006-01-02 03:04:05")
	fmt.Println(str)
	str = timeObj.Format("2006/01/02 03:04:05")
	fmt.Println(str)
	// 24小时制
	str = timeObj.Format("2006/01/02 15:04:05")
	fmt.Println(str)
}
