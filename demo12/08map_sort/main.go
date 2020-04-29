package main

import (
	"fmt"
	"sort"
)

func main() {

	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	var keySlice []int
	for key, _ := range map1 {
		keySlice = append(keySlice, key)
	}

	fmt.Println(keySlice)

	// 升序
	sort.Ints(keySlice)

	for _, v := range keySlice {
		fmt.Printf("key=%v value=%v\n", v, map1[v])
	}

}
