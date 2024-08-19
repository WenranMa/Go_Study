# 344_反转字符串_Reverse String
Write a function that reverses a string. The input string is given as an array of characters char[]. Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory. You may assume all the characters consist of printable ascii characters.

Example:  

    Input: ["h","e","l","l","o"]  
    Output: ["o","l","l","e","h"]  
    Input: ["H","a","n","n","a","h"]  
    Output: ["h","a","n","n","a","H"]

```go
func reverseString(s []byte) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}
```