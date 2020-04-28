package main

import "fmt"

func main() {
	var numSlice = []int{1, 10, 3, 6, 9, 21}
	l := len(numSlice)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1; j++ {
			if numSlice[j] > numSlice[j+1] {

				numSlice[j], numSlice[j+1] = numSlice[j+1], numSlice[j]
			}
		}
	}
	fmt.Println(numSlice)
}
