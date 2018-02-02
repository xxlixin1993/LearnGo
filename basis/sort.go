package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func main() {
	testS := StringSlice{}
	testS = append(testS, "1")
	testS = append(testS, "3")
	testS = append(testS, "2")
	fmt.Println(testS)
	// asc
	sort.Sort(testS)
	fmt.Println(testS)
	// desc
	sort.Sort(sort.Reverse(testS))
	fmt.Println(testS)
}

func (ss StringSlice) Len() int {
	return len(ss)
}

func (ss StringSlice) Less(i int, j int) bool {
	return ss[i] < ss[j]
}

func (ss StringSlice) Swap(i int, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
