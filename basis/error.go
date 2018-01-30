package main

import (
	"syscall"
	"fmt"
	"errors"
	"os"
)

func main(){
	var err error = syscall.Errno(2)

	fmt.Println(os.IsExist(err)) // false
	fmt.Println(os.IsNotExist(err)) // true

	fmt.Println(err.Error()) // "no such file or directory"
	fmt.Println(err)         // "no such file or directory"
}


func Errorf(format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(format, args...))
}