package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	//
	//str := "610221195107040037"
	str := "610421195308143818"

	//方法一
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

	fmt.Println(md5str1)
}
