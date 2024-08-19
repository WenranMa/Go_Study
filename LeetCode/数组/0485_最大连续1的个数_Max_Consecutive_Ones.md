# 485_最大连续1的个数_Max_Consecutive_Ones
Given a binary array, find the maximum number of consecutive 1s in this array.

Example: 

    Input: [1,1,0,1,1,1]   
    Output: 3   
    Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.   

Note:   
The input array will only contain 0 and 1.   
The length of input array is a positive integer and will not exceed 10,000

### 解：

就用一共变量计数即可，如果遇到0则重置。退出循环时，再判断一次。

```go
func findMaxConsecutiveOnes(nums []int) int {
    ans := 0
    c := 0
    for _, e := range nums {
        if e == 1 {
            c += 1
        } else {
            if ans < c {
                ans = c
            }
            c = 0
        }
    }
    if ans < c {
        ans = c
    }
    return ans
}
```