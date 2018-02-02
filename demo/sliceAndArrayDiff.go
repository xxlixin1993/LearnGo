package main

import "fmt"

func main() {
	var array = [3]string{"a", "b", "c"}
	var slice = []string{"a", "b", "c"}

	swapSlice(slice)
	swapArray(array)

	fmt.Println(array)
	fmt.Println(slice)
}

func swapSlice(s []string) {
	s[1], s[2] = s[2], s[1]
}
func swapArray(s [3]string) {
	s[1], s[2] = s[2], s[1]
}
