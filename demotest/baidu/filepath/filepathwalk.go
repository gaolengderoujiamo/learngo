package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	// 如下语句可以过滤路径
	//if path == "E:\\个人留存\\_______装机与软件留存\\娱乐软件\\Office2007" {
	//	return filepath.SkipDir
	//}
	fmt.Printf("%s\n", path)
	return nil
}

func main() {
	//遍历打印所有的文件名
	filepath.Walk("E:/个人留存/_______装机与软件留存/娱乐软件", walkFunc)
}
