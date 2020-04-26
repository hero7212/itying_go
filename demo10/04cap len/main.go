package main

import "fmt"

func main() {

	var sliceA = make([]int, 4, 8)
	fmt.Println(sliceA) // [0 0 0 0]
	fmt.Printf("%d--%d", len(sliceA), cap(sliceA))
}
