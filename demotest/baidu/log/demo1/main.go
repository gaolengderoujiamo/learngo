package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("./demotest/baidu/logfiles/Info_First.log", os.O_APPEND|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		log.Fatalln("create file error!")
	}
	// Ldate         = 1 << iota     // 形如 2009/01/23 的日期
	// Ltime                         // 形如 01:23:23   的时间
	// Lmicroseconds                 // 形如 01:23:23.123123   的时间
	// Llongfile                     // 全路径文件名和行号: /a/b/c/d.go:23
	// Lshortfile                    // 文件名和行号: d.go:23
	// LstdFlags     = Ldate | Ltime // 日期和时间
	debugLog := log.New(file, "[Info]", log.Lshortfile)
	debugLog.Println("A Info message here")
	debugLog.SetPrefix("[Debug]")
	debugLog.Println("A Debug message here")
}
