# 744_比目标字符大的最小字符_Find_Smallest_Letter_Greater_Than_Target
Given a list of sorted characters letters containing only lowercase letters, and given a target letter target, find the smallest element in the list that is larger than the given target.

Letters also wrap around. For example, if the target is target = 'z' and letters = ['a', 'b'], the answer is 'a'.

### Examples:

    Input:  
    letters = ["c", "f", "j"]  
    target = "a"  
    Output: "c"  

    Input:  
    letters = ["c", "f", "j"]  
    target = "c"  
    Output: "f"  

    Input:  
    letters = ["c", "f", "j"]  
    target = "d"  
    Output: "f"  

    Input:  
    letters = ["c", "f", "j"]  
    target = "g"  
    Output: "j"  

    Input:  
    letters = ["c", "f", "j"]  
    target = "j"  
    Output: "c"

    Input:  
    letters = ["c", "f", "j"]  
    target = "k"  
    Output: "c"  

Note:
letters has a length in range [2, 10000].
letters consists of lowercase letters, and contains at least 2 unique letters.
target is a lowercase letter.

### 解：

```go
func nextGreatestLetter(letters []byte, target byte) byte {
    l := 0
    r := len(letters) - 1
    for l < r {
        m := (l + r) / 2
        if letters[m] > target { //如果大于应保留m
            r = m
        } else if letters[m] <= target {
            l = m + 1
        }
    } // l should be equal to (l+r)/2
    if letters[l] <= target {
        return letters[0]
    }
    return letters[l]
}
```