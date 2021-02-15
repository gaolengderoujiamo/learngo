package main

import (
	"fmt"
	strings "strings"
)

func main() {

	filename := `G:\医院数据\陕西省肿瘤医院\pacsimge\MC187\186999`
	p := strings.Split(filename, "\\")
	for _, v := range p {
		fmt.Println(v)
	}
}
