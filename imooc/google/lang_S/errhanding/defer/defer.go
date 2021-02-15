package main

import (
	"bufio"
	"fmt"
	"learngo/imooc/google/lang_S/errhanding/fib"
	"log"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func writeFile(filename string) {
	file, e := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if e != nil {
		if pathError, ok := e.(*os.PathError); !ok {
			panic(e)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func getPwd() string {
	pwd, err := os.Getwd() // 当前工作目录
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	return pwd
}

func main() {
	//pwd := getPwd()
	writeFile("fib.txt")
}
