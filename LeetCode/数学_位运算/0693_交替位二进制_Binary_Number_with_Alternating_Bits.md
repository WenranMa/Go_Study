# 693_交替位二进制_Binary_Number_with_Alternating_Bits
Given a positive integer, check whether it has alternating bits: namely, if two adjacent bits will always have different values.

Example:

    Input: 5  
    Output: True   
    Explanation: The binary representation of 5 is: 101

    Input: 7
    Output: False  
    Explanation: The binary representation of 7 is: 111.

    Input: 11  
    Output: False  
    Explanation: The binary representation of 11 is: 1011.

    Input: 10  
    Output: True  
    Explanation: The binary representation of 10 is: 1010.

### 解：

```go
func hasAlternatingBits(n int) bool {
    for n != 0 {
        a := (n & 1) == 0
        b := (n & 2) == 0
        if a == b {
            return false
        }
        n = n >> 1
    }
    return true
}
```