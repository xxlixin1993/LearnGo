package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	s := "api_key"  + "param"  + "time"  + "version"
	signByte := []byte(s)
	hash := md5.New()
	hash.Write(signByte)
	fmt.Println(hex.EncodeToString(hash.Sum(nil)))
}
