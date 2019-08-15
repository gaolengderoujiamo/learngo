package main

import "fmt"

type father int

func (this *father) prepare() {
	fmt.Println("father")
}

type son struct {
	father
}

//func (this *son) prepare() {
//	fmt.Println("son")
//}

type grandson struct {
	son
}

//func (this *grandson) prepare() {
//	fmt.Println("grandson")
//}

func main() {
	g := grandson{}
	g.prepare()
}

// prepare()仅执行一次, 首先查找当前类, 然后是父类, ...
