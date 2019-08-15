package main

import "learngo/demotest/demo/interview/sleepsort"

func main() {
	data := []int{32, 12, 0, 34, 2222, 1, 3, 15}
	sleepsort.SleepSort(data, 8)
}
