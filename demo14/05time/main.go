package main

import (
	"fmt"
	"time"
)

/*
当前时间戳
*/
func main() {
	var str = "2020-05-01 11:17:27"
	var tmp = "2006-01-02 15:04:05"
	timeObj, _ := time.ParseInLocation(tmp, str, time.Local)
	fmt.Println(timeObj)

	fmt.Println(timeObj.Unix())
}
