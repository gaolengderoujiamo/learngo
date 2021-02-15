package main

import "log"

type People interface {
	Sex()
}

type Man struct{}

func (m *Man) Sex() {
	log.Print("Man")
}

type Woman struct{}

func (m *Woman) Sex() {
	log.Print("Woman")
}

func main() {
	var t1 People

	t1 = new(Man)
	t1.Sex()

	t1 = new(Woman)
	t1.Sex()
}
