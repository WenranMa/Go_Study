# 944_删列造序_Delete Columns to Make Sorted
We are given an array A of N lowercase letter strings, all of the same length. Now, we may choose any set of deletion indices, and for each string, we delete all the characters in those indices. For example, if we have an array A = ["abcdef","uvwxyz"] and deletion indices {0, 2, 3}, then the final array after deletions is ["bef", "vyz"], and the remaining columns of A are ["b","v"], ["e","y"], and ["f","z"].  (Formally, the c-th column is [A[0][c], A[1][c], ..., A[A.length-1][c]].). Suppose we chose a set of deletion indices D such that after deletions, each remaining column in A is in non-decreasing sorted order. Return the minimum possible value of D.length.

Example:  
Input: ["cba","daf","ghi"]  
Output: 1  
Explanation:   
After choosing D = {1}, each column ["c","d","g"] and ["a","f","i"] are in non-decreasing sorted order.  
If we chose D = {}, then a column ["b","a","h"] would not be in non-decreasing sorted order.

Input: ["a","b"]
Output: 0  
Explanation: D = {}  

Input: ["zyx","wvu","tsr"]  
Output: 3  
Explanation: D = {0, 1, 2}
 
Note:  
1 <= A.length <= 100  
1 <= A[i].length <= 1000

### 解：

分别比较每一列就可以了，如果有降序，则该列需要删除，ans+=1.

```go
func minDeletionSize(A []string) int {
    l := len(A)
    if l <= 1 {
        return 0
    }
    sl := len(A[0])
    ans := 0
    for j := 0; j < sl; j++ {
        for i := 1; i < l; i++ {
            if A[i][j] < A[i-1][j] {
                ans += 1
                break
            }
        }
    }
    return ans
}
```