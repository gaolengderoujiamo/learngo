package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!"
	// 使用len获得字节的长度
	fmt.Println(len(s))

	// utf-8
	// 英文每个1字节，中文每个3字节
	// 使用[]byte获得字节
	for _,v := range []byte(s) {
		fmt.Printf("%X ",v)
	}
	fmt.Println()
	for i, ch := range s {	// ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	// 使用utf8.RuneCountInString获得字符数量
	fmt.Println("Rune count:",utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}



}
