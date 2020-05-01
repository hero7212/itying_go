package main

import (
	"fmt"
	"time"
)

func main() {
	timeObj := time.Now()

	fmt.Println(timeObj)
	timeObj = timeObj.Add(time.Hour)
	fmt.Println(timeObj)
}
