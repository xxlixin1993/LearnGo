package main

import "fmt"

func main() {
	c := make(map[string]int)

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1000000; j++ {
				c[fmt.Sprintf("%d", j)] = j
			}

		}()
	}
	// map是并发不安全的 多试几次会报错
	fmt.Println(c)
}
