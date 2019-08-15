package main

import (
	"learngo/lang/errHandling/filelisteningserver/filelistening"
	"net/http"
	"os"

	"github.com/gpmgo/gopm/modules/log"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		// panic
		defer func() {
			if r := recover(); r != nil {
				log.Warn("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Warn("Error handling request: %s",
				err.Error())

			// userError
			// Type Assertion
			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

// 可以给用户看到的error
type userError interface {
	// 供系统
	error
	// 暴露给用户
	Message() string
}

func main() {
	http.HandleFunc("/",
		errWrapper(filelistening.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
