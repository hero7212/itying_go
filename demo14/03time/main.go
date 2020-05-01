package main

import (
	"fmt"
	"time"
)

/*
当前时间戳
*/
func main() {
	timeObj := time.Now()
	// 时间戳
	unixtime := timeObj.Unix()
	fmt.Println("当前毫秒时间戳：", unixtime)

	unixNatime := timeObj.UnixNano()
	fmt.Println("当前纳秒时间戳：", unixNatime)
}
