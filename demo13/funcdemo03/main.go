package main

import "fmt"

func sortInt(slice []int, s string) {
	l := len(slice)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			m := map[string]bool{
				"asc":  slice[i] > slice[j],
				"desc": slice[i] < slice[j],
			}
			if m[s] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

func main() {

	var slice = []int{1, 9, 2, 3, 566, 36, 6}
	sortInt(slice, "asc")
	fmt.Println(slice)
	sortInt(slice, "desc")
	fmt.Println(slice)
}
