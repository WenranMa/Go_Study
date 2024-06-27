package main

import "fmt"

func main() {
	isMatch := func(i int) bool {
		switch i {
		case 1:
			fallthrough
		case 2:
			return true
		}
		return false
	}
	fmt.Println(isMatch(1)) // true
	fmt.Println(isMatch(2)) // true
	match := func(i int) bool {
		switch i {
		case 1, 2:
			return true
		}
		return false
	}
	fmt.Println(match(1)) // true
	fmt.Println(match(2)) // true
}
