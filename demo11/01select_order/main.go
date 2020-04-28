package main

import "fmt"

func main() {
	var numSlice = []int{1, 10, 3, 6, 9, 21}
	l := len(numSlice)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if numSlice[i] > numSlice[j] {
				// temp := numSlice[i]
				// numSlice[i] = numSlice[j]
				// numSlice[j] = temp
				numSlice[i], numSlice[j] = numSlice[j], numSlice[i]
			}
		}
	}
	fmt.Println(numSlice)
}
