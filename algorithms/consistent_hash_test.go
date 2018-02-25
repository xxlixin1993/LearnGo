package main

import (
	"testing"
)


func TestAddAndGet(t *testing.T) {
	var container1 = NewContainer(40, nil)
	var container2 = NewContainer(40, nil)

	container1.Add("127.0.0.1", "192.168.0.1", "172.16.0.1", "10.0.0.1")
	container2.Add("127.0.0.1", "192.168.0.1", "172.16.0.1", "10.0.0.1")

	if container1.Get("foo") != container2.Get("foo") {
		t.Errorf("Get keys error, expect to return the same")
	}
}
