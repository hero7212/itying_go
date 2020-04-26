package main

import "fmt"

func main() {
	// var sliceA []int
	// fmt.Printf("%v -- %v--%v \n", sliceA, len(sliceA), cap(sliceA))
	// sliceA = append(sliceA, 12)
	// sliceA = append(sliceA, 24)
	// fmt.Printf("%v -- %v--%v \n", sliceA, len(sliceA), cap(sliceA))

	// var sliceA []int
	// sliceA = append(sliceA, 12, 23, 35, 465)
	// fmt.Printf("%v -- %v--%v \n", sliceA, len(sliceA), cap(sliceA))

	// append可以合并切片

	sliceA := []string{"php", "java"}
	sliceB := []string{"nodejs", "go"}

	sliceA = append(sliceA, sliceB...)
	fmt.Println(sliceA)
}
