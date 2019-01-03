package main

import (
	"./node"
	"bufio"
	"fmt"
	"os"
)

const (
	fileName    = "large.in"
	fileSize    = 100000000 //512
	chunkCount  = 4
	fileNameOut = "test.out"
)

func main() {
	//mergeDemo()

	//fileDemo()

	//external sorting.

	node.Init()

	p := createPipeline(fileName, fileSize, chunkCount)
	writeToFile(p, fileNameOut)
	printFile(fileNameOut)

}

func createPipeline(filename string, fileSize, chunkCount int) <-chan int {

	chunkSize := fileSize / chunkCount
	sortResults := []<-chan int{}

	for i := 0; i < chunkCount; i++ {

		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0) //???

		source := node.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, node.InMemSort(source))

		fmt.Println()
	}

	return node.MergeN(sortResults...)
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	node.WriterSink(writer, p)
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := node.ReaderSource(file, -1)

	c := 0
	for v := range p {
		fmt.Println(v)
		c++
		if c >= 50 {
			break
		}
	}
}

func mergeDemo() {
	p := node.ArraySource(2, 3, 4, 2, 2, 2, 5, 5, 234, 1, 4, 6, 32)

	/*
		for v := range p {
			fmt.Printf("source: %v\n", v)
		}
	*/
	//如果运行上面的打印代码，则channel数据已经消费 下面的InMemSort传入的为空chaneel

	s := node.InMemSort(p)
	for v := range s {
		fmt.Println("sorted: ", v)
	}

	p1 := node.ArraySource(2, 3, 4, 2, 2, 2, 5, 5, 234, 1, 4, 6, 32)
	p2 := node.ArraySource(2, 3, 23, 34, 12312, 23, 663, 773, 354, 213, 11, 2)
	out := node.Merge(node.InMemSort(p1), node.InMemSort(p2))
	for v := range out {
		fmt.Println("... ", v)
	}
}

func fileDemo() {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := node.RandomSource(fileSize)

	w := bufio.NewWriter(file)
	node.WriterSink(w, p) //???
	w.Flush()

	file, err = os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = node.ReaderSource(bufio.NewReader(file), -1) //???

	c := 0
	for v := range p {
		fmt.Println(v)
		c++
		if c > 100 {
			break
		}
	}
}
