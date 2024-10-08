package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./node"
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

	//sortDemo()

	p := createNetworkPipeline("small.in", 512, 4)
	writeToFile(p, "test_net.out")
	printFile("test_net.out")
}

func createNetworkPipeline(filename string, fileSize, chunkCount int) <-chan int {

	node.Init()
	chunkSize := fileSize / chunkCount
	//sortResults := []<-chan int{}

	sortAddrs := []string{}
	for i := 0; i < chunkCount; i++ {

		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0) //???

		source := node.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000+i)
		node.NetworkSink(addr, node.InMemSort(source))
		sortAddrs = append(sortAddrs, addr)

	}

	sortResults := []<-chan int{}
	for _, addr := range sortAddrs {
		sortResults = append(sortResults, node.NetworkSource(addr))
	}

	return node.MergeN(sortResults...)
}

func createPipeline(filename string, fileSize, chunkCount int) <-chan int {

	node.Init()
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

func sortDemo() {
	p := createPipeline(fileName, fileSize, chunkCount)
	writeToFile(p, fileNameOut)
	printFile(fileNameOut)
}

/*
### Pipeline: 外部排序Pipeline
选自慕课网：搭建并行处理管道

- 原始数据过大，无法一次读入内存，所以分块读入内存。每个块数据进行内部排序（直接调用API排序），
最后讲各个节点归并，归并选择两两归并。
*/
