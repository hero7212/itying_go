package main

import "fmt"

func main() {

	sliceA := []int{1, 2, 3, 45}
	sliceB := sliceA
	sliceB[0] = 11
	fmt.Println(sliceA)
	fmt.Println(sliceB)
}
