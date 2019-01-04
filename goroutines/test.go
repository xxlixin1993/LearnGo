package main

import (
	"fmt"
)
func main() {
	arrayA := [2]int{100, 200}

	testArrayPoint(&arrayA)   // 1.传数组指针

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
}

func testArrayPoint(x *[2]int) {
	fmt.Printf("func Array : %p , %v\n", x, *x)
	(*x)[1] += 100
}
