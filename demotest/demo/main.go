package main

import "fmt"

var cwd string

func main() {
	fmt.Printf("%*s</%s>", 7, "s--", "ppty")
	var s interface{}
	s = 2
	i := s.(int) + 1
	fmt.Println(i)
}
