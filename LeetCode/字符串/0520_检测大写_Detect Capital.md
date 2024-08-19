# 520_检测大写_Detect Capital
Given a word, you need to judge whether the usage of capitals in it is right or not. We define the usage of capitals in a word to be right when one of the following cases holds:

All letters in this word are capitals, like "USA".  
All letters in this word are not capitals, like "leetcode".  
Only the first letter in this word is capital if it has more than one letter, like "Google".  
Otherwise, we define that this word doesn't use capitals in a right way.  

Example:  

    Input: "USA"  
    Output: True  
    Input: "FlaG"  
    Output: False  

Note: The input will be a non-empty word consisting of uppercase and lowercase latin letters.

### 解：

```go
func detectCapitalUse(word string) bool {
    l := len(word)
    if l <= 1 {
        return true
    }
    if word[0] >= 65 && word[0] <= 90 {
        if word[1] >= 65 && word[1] <= 90 {
            for i := 1; i < l; i++ {
                if word[i] >= 97 && word[i] <= 122 {
                    return false
                }
            }
        } else if word[1] >= 97 && word[1] <= 122 {
            for i := 1; i < l; i++ {
                if word[i] >= 65 && word[i] <= 90 {
                    return false
                }
            }
        }
    } else if word[0] >= 97 && word[0] <= 122 {  //全小写返回true, 因为没用到大写也算正确
        for i := 1; i < l; i++ {
            if word[i] >= 65 && word[i] <= 90 {
                return false
            }
        }
    }
    return true
}
```