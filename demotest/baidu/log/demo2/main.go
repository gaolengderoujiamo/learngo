package main

import (
	"log"
	"os"
)

func main() {
	logfile, err := os.OpenFile("./demotest/baidu/logfiles/log_debug.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	defer logfile.Close()
	if err != nil {
		log.Fatalln("create logfile error!")
	}

	// Ldate         = 1 << iota     // 形如 2009/01/23 的日期
	// Ltime                         // 形如 01:23:23   的时间
	// Lmicroseconds                 // 形如 01:23:23.123123   的时间
	// Llongfile                     // 全路径文件名和行号: /a/b/c/d.go:23
	// Lshortfile                    // 文件名和行号: d.go:23
	// LstdFlags     = Ldate | Ltime // 日期和时间
	debuglog := log.New(logfile, "[Debug]", log.Llongfile)
	debuglog.Println("flag = [log.Llongfile], A debug message here")
	debuglog.SetFlags(debuglog.Flags() | log.LstdFlags)
	debuglog.Println("flag = [log.Llongfile | log.LstdFlags], A debug message here")
	debuglog.SetPrefix("[Info]")
	debuglog.Println("flag = [log.Llongfile | log.LstdFlags], A info message here")
	debuglog.SetFlags(log.Lshortfile | log.Llongfile | log.LstdFlags)
	debuglog.Println("flag = [log.Lshortfile | log.LstdFlags], A info message here")

}
