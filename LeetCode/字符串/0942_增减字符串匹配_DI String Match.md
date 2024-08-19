# 942_增减字符串匹配_DI String Match
Given a string S that only contains "I" (increase) or "D" (decrease), let N = S.length.

Return any permutation A of [0, 1, ..., N] such that for all i = 0, ..., N-1:

If S[i] == "I", then A[i] < A[i+1]  
If S[i] == "D", then A[i] > A[i+1]
 
Example:  
Input: "IDID"   
Output: [0,4,1,3,2]   
Input: "III"   
Output: [0,1,2,3]   
Input: "DDI"   
Output: [3,2,0,1]

Note:   
1 <= S.length <= 10000   
S only contains characters "I" or "D".   


### 解：

遇到I就就从最小的来，然后+1，遇到D就从最大的来，然后-1。

```go
func diStringMatch(S string) []int {
    j := len(S)
    i := 0
    ans := []int{}
    for _, c := range S {
        if c == 'I' {
            ans = append(ans, i)
            i++
        } else if c == 'D' {
            ans = append(ans, j)
            j--
        }
    }
    ans = append(ans, i)
    return ans
}
```