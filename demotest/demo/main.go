package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var cwd string

func main() {
	//t1()
	//t2()
	//t3()
	//t4()
	//t5()
	//selectTest()
	deferTest()
}

func deferFunc(name string, a, b int) int {
	fmt.Println(name, a, b)
	return a + b
}

func deferTest() {
	a, b := 1, 2
	defer deferFunc("one", a, deferFunc("two", a, b))
	a = 0
	defer deferFunc("three", a, deferFunc("four", a, b))
	panic("five")
}

func selectTest() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case r := <-ch:
			fmt.Println(r)
		case ch <- i:
		}
	}
}

func t5() {
	f := func() func() {
		fmt.Println("0123456789")
		return func() {
			fmt.Println("asdasdasdasd")
		}
	}
	defer f()()
	defer func() {
		fmt.Println("qwert")
	}()
	fmt.Println("0000000000")
}

func t4() {
	fmt.Printf("%.3s \n%-5d%s", "qwertyuio11111p", 123, "sd")
}

func t3() {
	var nn map[string]int
	n := make(map[string]int)
	m := map[string]int{}
	fmt.Println(nn == nil)
	fmt.Println(n == nil)
	fmt.Println(m == nil)
}

func t2() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	fmt.Print(<-c)
	c <- 4
}

func t1() {
	str1 := `<a href="http://www.zhenai.com/zhenghun/aba/nv">阿坝女士征婚</a>`
	str2 := `<span class="nickName" data-v-3c42fade>劣酒灼心</span>    </div> <div class="des f-cl" data-v-3c42fade>阿坝 | 25岁 | 中专 | 未婚 | 165cm | 5001-8000元</div> <div class="actions" data-v-3c42fade><div class="item sayHi" data-v-3c42fade>打招呼</div>`
	var genderRe = regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[\w]+/([\w]+)">[^<]*士征婚</a>`)
	var profileRe = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^\|]+)\|([^\|]+)\|([^\|]+)\|([^\|]+)\|([^\|]+)\|([^<]+)</div>`)
	submatch1 := genderRe.FindStringSubmatch(str1)
	for _, s := range submatch1 {
		fmt.Println(s)
	}

	submatch2 := profileRe.FindStringSubmatch(str2)
	for _, s := range submatch2 {
		fmt.Println(s)
	}

	str3 := " qwer  "
	fmt.Println(strings.Trim(str3, " "))

	str4 := "165力魔"
	str4rune := []rune(str4)
	fmt.Println(len(str4rune))
	fmt.Println(string(str4rune[:len(str4rune)-2]))

	ageRune := []rune("25岁")
	age, err := strconv.Atoi(string(ageRune[:len(ageRune)-1]))
	if err != nil {
		panic(err)
	}
	fmt.Println(age)

}
