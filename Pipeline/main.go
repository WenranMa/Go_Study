package main

import (
	"./node"
	"bufio"
	"fmt"
	"os"
)

const (
	filename = "large.in"
	count    = 100000000
)

func main() {
	//mergeDemo()

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := node.RandomSource(count)

	w := bufio.NewWriter(file)
	node.WriterSink(w, p) //???
	w.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = node.ReaderSource(bufio.NewReader(file)) //???

	c := 0
	for v := range p {
		fmt.Println(v)
		c++
		if c > 100 {
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
	//如果运行上面的打印代码，则chaneel数据已经消费 下面的InMemSort传入的为空chaneel

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

func smallFileDemo() {

}
