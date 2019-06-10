package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating slice...")
	var s []int // Zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{1, 23, 4, 5, 6}
	fmt.Println(s1)

	s2 := make([]int, 16)
	//fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d",
	//	s2,len(s2),cap(s2))

	s3 := make([]int, 10, 32)

	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice...")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	// 从前面删除元素会使cap-1, 尾部删除则不会
	for i := 3; i > 0; i-- {
		fmt.Println("Popping from front")
		front := s2[0]
		s2 = s2[1:]
		fmt.Println(front)
		printSlice(s2)
	}

	for i := 3; i > 0; i-- {
		fmt.Println("Popping from back")
		back := s2[len(s2)-1]
		s2 = s2[:len(s2)-1]
		fmt.Println(back)
		printSlice(s2)
	}
}
