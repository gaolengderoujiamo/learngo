package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	for ; n >= 0; n /= 2 {
		if n == 0 {
			return "0"
		}
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// 读文件
func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	//fmt.Print(convertToBin(0))
	printFile("lang/basic/basic/loop/abc.txt")
	s := `asd
		pwd
		"sdsd"

		123`
	printFileContents(strings.NewReader(s))

}
