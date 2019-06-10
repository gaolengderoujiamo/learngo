package main

import (
	"io/ioutil"
	"log"
	"os"
)

const log_filename = "./demotest/baidu/logfiles/Info_First.log"

func init() {
	log.SetPrefix("[Error] ")
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func checkOpenFileError(err error) {
	if err != nil {
		log.Fatalln("open file error!")
	}
}

func main() {
	os.Mkdir("./demotest/baidu/os/demo1/test_go", 0777)
	os.MkdirAll("./demotest/baidu/os/demo1/test_go/t1/s1", 0777)
	os.RemoveAll("./demotest/baidu/os/demo1/test_go")
	file, err := os.Open(log_filename)
	checkOpenFileError(err)
	defer file.Close()

	readFile, e := ioutil.ReadFile(log_filename)
	checkOpenFileError(e)
	log.Println(string(readFile))
}
