package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

// 测试 go run readFile.go ../README.md
func main() {
	content := make(map[string]int)

	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "readFile: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			content[line]++
		}
	}

	fmt.Println(content)
}
