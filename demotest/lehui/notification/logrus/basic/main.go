package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	//demo1()
	//demo2()
	//demo3()
	//demo4()
	demo5("100", "127.0.0.1")
}

func demo5(request_id, user_ip string) {
	requestLogger := logrus.WithFields(logrus.Fields{"request_id": request_id, "user_ip": user_ip})
	requestLogger.Info("something happened on that request") // will log request_id and user_ip
	requestLogger.Warn("something not great happened")
}

func demo4() {
	// logrus提供了New()函数来创建一个logrus的实例.
	// 项目中,可以创建任意数量的logrus实例.
	var log = logrus.New()
	// 为当前logrus实例设置消息的输出,同样地,
	// 可以设置logrus实例的输出到任意io.writer
	log.Out = os.Stdout

	// 为当前logrus实例设置消息输出格式为json格式.
	// 同样地,也可以单独为某个logrus实例设置日志级别和hook,这里不详细叙述.
	log.Formatter = &logrus.JSONFormatter{}

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}

func demo3() {
	// 设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	logrus.SetLevel(logrus.WarnLevel)

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}
func demo2() {
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Warn("A walrus appears")
}

func demo1() {
	logrus.Debug("Useful debugging information.")
	logrus.Info("Something noteworthy happened!")
	logrus.Warn("You should probably take a look at this.")
	logrus.Error("Something failed but I'm not quitting.")
	//logrus.Fatal("Bye.")         //log之后会调用os.Exit(1)
	logrus.Panic("I'm bailing.")
	//log之后会panic()
}
