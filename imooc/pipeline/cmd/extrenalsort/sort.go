package main

import (
	"bufio"
	"fmt"
	"learngo/imooc/pipeline/pipeline"
	"os"
	"strconv"
)

func main() {
	p, files := createNetworkPipeline("large.in", 800000000, 4)
	defer func() {
		for _, file := range files {
			file.Close()
		}
	}()
	writeToFile(p, "large.out")
	printFile("large.out")

}
func printFile(filename string) {
	file, e := os.Open(filename)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	source := pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range source {
		fmt.Println(v)
		count++
		if count > 100 {
			break
		}
	}
}
func writeToFile(in <-chan int, filename string) {
	file, e := os.Create(filename)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	pipeline.WriterSink(writer, in)
}

func createPipeline(filename string, fileSize, chunkCount int) (<-chan int, []*os.File) {

	pipeline.Init()
	chunkSize := fileSize / chunkCount

	files := []*os.File{}
	sortResults := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, e := os.Open(filename)
		if e != nil {
			panic(e)
		}
		files = append(files, file)

		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, pipeline.InMemSort(source))
	}

	return pipeline.MergeN(sortResults...), files
}

func createNetworkPipeline(filename string, fileSize, chunkCount int) (<-chan int, []*os.File) {

	pipeline.Init()
	chunkSize := fileSize / chunkCount

	files := []*os.File{}
	sortAddr := []string{}
	for i := 0; i < chunkCount; i++ {
		file, e := os.Open(filename)
		if e != nil {
			panic(e)
		}
		files = append(files, file)

		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000+i)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}

	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}

	return pipeline.MergeN(sortResults...), files
}
