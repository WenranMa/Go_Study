# 884_两句话中不常见单词_Uncommon_Words_from_Two_Sentences

We are given two sentences A and B. (A sentence is a string of space separated words.  Each word consists only of lowercase letters.) A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence. Return a list of all uncommon words. You may return the list in any order.

Example:

    Input: A = "this apple is sweet", B = "this apple is sour"  
    Output: ["sweet","sour"]  
    Input: A = "apple apple", B = "banana"  
    Output: ["banana"]  

Note:  
- 0 <= A.length <= 200  
- 0 <= B.length <= 200  
- A and B both contain only spaces and lowercase letters.  

### 解：

```go
func uncommonFromSentences(A string, B string) []string {
    m := make(map[string]int)
    res := []string{}
    a := strings.Split(A, " ")
    b := strings.Split(B, " ")
    for _, s := range a {
        m[s] += 1
    }
    for _, s := range b {
        m[s] += 1
    }
    for k, _ := range m {
        if m[k] == 1 {
            res = append(res, k)
        }
    }
    return res
}
```