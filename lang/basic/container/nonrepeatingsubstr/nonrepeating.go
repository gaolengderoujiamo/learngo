package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOcurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {

		if lastI, ok := lastOcurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOcurred[ch] = i
	}
	return maxLength
}

func main() {

	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(""))
	fmt.Println(
		lengthOfNonRepeatingSubStr("b"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcdefgh"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("我爱慕课网"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("三二一二三"))
}
