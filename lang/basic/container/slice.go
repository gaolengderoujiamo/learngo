package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func theBasicSliceTest(){
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6]",arr[2:6])
	fmt.Println("arr[2:6]",arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 =",s1)
	s2 := arr[:]
	fmt.Println("s2 =",s2)

	updateSlice(s1)
	fmt.Println("afterFunc:",s1)
	fmt.Println(arr)

	updateSlice(s2)
	fmt.Println("afterFunc:",s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println("I'm the source",s2)
	s2 = s2[:5]
	fmt.Println("second",s2)
	s2 = s2[3:]
	updateSlice(s2)
	fmt.Println("third",s2)
	fmt.Println(arr)
}

func theSecondSliceTest(){
	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[2:6]	//{2,3,4,5}
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n",
		s1,len(s1),cap(s1))
	s2 := s1[3:5]	//{5,6}
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n",
		s2,len(s2),cap(s2))
	s3 := s2[:3]
	fmt.Printf("s3=%v,len(s3)=%d,cap(s3)=%d\n",
		s3,len(s3),cap(s3))
}

func testSliceAppend(){
	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2,10)
	s4 := append(s3,11)
	s5 := append(s4,12)
	fmt.Println("s3 =",s3)
	fmt.Println("s4 =",s4)
	fmt.Println("s5 =",s5)
	// s4 and s5 no longer view arr.
	fmt.Println(arr)

}

func main() {
	//theBasicSliceTest()
	//theSecondSliceTest()
	testSliceAppend()
}
