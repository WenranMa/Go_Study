# 数据结构

## Array (Slice)
```go
// 插入删除元素
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s = insert(s, 5, 100)
	fmt.Println(s)

	s = remove(s, 5)
	fmt.Println(s)
}
func insert(s []int, index int, n int) []int {
	s = append(s, 0)
	copy(s[index+1:], s[index:])
	s[index] = n
	return s
}

func remove(nums []int, index int) []int {
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
	return nums[:len(nums)-1]
}
```
