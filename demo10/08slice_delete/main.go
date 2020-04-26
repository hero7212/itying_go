package main

import "fmt"

func main() {

	a := []int{30, 31, 32, 33, 34, 35, 36, 37}

	a = append(a[:2], a[3:]...) // 删除的元素是32
	fmt.Println(a)

	sliceB := []int{30, 31, 32, 33, 34, 35, 36, 37}

	sliceB = append(sliceB[:5], sliceB[7:]...)
	fmt.Println(sliceB)
}
