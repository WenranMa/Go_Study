# 804_唯一摩斯码单词_Unique_Morse_Code_Words

International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes, as follows: "a" maps to ".-", "b" maps to "-...", "c" maps to "-.-.", and so on.

For convenience, the full table for the 26 letters of the English alphabet is given below:   
[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]

Now, given a list of words, each word can be written as a concatenation of the Morse code of each letter. For example, "cba" can be written as "-.-..--...", (which is the concatenation "-.-." + "-..." + ".-"). We'll call such a concatenation, the transformation of a word.  
Return the number of different transformations among all words we have.

Example:  
    Input: words = ["gin", "zen", "gig", "msg"]  
    Output: 2  
    Explanation:   
    The transformation of each word is:  
    "gin" -> "--...-."  
    "zen" -> "--...-."  
    "gig" -> "--...--."  
    "msg" -> "--...--."  
    There are 2 different transformations, "--...-." and "--...--.".

### 解：

```go
func uniqueMorseRepresentations(words []string) int {
	morse := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]int)
	for _, w := range words {
		var s strings.Builder
		for _, c := range w {
			s.WriteString(morse[c-97])
		}
		m[s.String()] = 1
	}
	return len(m)
}
```