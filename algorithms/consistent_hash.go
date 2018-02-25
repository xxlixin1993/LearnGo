package main

import (
	"strconv"
	"sort"
	"hash/crc32"
)

const KVirtualNodes = 40

// Hash function can be customized
type HashFunc func(data []byte) uint32

type Container struct {
	virtual  int
	nodes    []int
	hashMap  map[int]string
	hashFunc HashFunc
}


func NewContainer(virtualNodes int, fn HashFunc) *Container {
	c := &Container{
		virtual:  virtualNodes,
		hashFunc: fn,
		hashMap:  make(map[int]string),
	}
	if c.hashFunc == nil {
		c.hashFunc = crc32.ChecksumIEEE
	}
	return c
}

// Add nodes to the hashMap
func (c *Container) Add(nodes ...string) {
	for i := 0; i < c.virtual; i++ {
		// TODO weight
		for _, node := range nodes {
			hash := int(c.hashFunc([]byte(strconv.Itoa(i) + node)))
			c.nodes = append(c.nodes, hash)
			c.hashMap[hash] = node
		}
	}

	sort.Ints(c.nodes)
}

// Find the search cache key in the hashMap
func (c *Container) Get(key string) string {
	if c.IsEmpty() {
		return ""
	}

	hash := int(c.hashFunc([]byte(key)))

	idx := sort.SearchInts(c.nodes,hash)

	// Cycled back
	if idx == len(c.nodes) {
		idx = 0
	}

	return c.hashMap[c.nodes[idx]]
}

// Return the container whether empty or not
func (c *Container) IsEmpty() bool {
	return len(c.nodes) == 0
}

