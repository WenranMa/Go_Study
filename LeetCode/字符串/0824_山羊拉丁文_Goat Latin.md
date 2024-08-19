# 824_山羊拉丁文_Goat Latin

A sentence S is given, composed of words separated by spaces. Each word consists of lowercase and uppercase letters only. We would like to convert the sentence to "Goat Latin" (a made-up language similar to Pig Latin.)

The rules of Goat Latin are as follows:  
If a word begins with a vowel (a, e, i, o, or u), append "ma" to the end of the word. For example, the word 'apple' becomes 'applema'. If a word begins with a consonant (i.e. not a vowel), remove the first letter and append it to the end, then add "ma". For example, the word "goat" becomes "oatgma".  
Add one letter 'a' to the end of each word per its word index in the sentence, starting with 1. For example, the first word gets "a" added to the end, the second word gets "aa" added to the end and so on. Return the final sentence representing the conversion from S to Goat Latin. 

Example:

    Input: "I speak Goat Latin"  
    Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"  
    Input: "The quick brown fox jumped over the lazy dog"  
    Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"  
 
Notes:
S contains only uppercase, lowercase and spaces. Exactly one space between each word.  
1 <= S.length <= 150.

### 解：

模拟

```go
func toGoatLatin(S string) string {
    arr := strings.Split(S, " ")
    ws := []string{}
    for i, w := range arr {
        out := encode(w)
        for j := 0; j <= i; j++ {
            out = out + "a"
        }
        ws = append(ws, out)
    }
    ans := strings.Join(ws, " ")
    return ans
}

func encode(in string) string {
    if in == "" {
        return in
    }
    out := ""
    if in[0] == 'a' || in[0] == 'e' || in[0] == 'i' || in[0] == 'o' || in[0] == 'u' || in[0] == 'A' || in[0] == 'E' || in[0] == 'I' || in[0] == 'O' || in[0] == 'U' {
        out = in + "ma"
    } else {
        out = in[1:len(in)] + string(in[0]) + "ma"
    }
    return out
}
```