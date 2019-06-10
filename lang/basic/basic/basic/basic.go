package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
1.定义变量并赋值 :=, 这样的初始化方式只能应用于函数内, 不能应用在函数外面
*/

func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func tryangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func enums() {
	const (
		cpp = iota
		java
		python
		javascript
		_
		sql
	)
	const (
		b int64 = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, javascript, sql)
	fmt.Print(b, kb, mb, gb, tb, pb)
}

func main() {
	euler()
	tryangle()
	enums()
}
