package mock

import (
	"fmt"
)

type Retriver struct {
	Contents string
}

// toString
func (r *Retriver) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s}",
		r.Contents)
}

func (r *Retriver) Post(url string,
	form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriver) Get(url string) string {
	return r.Contents
}
