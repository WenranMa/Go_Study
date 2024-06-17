package main

import "fmt"

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// func main() {
// 	m1 := make(map[int]int, 0)
// 	m1[1] = 2

// 	i := 100
// 	for k, v := range m1 {
// 		m1[i] = i
// 		i += 1
// 		fmt.Println(k, v)
// 	}

// 	s := []int{1, 2, 3}
// 	ms := make(map[int]*int, 0)
// 	for i, n := range s {
// 		ms[i] = &n
// 	}
// 	for k, v := range ms {
// 		fmt.Println(k, *v)
// 	}

// 	var m = map[string]int{
// 		"A": 21,
// 		"B": 22,
// 		"C": 23,
// 	}
// 	counter := 0
// 	for k, v := range m {
// 		if counter == 0 {
// 			delete(m, "A")
// 		}
// 		counter++
// 		fmt.Println(k, v)
// 	}
// 	fmt.Println("counter is ", counter)
// }
