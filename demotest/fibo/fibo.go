package main

import "fmt"

func main() {
	f := fibo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func fibo() func() int {
	x, y := 0, 1
	return func() int {

		x, y = y, x+y
		return x
	}
}
