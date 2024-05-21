package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && isValid(s[i]) {
			i++
		}
		for i < j && isValid(s[j]) {
			j--
		}
		if !equal(s[i], s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func isValid(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func equal(a, b byte) bool {
	la, lb := a, b
	if a >= 'A' && a <= 'Z' {
		la = a + 'a' - 'A'
	}
	if b >= 'A' && b <= 'Z' {
		lb = b + 'a' - 'A'
	}
    fmt.Println(a, b, la, lb)

	return la == lb
}

