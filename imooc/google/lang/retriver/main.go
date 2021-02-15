package main

import (
	"fmt"
	"learngo/imooc/google/lang/retriver/mock"
	real3 "learngo/imooc/google/lang/retriver/real"
	"time"
)

// 接口内方法的定义不需要func关键字, 不需要接受者
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

// 接口的组合
// 接口内可以组合接口, 同时还能添加自己的方法
type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.imooc.com"

/**
 ** 使用者定义接口!
 ** 使用者定义接口!
 ** 使用者定义接口!
 */

// 接口在此处被使用!!
// 根据 [由使用者定义接口]的原理↓
// 接口在本包下定义

// 实现者在
func downLoad(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "goland",
		})
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another facked imooc.com",
	})
	return s.Get(url)
}

/*
	接口变量有两个内容,
	一是实现者的类型, 二是实现者的值(也可以是指针)
*/

func main() {
	var r Retriever
	retrirver := mock.Retriver{
		"this is a fake imooc.com"}
	r = &retrirver
	inspect(r)
	r = &real3.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// ② Type assertion 的方式查看接口变量
	if realRetriver, ok := r.(*real3.Retriever); ok {
		fmt.Println(realRetriver.TimeOut)
	} else {
		fmt.Println("not a mock retriver")
	}

	fmt.Println("Try a session...")
	fmt.Println(session(&retrirver))
	//fmt.Println(downLoad(r))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type swich:")
	// ① Type Switch的方式查看接口的变量
	// r.(type) 可以得到r的类型
	switch v := r.(type) {
	case *mock.Retriver:
		fmt.Println("Contents:", v.Contents)
	case *real3.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
