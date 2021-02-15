package main

import "fmt"

/*
	数组是值传递
*/

func printArray(arr []int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}

}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{1, 3, 5, 7, 10}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	//for _,v := range arr3{
	//	sum += v
	//	fmt.Println(v)
	//}
	//fmt.Println("和为：",sum)
	printArray(arr3[:])
	fmt.Println(arr3)
}
