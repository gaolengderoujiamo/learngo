package main

import "fmt"

func swap(a,b *int){
	*a,*b = *b,*a
}

func swapTwo(a,b int) (int,int){
	return b,a
}

func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
	a,b := 3,4
	//swap(&a,&b)
	a,b = swapTwo(a,b)
	fmt.Println(a,b)

}
