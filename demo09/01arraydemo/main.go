package main

import "fmt"

func main() {

	var arr1 [3]int
	var arr2 [4]int
	var strArr [3]string

	fmt.Printf("arr1:%T	arr2:%T strArr:%T \n", arr1, arr2, strArr)
	fmt.Printf("arr1:%v	arr2:%v strArr:%v \n", arr1, arr2, strArr)

	fmt.Println(len(arr1))

	var arr3 = [3]int{1, 2, 3}
	fmt.Println(arr3)

	var arr4 = [...]int{1, 2}
	fmt.Println(arr4)

	arr5 := [...]int{0: 1, 1: 10, 3: 30}
	fmt.Println(arr5)
}
