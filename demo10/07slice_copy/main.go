package main

import "fmt"

func main() {

	// sliceA := []int{1, 2, 3, 45}
	// sliceB := sliceA
	// sliceB[0] = 11
	// fmt.Println(sliceA)
	// fmt.Println(sliceB)

	sliceA := []int{1, 2, 3, 45}
	sliceB := make([]int, 4, 4)
	copy(sliceB, sliceA)
	sliceB[0] = 11
	fmt.Println(sliceA)
	fmt.Println(sliceB)
}
