package main

import "fmt"

func main() {

	var arr = [...]int{12, 23, 45}
	var sum = 0
	var l = len(arr)
	for i := 0; i < l; i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
	fmt.Println(float64(sum) / float64(l))
}
