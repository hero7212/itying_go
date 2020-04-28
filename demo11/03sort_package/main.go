package main

import (
	"fmt"
	"sort"
)

func main() {
	var numSlice = []int{1, 10, 3, 6, 9, 21}
	sort.Ints(numSlice)
	fmt.Println(numSlice)
}
