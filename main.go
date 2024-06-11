package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 164, 200 再看

func main() {
	nums := []int{1, 2, 4, 3, 5, 7, 6, 8, 4}
	fmt.Println(findDuplicate(nums)) // 输出应为 9，即 10-1 的差值
}

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]

		fmt.Println(slow, fast)
	}
	fast = 0
	for fast != slow {
		slow = nums[slow]
		fast = nums[fast]
		fmt.Println("xx", slow, fast)
	}
	return slow
}
