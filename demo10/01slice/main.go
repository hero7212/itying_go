package main

import "fmt"

func main() {

	var arr1 []int
	fmt.Printf("%v - %T- 长度%v\n", arr1, arr1, len(arr1))

	var arr2 = []int{1, 2, 34, 45}
	fmt.Printf("%v - %T- 长度%v\n", arr2, arr2, len(arr2))

	var arr3 = []int{1: 2, 2: 4, 3: 6}
	fmt.Printf("%v - %T- 长度%v\n", arr3, arr3, len(arr3))

	fmt.Println(arr1 == nil) // 切片的默认值是nil
	fmt.Println(arr2 == nil) // false
}
