package main

import "fmt"

func main() {
	var a = 20  /* 声明实际变量 */
	var ip *int /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	m := make(map[string]int)
	// map查找
	v, ok := m["s"]

	fmt.Println(v, ok)

	// map 赋值
	medals := []string{"gold", "silver", "bronze"}
	fmt.Println(medals)

	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}


	str := "汉字"
	fmt.Printf("%s", str[0:3])
	//fmt.Println("汉字")

}
