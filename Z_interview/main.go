package main

import (
	"fmt"
)
func main() {
	v:= []int{-5,-3,-5}
	fmt.Println(maxSubArray(v))
}

func maxSubarraySumCircular(nums []int) int {
    n := len(nums)
    preMax, maxRes := nums[0], nums[0]
    preMin, minRes := nums[0], nums[0]
    sum := nums[0]
    for i := 1; i < n; i++ {
        preMax = max(preMax + nums[i], nums[i])
        maxRes = max(maxRes, preMax)
        preMin = min(preMin + nums[i], nums[i])
        minRes = min(minRes, preMin)
        sum += nums[i]

		fmt.Println("preMax:",preMax,"maxRes:",maxRes,"preMin:",preMin,"minRes:",minRes,"sum:",sum)
    }
    if maxRes < 0 {
        return maxRes
    } else {
        return max(maxRes, sum - minRes)
    }
}

func maxSubArray(nums []int) int {
	mx := -10001
	sum := 0
	for _, n := range nums {
		sum += n
		mx = max(mx, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return mx
}