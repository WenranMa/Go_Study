# 917_仅反转字母_Reverse Only Letters

Given a string S, return the "reversed" string where all characters that are not a letter stay in the same place, and all letters reverse their positions.

Example:

    Input: "ab-cd"
    Output: "dc-ba"

    Input: "a-bC-dEf-ghIj"
    Output: "j-Ih-gfE-dCba"

    Input: "Test1ng-Leet=code-Q!"
    Output: "Qedo1ct-eeLg=ntse-T!"

### 解：

前后两个指针移动，指向的字符不是字母，则指针移动，如果指向的字符都是字母，则交换再移动。

```go
func reverseOnlyLetters(S string) string {
    chars := []byte(S)
    l := len(chars)
    for i, j := 0, l-1; i < j; {
        if !isChar(chars[i]) {
            i++
        }
        if !isChar(chars[j]) {
            j--
        }
        if isChar(chars[i]) && isChar(chars[j]) {
            chars[i], chars[j] = chars[j], chars[i]
            i++
            j--
        }
    }
    return string(chars)
}
func isChar(c byte) bool {
    return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}
```