package main

import (
	"fmt"
	"strconv"
)

type S struct {
	a int
	b string
}

type C float64
type F float64

func main() {
	//test2()
	//test3()

	//fmt.Println(test4(1, 3))

	//var a int = 11
	//var b int = 12
	//fmt.Println(a, b)

	//a, b := test5()

	//fmt.Println(a, b)

	// f := p()

	// fmt.Println(f)
	// fmt.Println(p() == p())

	// var a = 3
	// //fmt.Println(a.(string))

	// _, ok := interface{}(a).(int)
	// fmt.Println(ok)

	// var cc C = 12.0
	// var ff F = 123.0

	// fmt.Println(cc - ff)

	fmt.Println(convertToBase7(10))
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

func test4(x, y int) (z int) {

	if z := x + y; z > 5 {

		fmt.Println("?")
		return z
	}

	return
	// { // 不能在一个级别，引发 "z redeclared in this block" 错误。
	// 	var z = x + y
	// 	// return   // Error: z is shadowed during return
	// 	return z // 必须显式返回。
	// }
}

func test5() (a int, b int) {
	a = 5
	b = 3
	return 2, 4
}

func p() *int {
	v := 1
	return &v
}

func transpose(A [][]int) [][]int {
	r := len(A)
	if r == 0 {
		return A
	}
	c := len(A[0])
	if c == 0 {
		return A
	}
	ans := [][]int{}
	for i := 0; i < c; i++ {

		row := []int{}
		for j := 0; j < r; j++ {
			row = append(row, A[j][i])
		}
		ans = append(ans, row)
	}
	return ans
}
