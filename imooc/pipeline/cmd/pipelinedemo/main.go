package main

import (
	"bufio"
	"fmt"
	"learngo/imooc/pipeline/pipeline"
	"os"
)

func main() {
	const filename = "small.in"
	const n = 64
	file, e := os.Create(filename)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	source := pipeline.RandomSource(n)

	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, source)
	defer writer.Flush()

	file, e = os.Open(filename)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	readerSource := pipeline.ReaderSource(reader, -1)

	count := 0
	for v := range readerSource {
		fmt.Println(v)
		count++
		//if count > 100 {
		//	break
		//}
	}
}
func mergeDemo() {
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(3, 7, 2, 8, 13, 9)),
		pipeline.InMemSort(pipeline.ArraySource(0, 9, 12, 5, 2, 7, 3)))
	for v := range p {
		fmt.Println(v)
	}
}
