package main

import (
	"fmt"
	"log"
)

func test_deferpanic() {
	defer func() {
		fmt.Println("--first--")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	log.Panicln("test for defer Panic")
	defer func() {
		fmt.Println("--second--")
	}()
}

func main() {
	test_deferpanic()
}
