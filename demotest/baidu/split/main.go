package main

import (
	"fmt"
	"strings"
)

func main() {

	p := strings.Split("MC221/220453/223989/704758/", "/")
	n := len(p)
	//for _, s := range p {
	//	fmt.Println(s)
	//}
	prepath := p[n-4] + "/" + p[n-3] + "/" + p[n-2] + "/"
	fmt.Println(prepath)
}
