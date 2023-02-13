# 1641. Count Sorted Vowel Strings

### Medium

Given an integer n, return the number of strings of length n that consist only of vowels (a, e, i, o, u) and are lexicographically sorted.

A string s is lexicographically sorted if for all valid i, s[i] is the same as or comes before s[i+1] in the alphabet.

### Example 1:

Input: n = 1
Output: 5
Explanation: The 5 sorted strings that consist of vowels only are ["a","e","i","o","u"].

### Example 2:

Input: n = 2
Output: 15
Explanation: The 15 sorted strings that consist of vowels only are
["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"].
Note that "ea" is not a valid string since 'e' comes after 'a' in the alphabet.

### Example 3:

Input: n = 33
Output: 66045

Constraints:

1 <= n <= 50

```go
/*
杨辉三角：n = 1时候，a e i o u（反序）分别出现：
1 1 1  1  1 
1 2 3  4  5
1 3 6  10 15 
1 4 10 20 35
1 5 15 35 70

O(n) time.
*/
func countVowelStrings(n int) int {
	vowel := []int{1, 1, 1, 1, 1}
	for i := 0; i < n; i++ {
		for j := 1; j < 5; j++ {
			vowel[j] += vowel[j-1]
		}
	}
	return vowel[4]
}

```