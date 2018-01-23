package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	echoInput()
}

// 测试 echo 'test echoInput' | ./input
func echoInput() {
	// 从标准输入中获取
	input := bufio.NewScanner(os.Stdin)
	// 一行一行读
	for input.Scan() {
		// 输出内容
		fmt.Println("123" + input.Text())
	}
}

