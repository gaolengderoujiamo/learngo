package main

import (
	"fmt"
	"learngo/lang_S/container/queue"
)

func main() {
	q := queue.Queue{1, 2, 3}
	q.Push("asd")
	for !q.IsEmpty() {
		fmt.Println(q.Pop())
	}
}
