# 709_转成小写字母_To_Lower Case
Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.

Example:  
Input: "Hello"  
Output: "hello"  
Input: "here"  
Output: "here"  
Input: "LOVELY"  
Output: "lovely"

```go
func toLowerCase(str string) string {
	m := make(map[byte]byte)
	for i := 0; i < 26; i++ {
		m[byte(i+65)] = byte(i + 97) // 65 is A, 97 is a
	}
	chars := []byte(str)
	for i, r := range chars {
		if r >= 65 && r <= 90 {
			chars[i] = m[r]
		}
	}
	return string(chars)
}
```

```go
func toLowerCase(str string) string {
	chars := []byte(str)
	for i, r := range chars {
		if r >= 65 && r <= 90 {
			chars[i] = chars[i] + 32
		}
	}
	return string(chars)
}
```