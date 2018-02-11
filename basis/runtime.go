package main

import (
	"runtime"
	"fmt"
)


func main() {
	test()
}

func test() {
	test2()
}

func test2(){
	pc,file,line,ok := runtime.Caller(2)
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)
	f := runtime.FuncForPC(pc)
	fmt.Println(f.Name())
	fmt.Println("----------------")

	pc,file,line,ok = runtime.Caller(0)
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)
	f = runtime.FuncForPC(pc)
	fmt.Println(f.Name())
	fmt.Println("----------------")

	pc,file,line,ok = runtime.Caller(1)
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)
	f = runtime.FuncForPC(pc)
	fmt.Println(f.Name())
}
