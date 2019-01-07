package main

import (
	"fmt"
)

type S struct {
	a int
	b string
}

func main() {
	//test2()
	test3()
}

func test2() {
	data := make([]int, 10, 20)
	data[0] = 1
	data[1] = 2
	dataappend := make([]int, 10, 20) //len <=10 则 	result[0] = 99 会 影响源Slice
	dataappend[0] = 1
	dataappend[1] = 2
	result := append(data, dataappend...)

	fmt.Println("length:", len(result), ":", result)

	result[0] = 99
	result[11] = 98
	fmt.Println("length:", len(data), ":", data)
	fmt.Println("length:", len(result), ":", result)
	fmt.Println("length:", len(dataappend), ":", dataappend)

}

func test3() {
	s := S{}
	s.a = 100
	s.b = "??"

	ss := make([]S, 10)
	ss = append(ss, s)
	fmt.Println("s :", s)
	fmt.Println("ss :", ss)

	s.a = 200

	fmt.Println("s :", s)
	fmt.Println("ss:", ss)

	fmt.Println("========")
	s1 := &S{}
	s1.a = 100
	s1.b = "??"

	ss1 := make([]*S, 10)
	ss1 = append(ss1, s1)
	fmt.Println("s1 :", *s1)
	fmt.Println("ss1:", *ss1[10])

	s1.a = 200

	fmt.Println("s1 :", *s1)
	fmt.Println("ss1:", *ss1[10])

}
