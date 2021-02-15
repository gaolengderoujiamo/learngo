package main

import (
	"bufio"
	"errors"
	"fmt"
	"learngo/imooc/google/lang/functional/fib/fib"
	"os"
)

func writeFile(fileName string) {
	file, err := os.OpenFile(
		fileName, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		} else {
			panic(err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
}
