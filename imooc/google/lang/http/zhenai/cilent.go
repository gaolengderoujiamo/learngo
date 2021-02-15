package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	url := "http://album.zhenai.com/u/1069129647"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	//request.Header.Add("Host", "album.zhenai.com")
	//cookie1 := &http.Cookie{Name: "sid", Value: "babb4716-d982-4272-9ae3-19c929b0494b", HttpOnly: true}
	//request.AddCookie(cookie1)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, e := httputil.DumpResponse(resp, true)
	if e != nil {
		panic(e)
	}

	fmt.Printf("%s\n", bytes)
}
