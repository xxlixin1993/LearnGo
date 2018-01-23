package main

import (
	"fmt"
	"strings"
)

func main() {
	testFor()
	testForRange()
	testJoin()
}

func testFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func testForRange() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for k, v := range i {
		fmt.Printf("key: %s value: %s", k, v)
		fmt.Println()
	}
}

// +=连接原字符串、空格和下个参数，产生新字符串在大数据量的时候效率低 建议使用Join（string包）函数
func testJoin() {
	i := []string{"a", "b", "c"}
	fmt.Println(strings.Join(i[0:], " "))
}
