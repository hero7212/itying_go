package main

import (
	"fmt"
	"sort"
)

func mapSort(map1 map[string]string) string {

	var sliceKey []string

	for k, _ := range map1 {
		sliceKey = append(sliceKey, k)
	}

	sort.Strings(sliceKey)

	var str string
	for _, v := range sliceKey {
		str += fmt.Sprintf("%v=>%v", v, map1[v])
	}
	return str
}

func main() {
	m1 := map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
	}
	str := mapSort(m1)
	fmt.Println(str)
}
