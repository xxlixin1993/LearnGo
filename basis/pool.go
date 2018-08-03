package main

import (
	"sync"
	"fmt"
	"runtime"
)

// sync.Pool
func main() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	case2(p)
}

// 简单使用
func case1(p *sync.Pool) {
	a := p.Get().(int)
	p.Put(1)
	p.Put(2)
	b := p.Get().(int)
	c := p.Get().(int)
	fmt.Println(a, b, c)
}

// sync.Pool缓存的期限只是两次gc之间这段时间 所有不可以使用sync.Pool去实现一个socket连接池的
func case2(p *sync.Pool) {
	a := p.Get().(int)
	p.Put(1)
	runtime.GC()
	b := p.Get().(int)
	fmt.Println(a, b)
}
