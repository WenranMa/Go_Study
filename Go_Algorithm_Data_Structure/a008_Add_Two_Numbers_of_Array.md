# 两数相加，数组表示

```go
package main

import "fmt"

func main() {
	num1 := []int{1, 3, 3, 5}
	num2 := []int{3, 5, 6}
	fmt.Println(add(num1, num2))
}

func add(num1, num2 []int) []int {
	res := []int{}
	l1 := len(num1)
	l2 := len(num2)
	var b int
	for i, j := l1-1, l2-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var n1, n2 int
		if i >= 0 {
			n1 = num1[i]
		}
		if j >= 0 {
			n2 = num2[j]
		}
		n := n1 + n2 + b
		res = append([]int{n % 10}, res...)
		b = n / 10
	}
	if b > 0 {
		res = append([]int{b}, res...)
	}
	return res
}
```