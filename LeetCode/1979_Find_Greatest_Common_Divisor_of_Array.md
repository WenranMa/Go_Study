# 1979. Find Greatest Common Divisor of Array

### Easy

Given an integer array nums, return the greatest common divisor of the smallest number and largest number in nums.

The greatest common divisor of two numbers is the largest positive integer that evenly divides both numbers.

### Example 1:

Input: nums = [2,5,6,9,10]
Output: 2
Explanation:
The smallest number in nums is 2.
The largest number in nums is 10.
The greatest common divisor of 2 and 10 is 2.

### Example 2:

Input: nums = [7,5,6,8,3]
Output: 1
Explanation:
The smallest number in nums is 3.
The largest number in nums is 8.
The greatest common divisor of 3 and 8 is 1.

### Example 3:

Input: nums = [3,3]
Output: 3
Explanation:
The smallest number in nums is 3.
The largest number in nums is 3.
The greatest common divisor of 3 and 3 is 3.

Constraints:

2 <= nums.length <= 1000
1 <= nums[i] <= 1000

```go
/*
欧几里得辗转相除法：
gcd(x,y)表示x和y的最大公约数
进入运算时:x!=0,y!=0，本质上就是不断转换成求等价更小数的最大公约数。如果x%y=0，返回y，即最大公约数。
gcd(x,y)=gcd(y,x%y)

证明：
设k=x/y，b=x%y  则：x=ky+b
如果n能够同时整除x和y，则(y%n)=0,(ky+b)%n=0，则b%n=0，即n也同时能够整除y和b。
由上得出：同时能够整除y和(b=x%y)的数，也必然能够同时整除x和y。故而gcd(x,y)=gcd(y,x%y)。当(b=x%y)=0，即y可以整除x，这时的y也就是所求的最大公约数了。
*/

func findGCD(nums []int) int {
	min := 1001
	max := 0
	for _, n := range nums {
		if min > n {
			min = n
		}
		if max < n {
			max = n
		}
	}
	return gcd(max, min)
}

func gcd(x, y int) int {
	temp := 0
	for {
		temp = x % y
		if temp > 0 {
			x = y
			y = temp
		} else {
			return y
		}
	}
}
```