# 762_二进制表示中的质数个数计算位置_Prime_Number_of_Set_Bits_in_Binary_Representation

Given two integers L and R, find the count of numbers in the range [L, R] (inclusive) having a prime number of set bits in their binary representation.

(Recall that the number of set bits an integer has is the number of 1s present when written in binary. For example, 21 written in binary is 10101 which has 3 set bits. Also, 1 is not a prime.)

Example:

    Input: L = 6, R = 10  
    Output: 4  
    Explanation:  
    6 -> 110 (2 set bits, 2 is prime)  
    7 -> 111 (3 set bits, 3 is prime)   
    9 -> 1001 (2 set bits , 2 is prime)   
    10->1010 (2 set bits , 2 is prime)   

    Input: L = 10, R = 15  
    Output: 5   
    Explanation:   
    10 -> 1010 (2 set bits, 2 is prime)   
    11 -> 1011 (3 set bits, 3 is prime)  
    12 -> 1100 (2 set bits, 2 is prime)  
    13 -> 1101 (3 set bits, 3 is prime)  
    14 -> 1110 (3 set bits, 3 is prime)  
    15 -> 1111 (4 set bits, 4 is not prime)  

Note:  
- L, R will be integers L <= R in the range [1, 10^6].  
- R - L will be at most 10000.

### 解：

结合191计算1的个数。
结合204判断质数。

```go

func countPrimeSetBits(L int, R int) int {
	ans := 0
	for i := L; i <= R; i++ {
		b := countOnes(i)
		if b > 1 && isPrime(b) { // checkPrime(b)
			ans += 1
		}
	}
	return ans
}

// 这是一个偷懒的写法
func checkPrime(n int) bool {
	if n == 2 || n == 3 || n == 5 || n == 7 || n == 11 || n == 13 || n == 17 || n == 19 {
		return true
	}
	return false
}

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func countOnes(n int) int {
	b := 0
	for n != 0 {
		b += n & 1
		n = n >> 1
	}
	return b
}
```