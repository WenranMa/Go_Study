package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scores := make([]int, n)
	scanner.Scan()
	temp := strings.Split(scanner.Text(), " ")
	for i := 0; i < n; i++ {
		s, _ := strconv.Atoi(temp[i])
		scores = append(scores, s)
	}
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	fmt.Println(maxResult(scores, k))
}
func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	queue := make([]int, n) // 模拟双端队列
	qi, qj := 0, 1
	for i := 1; i < n; i++ {
		for qi < qj && queue[qi] < i-k {
			qi++
		}
		dp[i] = dp[queue[qi]] + nums[i]
		for qi < qj && dp[queue[qj-1]] <= dp[i] {
			qj--
		}
		queue[qj] = i
		qj++
	}
	return dp[n-1]
}
