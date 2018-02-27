package main

import (
	"runtime"
	"fmt"
)


func main(){
	ms := &runtime.MemStats{}
	runtime.ReadMemStats(ms)
	fmt.Printf("%d kb",ms.Alloc / 1024)
}

