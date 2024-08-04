package main

import (
	"fmt"
	"io"
)

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func match(odd int, evens []int, visited map[int]int, suited map[int]int) bool {
	for _, even := range evens {
		if isPrime(odd+even) && visited[even] == 0 {
			visited[even] = 1
			if suited[even] == 0 || match(suited[even], evens, visited, suited) {
				suited[even] = odd
				return true
			}
		}
	}
	return false
}

func main() {
	for {
		var n int
		var num int
		var odds []int
		var evens []int

		c, err := fmt.Scanf("%d\n", &n)
		if c == 0 || err == io.EOF {
			break
		}

		//分奇数和偶数
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &num)
			if num%2 == 0 {
				evens = append(evens, num)
			} else {
				odds = append(odds, num)
			}
		}

		var suited map[int]int = make(map[int]int)
		var res int
		for i := 0; i < len(odds); i++ {

			var visited map[int]int = make(map[int]int)

			ok := match(odds[i], evens, visited, suited)
			if ok {
				res++
			}
		}

		fmt.Println(res)
	}
}
