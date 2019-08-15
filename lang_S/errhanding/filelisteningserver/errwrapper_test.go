package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter,
	request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter,
	request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrNotExist
}

func errNotPermission(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrPermission
}

func errUnknow(writer http.ResponseWriter,
	request *http.Request) error {
	return fmt.Errorf("unknow error")
}

func noError(writer http.ResponseWriter,
	request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNotPermission, 403, "Forbidden"},
	{errUnknow, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

// 仅测试errWrapper, 使用假的request与response
func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		// 假的request与假的response
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"https://www.baidu.com", nil)
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

// 启动一个测试服务进行测试
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("expect (%d, %s); got (%d, %s)", expectedCode, expectedMsg, resp.StatusCode, body)
	}
}
