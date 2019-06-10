package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

// 该方法接收者为上面的struct, 所以是上面的结构体实现了interface
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}
