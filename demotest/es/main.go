package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	rbody := `{
			  "query": {
				"bool": {
				  "filter": {
					"term": {
					  "duns.keyword": "61011343520350111"
					}
				  }
				}
			  }
			}`

	resp, err := http.Post("http://10.63.72.12:9200/pacs/pacs/_count",
		"application/x-www-form-urlencoded",
		strings.NewReader(rbody))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}
