package main

import (
	"fmt"
	"time"
)

/*
当前时间戳
*/
func main() {

	unixTime := 1588303047

	timeObj := time.Unix(int64(unixTime), 0)

	str := timeObj.Format("2006-01-02 15-04-05")
	fmt.Println(str)
}
