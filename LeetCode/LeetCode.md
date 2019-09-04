# LeetCode

### String

##### 1108. Defanging an IP Address

    Given a valid (IPv4) IP address, return a defanged version of that IP address.
    A defanged IP address replaces every period "." with "[.]".

    Example:
    Input: address = "255.100.50.0"
    Output: "255[.]100[.]50[.]0"
```go
func defangIPaddr(address string) string {
    return strings.Replace(address, ".", "[.]", -1)
}
```

##### 1021. Remove Outermost Parentheses
A valid parentheses string is either empty (""), "(" + A + ")", or A + B, where A and B are valid parentheses strings, and + represents string concatenation.  For example, "", "()", "(())()", and "(()(()))" are all valid parentheses strings.

A valid parentheses string S is primitive if it is nonempty, and there does not exist a way to split it into S = A+B, with A and B nonempty valid parentheses strings.

Given a valid parentheses string S, consider its primitive decomposition: S = P_1 + P_2 + ... + P_k, where P_i are primitive valid parentheses strings.

Return S after removing the outermost parentheses of every primitive string in the primitive decomposition of S.

Example:

    Input: "(()())(())"
    Output: "()()()"
    Explanation:
    The input string is "(()())(())", with primitive decomposition "(()())" + "(())".
    After removing outer parentheses of each part, this is "()()" + "()" = "()()()".

    Input: "(()())(())(()(()))"
    Output: "()()()()(())"
    Explanation:
    The input string is "(()())(())(()(()))", with primitive decomposition "(()())" + "(())" + "(()(()))".
    After removing outer parentheses of each part, this is "()()" + "()" + "()(())" = "()()()()(())".

    Input: "()()"
    Output: ""
    Explanation:
    The input string is "()()", with primitive decomposition "()" + "()".
    After removing outer parentheses of each part, this is "" + "" = "".

 Note:

    S.length <= 10000
    S[i] is "(" or ")"
    S is a valid parentheses string
```go
func removeOuterParentheses(S string) string {
    counter := 0
    head := 1
    res := ""
    for i, c := range S {
        if c == '(' {
            counter++
        } else {
            counter--
        }
        if counter == 0 {
            res += S[head:i]
            head = i + 2
        }
    }
    return res
}
```

##### 763. Partition Labels
A string S of lowercase letters is given. We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.

Example:

    Input: S = "ababcbacadefegdehijhklij"
    Output: [9,7,8]
    Explanation:
    The partition is "ababcbaca", "defegde", "hijhklij".
    This is a partition so that each letter appears in at most one part.
    A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.
Note:

    S will have length in range [1, 500].
    S will consist of lowercase letters ('a' to 'z') only.

```go
func partitionLabels(S string) []int {
    l := len(S)
    res := []int{}
    tail := 0
    for head := 0; head < l; head = tail + 1 {
        tail = head
        for j := head; j <= tail; j++ {
            for k := l - 1; k > tail; k-- {
                if S[j] == S[k] && k > tail {
                    tail = k
                }
            }
        }
        res = append(res, tail-head+1)
    }
    return res
}
//O(n^3) ?

//Better solution O(n) time, O(n) space.
func partitionLabels(S string) []int {
    res := []int{}
    tail := 0
    head := 0
    lastIndex := make(map[rune]int) // map[int32]int is also ok.
    for i, c := range S {
        lastIndex[c] = i
    }
    for i, c := range S {
        m := math.Max(float64(tail), float64(lastIndex[c]))
        tail = int(m)
        if i == tail {
            res = append(res, tail-head+1)
            head = tail + 1
        }
    }
    return res
}
```

##### 657.Robot Return to Origin
There is a robot starting at position (0, 0), the origin, on a 2D plane. Given a sequence of its moves, judge if this robot ends up at (0, 0) after it completes its moves.

The move sequence is represented by a string, and the character moves[i] represents its ith move. Valid moves are R (right), L (left), U (up), and D (down). If the robot returns to the origin after it finishes all of its moves, return true. Otherwise, return false.

Note: The way that the robot is "facing" is irrelevant. "R" will always make the robot move to the right once, "L" will always make it move left, etc. Also, assume that the magnitude of the robot's movement is the same for each move.

Example:  
Input: "UD"  
Output: true   
Input: "LL"  
Output: false  

```go
func judgeCircle(moves string) bool {
    x, y := 0, 0
    for _, c := range moves {
        switch c {
        case 'U':
            y++
        case 'D':
            y--
        case 'R':
            x++
        case 'L':
            x--
        default:
            break
        }
    }
    if x == 0 && y == 0 {
        return true
    }
    return false
}
```

##### 929.Unique Email Addresses
Every email consists of a local name and a domain name, separated by the @ sign. For example, in alice@leetcode.com, alice is the local name, and leetcode.com is the domain name.  
Besides lowercase letters, these emails may contain '.'s or '+'s.  

If you add periods ('.') between some characters in the local name part of an email address, mail sent there will be forwarded to the same address without dots in the local name.  For example, "alice.z@leetcode.com" and "alicez@leetcode.com" forward to the same email address.  (Note that this rule does not apply for domain names.). 

If you add a plus ('+') in the local name, everything after the first plus sign will be ignored. This allows certain emails to be filtered, for example m.y+name@email.com will be forwarded to my@email.com.  (Again, this rule does not apply for domain names.). 

It is possible to use both of these rules at the same time.  
Given a list of emails, we send one email to each address in the list.  How many different addresses actually receive mails?   

Example:  
Input: ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"].   
Output:  
Explanation: "testemail@leetcode.com" and "testemail@lee.tcode.com" actually receive mails

```go
func numUniqueEmails(emails []string) int {
    m := make(map[string]int)
    for _, e := range emails {
        l := len(e)
        s := strings.IndexRune(e, '@')
        name := e[0:s]
        addr := e[s:l]
        s = strings.IndexRune(name, '+')
        if s != -1 {
            name = e[0:s]
        }
        name = strings.Replace(name, ".", "", -1)
        mail := name + "@" + addr
        m[mail] = 1
    }
    return len(m)
}
```

##### 709. To Lower Case
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
    m := make(map[rune]rune)
    for i := 0; i < 26; i++ {
        m[(rune)(i+65)] = (rune)(i + 97)
    }
    runes := []rune(str)
    for i, r := range runes {
        if r >= 65 && r <= 90 {
            runes[i] = m[r]
        }
    }
    return string(runes)
}
```

##### 824.Goat Latin
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

##### 557. Reverse Words in a String III
Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example:  
Input: "Let's take LeetCode contest"  
Output: "s'teL ekat edoCteeL tsetnoc"  
Note: In the string, each word is separated by single space and there will not be any extra space in the string.
```go
func reverseWords(s string) string {
    arr := strings.Split(s, " ")
    ans := []string{}
    for _, w := range arr {
        runes := []rune(w)
        for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
            runes[i], runes[j] = runes[j], runes[i]
        }
        ans = append(ans, string(runes))
    }
    return strings.Join(ans, " ")
}
```

##### 344. Reverse String
Write a function that reverses a string. The input string is given as an array of characters char[]. Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory. You may assume all the characters consist of printable ascii characters.

Example:  
Input: ["h","e","l","l","o"]  
Output: ["o","l","l","e","h"]  
Input: ["H","a","n","n","a","h"]  
Output: ["h","a","n","n","a","H"]  
```go
func reverseString(s []byte) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}
```

##### 942.DI String Match
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

##### 944.Delete Columns to Make Sorted
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

##### 806. Number of Lines To Write String
We are to write the letters of a given string S, from left to right into lines. Each line has maximum width 100 units, and if writing a letter would cause the width of the line to exceed 100 units, it is written on the next line. We are given an array widths, an array where widths[0] is the width of 'a', widths[1] is the width of 'b', ..., and widths[25] is the width of 'z'.

Now answer two questions: how many lines have at least one character from S, and what is the width used by the last such line? Return your answer as an integer list of length 2.

Example:  
Input:   
widths = [10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10]  
S = "abcdefghijklmnopqrstuvwxyz"  
Output: [3, 60]  
Explanation:   
All letters have the same length of 10. To write all 26 letters,
we need two full lines and one line with 60 units.  
Example:  
Input:   
widths = [4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10]  
S = "bbbcccdddaaa"  
Output: [2, 4]  
Explanation:   
All letters except 'a' have the same length of 10, and 
"bbbcccdddaa" will cover 9 * 10 + 2 * 4 = 98 units.
For the last 'a', it is written on the second line because
there is only 2 units left in the first line.
So the answer is 2 lines, plus 4 units in the second line.
 
Note:  
The length of S will be in the range [1, 1000].  
S will only contain lowercase letters.  
widths is an array of length 26.  
widths[i] will be in the range of [2, 10].  
```go
func numberOfLines(widths []int, S string) []int {
    l := len(S)
    // if l== 0 {
    //     return []int{0,0}
    // }
    n := 0
    row := 1
    last := 0

    for i := 0; i < l; i++ {
        n += widths[S[i]-97]
        if i == l-1 {
            last = n
        }
        if n > 100 {
            row += 1
            i--
            n = 0
        }
    }
    return []int{row, last}
}
```

##### 520. Detect Capital
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
    } else if word[0] >= 97 && word[0] <= 122 {
        for i := 1; i < l; i++ {
            if word[i] >= 65 && word[i] <= 90 {
                return false
            }
        }
    }
    return true
}
```

---


### Math and Bit Manipulation（位运算）

##### 461. Hamming Distance
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.  
Given two integers x and y, calculate the Hamming distance.

Note:  
0 ≤ x, y < 231.

Example:  
Input: x = 1, y = 4  
Output: 2    
Explanation:  
1   (0 0 0 1)  
4   (0 1 0 0)  

```go
func hammingDistance(x int, y int) int {
    h := x ^ y
    ans := 0
    for h != 0 {
        if h&1 == 1 {
            ans += 1
        }
        h = h >> 1
    }
    return ans
}
```

##### 476.Number Complement
Given a positive integer, output its complement number. The complement strategy is to flip the bits of its binary representation.  

Note:
The given integer is guaranteed to fit within the range of a 32-bit signed integer.  
You could assume no leading zero bit in the integer’s binary representation.  
Example:  
Input: 5.  
Output: 2.   
Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.  
Input: 1.   
Output: 0.  
Explanation: The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.
```go
//Solution: 00000101 ^ 00000111
func findComplement(num int) int {
    n := num
    ones := 0
    for n != 0 {
        n = n >> 1
        ones = ones<<1 + 1
    }
    return num ^ ones
}
```

##### 136.Single Number
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:  
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example:  
Input: [2,2,1].   
Output: 1.  
Input: [4,1,2,1,2].  
Output: 4
```go
func singleNumber(nums []int) int {
    ans := 0
    for _, n := range nums {
        ans = ans ^ n
    }
    return ans
}
```

##### 191.Number of 1 Bits
Write a function that takes an unsigned integer and return the number of '1' bits it has (also known as the Hamming weight).

Example 1:  
Input: 00000000000000000000000000001011   
Output: 3   
Input: 00000000000000000000000010000000  
Output: 1  
Input: 11111111111111111111111111111101  
Output: 31  

Note:  
Note that in some languages such as Java, there is no unsigned integer type. In this case, the input will be given as signed integer type and should not affect your implementation, as the internal binary representation of the integer is the same whether it is signed or unsigned.
In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 3 above the input represents the signed integer -3.

```go
func hammingWeight(num uint32) int {
    ans := 0
    for num != 0 {
        ans += (int)(num & 1)
        num = num >> 1
    }
    return ans
}
```

##### 693.Binary Number with Alternating Bits
Given a positive integer, check whether it has alternating bits: namely, if two adjacent bits will always have different values.

Example:  
Input: 5  
Output: True   
Explanation: The binary representation of 5 is: 101

Input: 7  
Output: False  
Explanation: The binary representation of 7 is: 111.

Input: 11  
Output: False  
Explanation: The binary representation of 11 is: 1011.

Input: 10  
Output: True  
Explanation: The binary representation of 10 is: 1010.
```go
func hasAlternatingBits(n int) bool {
    for n != 0 {
        a := (n & 1) == 0
        b := (n & 2) == 0
        if a == b {
            return false
        }
        n = n >> 1
    }
    return true
}
```

##### 728.Self Dividing Numbers
A self-dividing number is a number that is divisible by every digit it contains. For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0. Also, a self-dividing number is not allowed to contain the digit zero. Given a lower and upper number bound, output a list of every possible self dividing number, including the bounds if possible.

Example:  
Input:  
left = 1, right = 22  
Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]  

Note: The boundaries of each input argument are 1 <= left <= right <= 10000.
```go
func selfDividingNumbers(left int, right int) []int {
    ans := []int{}
    for i := left; i <= right; i++ {
        if isDividingNumber(i) == true {
            ans = append(ans, i)
        }
    }
    return ans
}

func isDividingNumber(num int) bool {
    bits := []int{}
    n := num
    for n > 0 {
        b := n % 10
        if b == 0 {
            return false
        }
        n = n / 10
        bits = append(bits, b)
    }
    for _, b := range bits {
        if num%b != 0 {
            return false
        }
    }
    return true
}
```

##### 509. Fibonacci Number
The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), for N > 1.
Given N, calculate F(N).

Example:  
Input: 2  
Output: 1  
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.  
Input: 3  
Output: 2  
Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.  
Input: 4  
Output: 3  
Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.  
```go
func fib(N int) int {
    if N == 0 {
        return 0
    }
    if N == 1 {
        return 1
    }
    return fib(N-1) + fib(N-2)
}
```

##### 908.Smallest Range I
Given an array A of integers, for each integer A[i] we may choose any x with -K <= x <= K, and add x to A[i]. After this process, we have some array B. Return the smallest possible difference between the maximum value of B and the minimum value of B.

Example:  
Input: A = [1], K = 0  
Output: 0  
Explanation: B = [1]  
Input: A = [0,10], K = 2  
Output: 6  
Explanation: B = [2,8]  
Input: A = [1,3,6], K = 3   
Output: 0   
Explanation: B = [3,3,3] or B = [4,4,4]   
 
Note:  
1 <= A.length <= 10000  
0 <= A[i] <= 10000  
0 <= K <= 10000  
```go
func smallestRangeI(A []int, K int) int {
    min, max := 10001, 0
    for _, m := range A {
        if m < min {
            min = m
        }
        if m > max {
            max = m
        }
    }
    ans := max - min - 2*K
    if ans < 0 {
        return 0
    }
    return ans
}
```

##### 412.Fizz Buzz
Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

Example:   
n = 15,   
Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
```go
func fizzBuzz(n int) []string {
    ans := []string{}
    for i := 1; i <= n; i++ {
        f1 := i % 3
        f2 := i % 5
        if f1 == 0 && f2 != 0 {
            ans = append(ans, "Fizz")
        } else if f1 != 0 && f2 == 0 {
            ans = append(ans, "Buzz")
        } else if f1 == 0 && f2 == 0 {
            ans = append(ans, "FizzBuzz")
        } else {
            ans = append(ans, strconv.Itoa(i))
        }
    }
    return ans
}
```

##### 762. Prime Number of Set Bits in Binary Representation
Given two integers L and R, find the count of numbers in the range [L, R] (inclusive) having a prime number of set bits in their binary representation.

(Recall that the number of set bits an integer has is the number of 1s present when written in binary. For example, 21 written in binary is 10101 which has 3 set bits. Also, 1 is not a prime.)

Example:  
Input: L = 6, R = 10  
Output: 4  
Explanation:  
6 -> 110 (2 set bits, 2 is prime)  
7 -> 111 (3 set bits, 3 is prime)   
9 -> 1001 (2 set bits , 2 is prime)   
10->1010 (2 set bits , 2 is prime)   

Input: L = 10, R = 15  
Output: 5   
Explanation:   
10 -> 1010 (2 set bits, 2 is prime)   
11 -> 1011 (3 set bits, 3 is prime)  
12 -> 1100 (2 set bits, 2 is prime)  
13 -> 1101 (3 set bits, 3 is prime)  
14 -> 1110 (3 set bits, 3 is prime)  
15 -> 1111 (4 set bits, 4 is not prime)  

Note:  
L, R will be integers L <= R in the range [1, 10^6].  
R - L will be at most 10000.  
```go
func countPrimeSetBits(L int, R int) int {
    ans := 0
    for i := L; i <= R; i++ {
        b := countOnes(i)
        if checkPrime(b) {
            ans += 1
        }
    }
    return ans
}
func checkPrime(n int) bool {
    if n == 2 || n == 3 || n == 5 || n == 7 || n == 11 || n == 13 || n == 17 || n == 19 {
        return true
    }
    return false
}
func countOnes(n int) int {
    b := 0
    for n != 0 {
        if n&1 == 1 {
            b++
        }
        n = n >> 1
    }
    return b
}
```

---

### Array 数组 and Two Pass 双指针

##### 1046. Last Stone Weight
We have a collection of rocks, each rock has a positive integer weight.

Each turn, we choose the two heaviest rocks and smash them together.  Suppose the stones have weights x and y with x <= y.  The result of this smash is:

If x == y, both stones are totally destroyed;
If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.
At the end, there is at most 1 stone left.  Return the weight of this stone (or 0 if there are no stones left.)

Example:

    Input: [2,7,4,1,8,1]
    Output: 1
    Explanation:
    We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
    we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
    we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
    we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of last stone.

Note:

    1 <= stones.length <= 30
    1 <= stones[i] <= 1000

```go
func lastStoneWeight(stones []int) int {
    for len(stones) > 1 {
        sort.Ints(stones)
        l := len(stones)
        m1 := 0
        m2 := 0
        if l > 1 {
            m1 = stones[l-1]
            m2 = stones[l-2]
            stones = stones[0 : l-2]
            if m1 > m2 {
                stones = append(stones, m1-m2)
            }
        }
    }
    if len(stones) == 1 {
        return stones[0]
    }
    return 0
}
```

##### 1051. Height Checker
Students are asked to stand in non-decreasing order of heights for an annual photo.

Return the minimum number of students not standing in the right positions.  (This is the number of students that must move in order for all students to be standing in non-decreasing order of height.)

Example:

    Input: [1,1,4,2,1,3]
    Output: 3
    Explanation:
    Students with heights 4, 3 and the last 1 are not standing in the right positions.

Note:

    1 <= heights.length <= 100
    1 <= heights[i] <= 100
```go
func heightChecker(heights []int) int {
    sorted := make([]int, len(heights))
    copy(sorted, heights)
    sort.Ints(sorted)
    res := 0
    for i, e := range heights {
        if e != sorted[i] {
            res += 1
        }
    }
    return res
}
```

##### 1002. Find Common Characters
Given an array A of strings made only from lowercase letters, return a list of all characters that show up in all strings within the list (including duplicates).  For example, if a character occurs 3 times in all strings but not 4 times, you need to include that character three times in the final answer.

You may return the answer in any order.

Example:

    Input: ["bella","label","roller"]
    Output: ["e","l","l"]

    Input: ["cool","lock","cook"]
    Output: ["c","o"]

Note:

    1 <= A.length <= 100
    1 <= A[i].length <= 100
    A[i][j] is a lowercase letter

```go
// Use one array "count" to store the frequency fo char in words.
// If the chars frequency decrease, update "count"
func commonChars(A []string) []string {
    count := [26]int{}
    for i, _ := range count {
        count[i] = math.MaxUint16
    }
    for _, word := range A {
        chars := [26]int{}
        for _, c := range word {
            chars[c-'a'] += 1
        }
        for i := 0; i < 26; i++ {
            if chars[i] < count[i] {
                count[i] = chars[i]
            }
        }
    }
    res := []string{}
    for i, n := range count {
        for j := 1; j <= n; j++ {
            res = append(res, string('a'+i))
        }
    }
    return res
}
```

##### 167.Two Sum II - Input array is sorted
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number. The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:  
Your returned answers (both index1 and index2) are not zero-based. You may assume that each input would have exactly one solution and you may not use the same element twice.

Example:  
Input: numbers = [2,7,11,15], target = 9  
Output: [1,2]  
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.  
```go
func twoSum(numbers []int, target int) []int {
    l, r := 0, len(numbers)-1
    for l < r {
        if numbers[l]+numbers[r] < target {
            l++
        } else if numbers[l]+numbers[r] > target {
            r--
        } else {
            return []int{l + 1, r + 1}
        }
    }
    return []int{-1, -1}
}
```

##### 448.Find All Numbers Disappeared in an Array
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once. Find all the elements of [1, n] inclusive that do not appear in this array. Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:  
Input: [4,3,2,7,8,2,3,1]  
Output: [5,6]
```go
func findDisappearedNumbers(nums []int) []int {
    l := len(nums)
    ans := []int{}
    for i := 0; i < l; {
        if nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        } else {
            i++
        }
    }
    for i, n := range nums {
        if n != i+1 {
            ans = append(ans, i+1)
        }
    }
    return ans
}
```

##### 442. Find All Duplicates in an Array
Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once. Find all the elements that appear twice in this array. Could you do it without extra space and in O(n) runtime?

Example:  
Input: [4,3,2,7,8,2,3,1]  
Output: [2,3]
```go 
func findDuplicates(nums []int) []int {
    l := len(nums)
    ans := []int{}
    for i := 0; i < l; {
        if nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        } else {
            i++
        }
    }
    for i, n := range nums {
        if n != i+1 {
            ans = append(ans, n)
        }
    }
    return ans
}
```

##### 268.Missing Number
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example:  
Input: [3,0,1]   
Output: 2  
Input: [9,6,4,2,3,5,7,0,1]  
Output: 8  

Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
```go
func missingNumber(nums []int) int {
    l := len(nums)
    for i := 0; i < l; {
        if nums[i] < l && nums[i] != i {
            nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
        } else {
            i++
        }
    }
    for i, n := range nums {
        if n != i {
            return i
        }
    }
    return l
}
```

##### 905.Sort Array By Parity
Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A. You may return any answer array that satisfies this condition.

Example 1:  
Input: [3,1,2,4]  
Output: [2,4,3,1]  
The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.  
 
Note:  
1 <= A.length <= 5000  
0 <= A[i] <= 5000  

```go
func sortArrayByParity(A []int) []int {
    if A == nil || len(A) == 0 || len(A) == 1 {
        return A
    }
    l := len(A)
    i, j := 0, l-1
    for i < j {
        for i < j && A[i]%2 == 0 {
            i++
        }
        for i < j && A[j]%2 != 0 {
            j--
        }
        A[i], A[j] = A[j], A[i]
    }
    return A
}
```

##### 977.Squares of a Sorted Array
Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.

Example:  
Input: [-4,-1,0,3,10]  
Output: [0,1,9,16,100]  
Input: [-7,-3,2,3,11]  
Output: [4,9,9,49,121]  
 
Note:  
1 <= A.length <= 10000  
-10000 <= A[i] <= 10000  
A is sorted in non-decreasing order.  
```go
func sortedSquares(A []int) []int {
    if A == nil || len(A) == 0 {
        return A
    }
    l := len(A)
    i := 0 // for positive
    for i < l && A[i] < 0 {
        i++
    }
    j := i - 1 //for negative
    ans := []int{}
    for i < l || j >= 0 {
        if j < 0 || i < l && A[i]*A[i] <= A[j]*A[j] {
            ans = append(ans, A[i]*A[i])
            i++
        } else {
            ans = append(ans, A[j]*A[j])
            j--
        }
    }
    return ans
}
```

##### 852.Peak Index in a Mountain Array
Let's call an array A a mountain if the following properties hold:

A.length >= 3  
There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]  
Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].  

Example:  
Input: [0,1,0]  
Output: 1  
Input: [0,2,1,0]  
Output: 1  

Note:  
3 <= A.length <= 10000  
0 <= A[i] <= 10^6  
A is a mountain, as defined above.  
```go
func peakIndexInMountainArray(A []int) int {
    l := len(A)
    for i := 0; i < l-1; i++ {
        if A[i] > A[i+1] {
            return i
        }
    }
    return 0
}
```

##### 832. Flipping an Image
Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image. To flip an image horizontally means that each row of the image is reversed.  For example, flipping [1, 1, 0] horizontally results in [0, 1, 1]. To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].

Example:  
Input: [[1,1,0],[1,0,1],[0,0,0]]   
Output: [[1,0,0],[0,1,0],[1,1,1]]    
Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]  
Input: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]  
Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]  
Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]  

Notes:  
1 <= A.length = A[0].length <= 20  
0 <= A[i][j] <= 1
```go
func flipAndInvertImage(A [][]int) [][]int {
    for _, row := range A {
        l := len(row)
        for j, k := 0, l-1; j <= k; {
            row[j], row[k] = row[k], row[j]
            row[j], row[k] = row[j]^1, row[k]^1
            j++
            k--
        }
    }
    return A
}
```

##### 985. Sum of Even Numbers After Queries
We have an array A of integers, and an array queries of queries. For the i-th query val = queries[i][0], index = queries[i][1], we add val to A[index].  Then, the answer to the i-th query is the sum of the even values of A. (Here, the given index = queries[i][1] is a 0-based index, and each query permanently modifies the array A.) Return the answer to all queries.  Your answer array should have answer[i] as the answer to the i-th query.

Example:  
Input: A = [1,2,3,4], queries = [[1,0],[-3,1],[-4,0],[2,3]]  
Output: [8,6,2,4]  
Explanation:   
At the beginning, the array is [1,2,3,4].  
After adding 1 to A[0], the array is [2,2,3,4], and the sum of even values is 2 + 2 + 4 = 8.  
After adding -3 to A[1], the array is [2,-1,3,4], and the sum of even values is 2 + 4 = 6.  
After adding -4 to A[0], the array is [-2,-1,3,4], and the sum of even values is -2 + 4 = 2.  
After adding 2 to A[3], the array is [-2,-1,3,6], and the sum of even values is -2 + 6 = 4.  
 
Note:  
1 <= A.length <= 10000  
-10000 <= A[i] <= 10000  
1 <= queries.length <= 10000  
-10000 <= queries[i][0] <= 10000  
0 <= queries[i][1] < A.length  
```go
func sumEvenAfterQueries(A []int, queries [][]int) []int {
    ans := []int{}
    for _, query := range queries {
        A[query[1]] += query[0]
        an := 0
        for _, a := range A {
            if a&1 == 0 {
                an += a
            }
        }
        ans = append(ans, an)
    }
    return ans
}
```

##### 561.Array Partition I
Given an array of 2n integers, your task is to group these integers into n pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.

Example:  
Input: [1,4,3,2]. 
Output: 4. 
Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).

Note:  
n is a positive integer, which is in the range of [1, 10000].  
All the integers in the array will be in the range of [-10000, 10000].  
```go
func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    l := len(nums)
    ans := 0
    for i := 0; i < l; i += 2 {
        ans += nums[i]
    }
    return ans
}
```

##### 922.Sort Array By Parity II
Given an array A of non-negative integers, half of the integers in A are odd, and half of the integers are even. Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is even, i is even. You may return any answer array that satisfies this condition.

Example:  
Input: [4,2,5,7].  
Output: [4,5,2,7].  
Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.  
 
Note:
2 <= A.length <= 20000.  
A.length % 2 == 0.  
0 <= A[i] <= 1000.  
```go
func sortArrayByParityII(A []int) []int {
    i, j := 0, 1
    l := len(A)
    for i < l || j < l {
        for ; i < l; i = i + 2 {
            if A[i]&1 == 1 {
                break
            }
        }
        for ; j < l; j = j + 2 {
            if A[j]&1 == 0 {
                break
            }
        }
        if i < l && j < l {
            A[i], A[j] = A[j], A[i]
        }
    }
    return A
}
```

##### 883.Projection Area of 3D Shapes
On a N * N grid, we place some 1 * 1 * 1 cubes that are axis-aligned with the x, y, and z axes. Each value v = grid[i][j] represents a tower of v cubes placed on top of grid cell (i, j). Now we view the projection of these cubes onto the xy, yz, and zx planes. A projection is like a shadow, that maps our 3 dimensional figure to a 2 dimensional plane. Here, we are viewing the "shadow" when looking at the cubes from the top, the front, and the side. Return the total area of all three projections.

Example:

Input: [[2]]  
Output: 5  
Input: [[1,2],[3,4]]  
Output: 17  
Explanation:   
Here are the three projections ("shadows") of the shape made with each axis-aligned plane.  
Input: [[1,0],[0,2]]  
Output: 8   
Input: [[1,1,1],[1,0,1],[1,1,1]]  
Output: 14  
Input: [[2,2,2],[2,1,2],[2,2,2]]  
Output: 21   
 
Note:   
1 <= grid.length = grid[0].length <= 50  
0 <= grid[i][j] <= 50

```go
func projectionArea(grid [][]int) int {
    //每行最大元素的和 + 每列最大元素的和 + 非零元素个数
    ans := 0
    l := len(grid)
    for i := 0; i < l; i++ {
        m := 0
        for j := 0; j < l; j++ {
            if grid[i][j] > m {
                m = grid[i][j]
            }
            if grid[i][j] != 0 {
                ans += 1
            }
        }
        ans += m
    }
    for j := 0; j < l; j++ {
        m := 0
        for i := 0; i < l; i++ {
            if grid[i][j] > m {
                m = grid[i][j]
            }
        }
        ans += m
    }
    return ans
}
```

##### 821.Shortest Distance to a Character
Given a string S and a character C, return an array of integers representing the shortest distance from the character C in the string.

Example:  
Input: S = "loveleetcode", C = 'e'   
Output: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]   

Note:  
S string length is in [1, 10000].   
C is a single character, and guaranteed to be in string S.   
All letters in S and C are lowercase.   
```go
func shortestToChar(S string, C byte) []int {
    l := len(S)
    ans := []int{}
    pos := []int{}
    for i, c := range S {
        if c == rune(C) {
            pos = append(pos, i)
        }
    }
    for i, _ := range S {
        min := l
        for _, p := range pos {
            d := 0
            if i < p {
                d = p - i
            } else {
                d = i - p
            }
            if d < min {
                min = d
            }
        }
        ans = append(ans, min)
    }
    return ans
}
```

##### 463.Island Perimeter
You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.

Example:  
Input:  
[[0,1,0,0],  
 [1,1,1,0],  
 [0,1,0,0],  
 [1,1,0,0]]

Output: 16

Explanation: The perimeter is the 16 yellow stripes in the image below:
```go
func islandPerimeter(grid [][]int) int {
    h := len(grid)
    w := len(grid[0])
    ans := 0
    for i, _ := range grid {
        for j, _ := range grid[i] {
            if grid[i][j] == 1 {
                edge := 4 - checkNeighbor(grid, i, j, h, w)
                ans += edge
            }

        }
    }
    return ans
}

func checkNeighbor(grid [][]int, i, j, h, w int) int {
    n := 0
    if i-1 >= 0 && grid[i-1][j] == 1 {
        n += 1
    }
    if i+1 < h && grid[i+1][j] == 1 {
        n += 1
    }
    if j-1 >= 0 && grid[i][j-1] == 1 {
        n += 1
    }
    if j+1 < w && grid[i][j+1] == 1 {
        n += 1
    }
    return n
}
```

##### 976.Largest Perimeter Triangle
Given an array A of positive lengths, return the largest perimeter of a triangle with non-zero area, formed from 3 of these lengths.

If it is impossible to form any triangle of non-zero area, return 0.

Example 1:  
Input: [2,1,2]   
Output: 5    
Input: [1,2,1]    
Output: 0      
Input: [3,2,3,4]    
Output: 10   
Input: [3,6,2,3]    
Output: 8    
 
Note:   
3 <= A.length <= 10000   
1 <= A[i] <= 10^6    
```go
func largestPerimeter(A []int) int {
    l := len(A)
    if l <= 2 {
        return 0
    }
    sort.Ints(A)
    ans := 0
    for i := l - 1; i >= 2; i-- {

        if A[i-1]+A[i-2] > A[i] {
            ans = A[i-2] + A[i-1] + A[i]
            break
        }
    }
    return ans
}
```

##### 896.Monotonic Array
An array is monotonic if it is either monotone increasing or monotone decreasing.

An array A is monotone increasing if for all i <= j, A[i] <= A[j].  An array A is monotone decreasing if for all i <= j, A[i] >= A[j].

Return true if and only if the given array A is monotonic.

Example:   
Input: [1,2,2,3]   
Output: true   
Input: [6,5,4,4]    
Output: true    
Input: [1,3,2]    
Output: false    
Input: [1,1,1]    
Output: true   
 
Note:   
1 <= A.length <= 50000    
-100000 <= A[i] <= 100000    
```go
func isMonotonic(A []int) bool {
    l := len(A)
    increasing := true
    decreasing := true
    for i := 1; i < l; i++ {
        if A[i-1] > A[i] {
            increasing = false
        }
        if A[i-1] < A[i] {
            decreasing = false
        }
    }
    return increasing || decreasing
}
```

##### 485.Max Consecutive Ones
Given a binary array, find the maximum number of consecutive 1s in this array.

Example:   
Input: [1,1,0,1,1,1]   
Output: 3   
Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.   

Note:   
The input array will only contain 0 and 1.   
The length of input array is a positive integer and will not exceed 10,000   
```go
func findMaxConsecutiveOnes(nums []int) int {
    ans := 0
    c := 0
    for _, e := range nums {
        if e == 1 {
            c += 1
        } else {
            if ans < c {
                ans = c
            }
            c = 0
        }
    }
    if ans < c {
        ans = c
    }
    return ans
}
```

##### 704.Binary Search
Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.

Example:   
Input: nums = [-1,0,3,5,9,12], target = 9   
Output: 4   
Explanation: 9 exists in nums and its index is 4   
Input: nums = [-1,0,3,5,9,12], target = 2   
Output: -1   
Explanation: 2 does not exist in nums so return -1 
 
Note:   
You may assume that all elements in nums are unique.   
n will be in the range [1, 10000].   
The value of each element in nums will be in the range [-9999, 9999].   
```go
func search(nums []int, target int) int {
    l := 0
    r := len(nums) - 1
    for l <= r {
        m := (l + r) / 2
        if nums[m] == target {
            return m
        } else if nums[m] < target {
            l = m + 1
        } else if nums[m] > target {
            r = m - 1
        }
    }
    return -1
}
```

##### 744.Find Smallest Letter Greater Than Target
Given a list of sorted characters letters containing only lowercase letters, and given a target letter target, find the smallest element in the list that is larger than the given target.

Letters also wrap around. For example, if the target is target = 'z' and letters = ['a', 'b'], the answer is 'a'.

Examples:  
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


##### 973. K Closest Points to Origin
We have a list of points on the plane.  Find the K closest points to the origin (0, 0). (Here, the distance between two points on a plane is the Euclidean distance.) You may return the answer in any order.  The answer is guaranteed to be unique (except for the order that it is in.) 

Example:  
Input: points = [[1,3],[-2,2]], K = 1   
Output: [[-2,2]]   
Explanation:   
The distance between (1, 3) and the origin is sqrt(10).   
The distance between (-2, 2) and the origin is sqrt(8).   
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.    
We only want the closest K = 1 points from the origin, so the answer is just [[-2,2]].   
Input: points = [[3,3],[5,-1],[-2,4]], K = 2   
Output: [[3,3],[-2,4]]   
(The answer [[-2,4],[3,3]] would also be accepted.)   
 
Note:  
1 <= K <= points.length <= 10000   
-10000 < points[i][0] < 10000   
-10000 < points[i][1] < 10000   
```go
func kClosest(points [][]int, K int) [][]int {
    if len(points) <= K {
        return points
    }
    sort.Slice(points, func(i, j int) bool {
        a := points[i][0]*points[i][0] + points[i][1]*points[i][1]
        b := points[j][0]*points[j][0] + points[j][1]*points[j][1]
        return a < b
    })
    ans := [][]int{}
    for i := 0; i < K; i++ {
        ans = append(ans, points[i])
    }
    return ans
}
```

##### 867.Transpose Matrix
Given a matrix A, return the transpose of A. The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.

Example:  
Input: [[1,2,3],[4,5,6],[7,8,9]]   
Output: [[1,4,7],[2,5,8],[3,6,9]]   

Input: [[1,2,3],[4,5,6]]
Output: [[1,4],[2,5],[3,6]]
 
Note:   
1 <= A.length <= 1000   
1 <= A[0].length <= 1000   
```go
func transpose(A [][]int) [][]int {
    r := len(A)
    if r == 0 {
        return A
    }
    c := len(A[0])
    if c == 0 {
        return A
    }
    ans := [][]int{}
    for i := 0; i < c; i++ {

        row := []int{}
        for j := 0; j < r; j++ {
            row = append(row, A[j][i])
        }
        ans = append(ans, row)
    }
    return ans
}
```

##### 766.Toeplitz Matrix
A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same element. Now given an M x N matrix, return True if and only if the matrix is Toeplitz.
 
Example:  
Input:  
matrix = [  
  [1,2,3,4],  
  [5,1,2,3],  
  [9,5,1,2]  
]  
Output: True  
Explanation:  
In the above grid, the diagonals are:   
"[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]".   
In each diagonal all elements are the same, so the answer is True.   

Input:  
matrix = [  
  [1,2],  
  [2,2]  
]  
Output: False  
Explanation:  
The diagonal "[1, 2]" has different elements.  

Note:  
matrix will be a 2D array of integers.  
matrix will have a number of rows and columns in range [1, 20].  
matrix[i][j] will be integers in range [0, 99].  
```go
func isToeplitzMatrix(matrix [][]int) bool {
    row := len(matrix)
    col := len(matrix[0])
    for r := row - 1; r >= 0; r-- {
        n := matrix[r][0]
        for i, j := r, 0; i < row && j < col; i, j = i+1, j+1 {
            if matrix[i][j] != n {
                return false
            }
        }
    }
    for c := 1; c < col; c++ {
        n := matrix[0][c]
        for i, j := 0, c; i < row && j < col; i, j = i+1, j+1 {

            if matrix[i][j] != n {
                return false
            }
        }
    }
    return true
}
```

##### 807.Max Increase to Keep City Skyline
In a 2 dimensional array grid, each value grid[i][j] represents the height of a building located there. We are allowed to increase the height of any number of buildings, by any amount (the amounts can be different for different buildings). Height 0 is considered to be a building as well. 

At the end, the "skyline" when viewed from all four directions of the grid, i.e. top, bottom, left, and right, must be the same as the skyline of the original grid. A city's skyline is the outer contour of the rectangles formed by all the buildings when viewed from a distance. See the following example.

What is the maximum total sum that the height of the buildings can be increased?

Example:  
Input: grid = [[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]  
Output: 35  
Explanation:   
The grid is:  
[ [3, 0, 8, 4],   
  [2, 4, 5, 7],  
  [9, 2, 6, 3],  
  [0, 3, 1, 0] ]  

The skyline viewed from top or bottom is: [9, 4, 8, 7]  
The skyline viewed from left or right is: [8, 7, 9, 3]

The grid after increasing the height of buildings without affecting skylines is:

gridNew = [ [8, 4, 8, 7],  
            [7, 4, 7, 7],  
            [9, 4, 8, 7],  
            [3, 3, 3, 3] ]  

Notes:  
1 < grid.length = grid[0].length <= 50.  
All heights grid[i][j] are in the range [0, 100].  
All buildings in grid[i][j] occupy the entire grid cell: that is, they are a 1 x 1 x grid[i][j] rectangular prism.
```go
func maxIncreaseKeepingSkyline(grid [][]int) int {
    rowMax := []int{}
    colMax := []int{}
    row := len(grid)
    col := len(grid[0])
    for i := 0; i < row; i++ {
        max := 0
        for j := 0; j < col; j++ {
            if max < grid[i][j] {
                max = grid[i][j]
            }
        }
        rowMax = append(rowMax, max)
    }
    for j := 0; j < col; j++ {
        max := 0
        for i := 0; i < row; i++ {
            if max < grid[i][j] {
                max = grid[i][j]
            }
        }
        colMax = append(colMax, max)
    }
    ans := 0
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            min := int(math.Min(float64(rowMax[i]), float64(colMax[j])))
            if grid[i][j] < min {
                ans += min - grid[i][j]
            }
        }
    }
    return ans
}
```

##### 861. Score After Flipping Matrix
We have a two dimensional matrix A where each value is 0 or 1. A move consists of choosing any row or column, and toggling each value in that row or column: changing all 0s to 1s, and all 1s to 0s. After making any number of moves, every row of this matrix is interpreted as a binary number, and the score of the matrix is the sum of these numbers. Return the highest possible score.

Example:  
Input: [[0,0,1,1],[1,0,1,0],[1,1,0,0]]  
Output: 39  
Explanation:  
Toggled to [[1,1,1,1],[1,0,0,1],[1,1,1,1]].  
0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39  
 
Note:  
1 <= A.length <= 20  
1 <= A[0].length <= 20  
A[i][j] is 0 or 1.  
```go
func matrixScore(A [][]int) int {
    A = flip(A)
    ans := 0
    for _, row := range A {
        r := 0
        for _, n := range row {
            r = r << 1
            r += n
        }
        ans += r
    }
    return ans
}

func flip(A [][]int) [][]int {
    r := len(A)
    c := len(A[0])
    //change 1st col to 1s.
    for i := 0; i < r; i++ {
        if A[i][0] == 0 {
            for j := 0; j < c; j++ {
                A[i][j] ^= 1
            }
        }
    }
    //change the rest 0s to 1s, when 0s are more than 1s
    for j := 1; j < c; j++ {
        n := 0
        for i := 0; i < r; i++ {
            n += A[i][j]
        }
        if n <= r/2 { // 0s more than 1s
            for i := 0; i < r; i++ {
                A[i][j] ^= 1
            }
        }
    }
    return A
}
```

---

### Map 

#### 890. Find and Replace Pattern
You have a list of words and a pattern, and you want to know which words in words matches the pattern.

A word matches the pattern if there exists a permutation of letters p so that after replacing every letter x in the pattern with p(x), we get the desired word.

(Recall that a permutation of letters is a bijection from letters to letters: every letter maps to another letter, and no two letters map to the same letter.)

Return a list of the words in words that match the given pattern.

You may return the answer in any order.

Example:

    Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
    Output: ["mee","aqq"]
    Explanation: "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}.
    "ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation,
    since a and b map to the same letter.

Note:

    1 <= words.length <= 50
    1 <= pattern.length = words[i].length <= 20

```go
//Two maps
// O(N*K) time, space
func findAndReplacePattern(words []string, pattern string) []string {
    res := []string{}
    for _, s := range words {
        if check(s, pattern) {
            res = append(res, s)
        }
    }
    return res
}

func check(s, pattern string) bool {
    mptos := make(map[byte]byte)
    mstop := make(map[byte]byte)
    l := len(s)
    for i := 0; i < l; i++ {
        if _, ok := mptos[pattern[i]]; !ok {
            mptos[pattern[i]] = s[i]
        }
        if _, ok := mstop[s[i]]; !ok {
            mstop[s[i]] = pattern[i]
        }
        if mptos[pattern[i]] != s[i] || mstop[s[i]] != pattern[i] {
            return false
        }
    }
    return true
}
```

##### 1160. Find Words That Can Be Formed by Characters
You are given an array of strings words and a string chars.

A string is good if it can be formed by characters from chars (each character can only be used once).

Return the sum of lengths of all good strings in words.

Example:

    Input: words = ["cat","bt","hat","tree"], chars = "atach"
    Output: 6
    Explanation:
    The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.

    Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
    Output: 10
    Explanation:
    The strings that can be formed are "hello" and "world" so the answer is 5 + 5 = 10.

Note:

    1 <= words.length <= 1000
    1 <= words[i].length, chars.length <= 100
    All strings contain lowercase English letters only.

```go
func countCharacters(words []string, chars string) int {
    res := 0
    for _, w := range words {
        if check(w, chars) {
            res = res + len(w)
        }
    }
    return res
}

func check(w, chars string) bool {
    m := make(map[rune]int)
    for _, c := range chars {
        m[c] += 1
    }
    for _, c := range w {
        m[c] -= 1
        if m[c] < 0 {
            return false
        }
    }
    return true
}
```

##### 1. Two Sum
Given an array of integers, return indices of the two numbers such that they add up to a specific target. You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:  
Given nums = [2, 7, 11, 15], target = 9,  
Because nums[0] + nums[1] = 2 + 7 = 9,  
return [0, 1].  
```go
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    ans := []int{}
    for i, n := range nums {
        x := target - n
        if _, ok := m[x]; ok {
            return []int{m[x], i}
        } else {
            m[n] = i
        }
    }
    return ans
}
```

##### 961.N-Repeated Element in Size 2N Array
In a array A of size 2N, there are N+1 unique elements, and exactly one of these elements is repeated N times. Return the element repeated N times.

Example:  
Input: [1,2,3,3]  
Output: 3  
Input: [2,1,2,5,3,2]  
Output: 2  
Input: [5,1,5,2,5,3,5,4]  
Output: 5  
 
Note:  
4 <= A.length <= 10000  
0 <= A[i] < 10000  
A.length is even  

```go
func repeatedNTimes(A []int) int {
    m := make(map[int]int)
    ans := 0
    for _, n := range A {
        m[n] += 1
        if m[n] > 1 {
            ans = n
            break
        }
    }
    return ans
}
```

##### 771.Jewels and Stones
You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".

Example:  
Input: J = "aA", S = "aAAbbbb"  
Output: 3  
Input: J = "z", S = "ZZ"  
Output: 0  

Note:  
S and J will consist of letters and have length at most 50.  
The characters in J are distinct.
```go
func numJewelsInStones(J string, S string) int {
    m := make(map[rune]int)
    ans := 0
    for _, c := range S {
        m[c] += 1
    }
    for _, c := range J {
        ans += m[c]
    }
    return ans
}
```

##### 804.Unique Morse Code Words
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
```go
func uniqueMorseRepresentations(words []string) int {
    morse := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
    m := make(map[string]int)
    for _, w := range words {
        s := ""
        for _, c := range w {
            s += morse[c-97]
        }
        m[s] = 1
    }
    return len(m)
}
```

##### 884.Uncommon Words from Two Sentences
We are given two sentences A and B. (A sentence is a string of space separated words.  Each word consists only of lowercase letters.) A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence. Return a list of all uncommon words. You may return the list in any order.

Example:  
Input: A = "this apple is sweet", B = "this apple is sour"  
Output: ["sweet","sour"]  
Input: A = "apple apple", B = "banana"  
Output: ["banana"]  

Note:  
0 <= A.length <= 200  
0 <= B.length <= 200  
A and B both contain only spaces and lowercase letters.  

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

##### 500.Keyboard Row
Given a List of words, return the words that can be typed using letters of alphabet on only one row's of American keyboard like the image below.

Example:  
Input: ["Hello", "Alaska", "Dad", "Peace"]  
Output: ["Alaska", "Dad"]  
 
Note:  
You may use one character in the keyboard more than once.  
You may assume the input string will only contain letters of alphabet.  
```go
func findWords(words []string) []string {
    m := make(map[rune]int)

    m['q'] = 1
    m['w'] = 1
    m['e'] = 1
    m['r'] = 1
    m['t'] = 1
    m['y'] = 1
    m['u'] = 1
    m['i'] = 1
    m['o'] = 1
    m['p'] = 1

    m['a'] = 2
    m['s'] = 2
    m['d'] = 2
    m['f'] = 2
    m['g'] = 2
    m['h'] = 2
    m['j'] = 2
    m['k'] = 2
    m['l'] = 2

    m['z'] = 3
    m['x'] = 3
    m['c'] = 3
    m['v'] = 3
    m['b'] = 3
    m['n'] = 3
    m['m'] = 3

    ans := []string{}
    for _, w := range words {
        wl := strings.ToLower(w)
        flag := true
        n := m[rune(wl[0])]
        for _, l := range wl {
            if m[l] != n {
                flag = false
                break
            }
        }
        if flag {
            ans = append(ans, w)
        }
    }
    return ans
}
```

##### 811.Subdomain Visit Count
A website domain like "discuss.leetcode.com" consists of various subdomains. At the top level, we have "com", at the next level, we have "leetcode.com", and at the lowest level, "discuss.leetcode.com". When we visit a domain like "discuss.leetcode.com", we will also visit the parent domains "leetcode.com" and "com" implicitly.

Now, call a "count-paired domain" to be a count (representing the number of visits this domain received), followed by a space, followed by the address. An example of a count-paired domain might be "9001 discuss.leetcode.com".

We are given a list cpdomains of count-paired domains. We would like a list of count-paired domains, (in the same format as the input, and in any order), that explicitly counts the number of visits to each subdomain.

Example 1:  
Input:   
["9001 discuss.leetcode.com"]   
Output:    
["9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"]   
Explanation:   
We only have one website domain: "discuss.leetcode.com". As discussed above, the subdomain "leetcode.com" and "com" will also be visited. So they will all be visited 9001 times.

Example 2:   
Input:    
["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]   
Output:   
["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]   
Explanation:    
We will visit "google.mail.com" 900 times, "yahoo.com" 50 times, "intel.mail.com" once and "wiki.org" 5 times. For the subdomains, we will visit "mail.com" 900 + 1 = 901 times, "com" 900 + 50 + 1 = 951 times, and "org" 5 times.

Notes:   
The length of cpdomains will not exceed 100.   
The length of each domain name will not exceed 100.   
Each address will have either 1 or 2 "." characters.   
The input count in any count-paired domain will not exceed 10000.   
The answer output can be returned in any order.   
```go
func subdomainVisits(cpdomains []string) []string {
    m := make(map[string]int)
    for _, domain := range cpdomains {
        d := strings.Split(domain, " ")
        n, _ := strconv.Atoi(d[0])
        m[d[1]] += n
        for {
            i := strings.IndexRune(d[1], '.')
            if i == -1 {
                break
            }
            d[1] = d[1][i+1:]
            m[d[1]] += n
        }
    }
    ans := []string{}
    for k, v := range m {
        n := strconv.Itoa(v)
        s := n + " " + k
        ans = append(ans, s)
    }
    return ans
}
```

##### 349. Intersection of Two Arrays
Given two arrays, write a function to compute their intersection.

Example:  
Input: nums1 = [1,2,2,1], nums2 = [2,2]   
Output: [2]  
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]  
Output: [9,4]  

Note:  
Each element in the result must be unique.
The result can be in any order.
```go
func intersection(nums1 []int, nums2 []int) []int {
    m := make(map[int]int)
    for _, n := range nums1 {
        m[n] = 1
    }
    ans := []int{}
    for _, n := range nums2 {
        if _, ok := m[n]; ok {
            ans = append(ans, n)
            delete(m, n)
        }
    }
    return ans
}
```

##### 350.Intersection of Two Arrays II
Given two arrays, write a function to compute their intersection.

Example:  
Input: nums1 = [1,2,2,1], nums2 = [2,2]  
Output: [2,2]  

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]  
Output: [4,9]  

Note:  
Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.

Follow up:  
What if the given array is already sorted? How would you optimize your algorithm?  
What if nums1's size is small compared to nums2's size? Which algorithm is better?  
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?  
```go
func intersect(nums1 []int, nums2 []int) []int {
    m := make(map[int]int)
    ans := []int{}
    for _, n := range nums1 {
        m[n] += 1
    }
    for _, n := range nums2 {
        if _, ok := m[n]; ok {
            ans = append(ans, n)
            m[n] -= 1
            if m[n] == 0 {
                delete(m, n)
            }
        }
    }
    return ans
}
```

##### 697. Degree of an Array
Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements. Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

Example:  
Input: [1, 2, 2, 3, 1]  
Output: 2  
Explanation:   
The input array has a degree of 2 because both elements 1 and 2 appear twice.
Of the subarrays that have the same degree:  
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]   
The shortest length is 2. So return 2.  
Input: [1,2,2,3,1,4,2]  
Output: 6  

Note:
nums.length will be between 1 and 50,000.
nums[i] will be an integer between 0 and 49,999.
```go
func findShortestSubArray(nums []int) int {
    //双指针 + Map
    // find degree lists
    m := make(map[int]int)
    for _, n := range nums {
        m[n] += 1
    }
    d := 0
    for _, v := range m {
        if d <= v {
            d = v
        }
    }
    ds := []int{}
    for k, v := range m {
        if d == v {
            ds = append(ds, k)
        }
    }
    // go over degree lists
    dis := []int{}
    for _, degree := range ds {
        l := 0
        r := len(nums) - 1
        for i, n := range nums {
            if n == degree {
                l = i
                break
            }
        }
        for j := r; j > 0; j-- {
            if nums[j] == degree {
                r = j
                break
            }
        }
        dis = append(dis, r-l+1)
    }
    ans := 50001
    for _, n := range dis {
        if n <= ans {
            ans = n
        }
    }
    return ans
}
```

---

### Stack

##### 921. Minimum Add to Make Parentheses Valid
Given a string S of '(' and ')' parentheses, we add the minimum number of parentheses ( '(' or ')', and in any positions ) so that the resulting parentheses string is valid.

Formally, a parentheses string is valid if and only if:

It is the empty string, or
It can be written as AB (A concatenated with B), where A and B are valid strings, or
It can be written as (A), where A is a valid string.
Given a parentheses string, return the minimum number of parentheses we must add to make the resulting string valid.

Example:

    Input: "())"
    Output: 1

    Input: "((("
    Output: 3

    Input: "()"
    Output: 0

    Input: "()))(("
    Output: 4

Note:

    S.length <= 1000
    S only consists of '(' and ')' characters.

```go
// Double counter, simulate stack, O(1) space, O(N) time.
func minAddToMakeValid(S string) int {
    c1 := 0
    c2 := 0
    for _, c := range S {
        if c == '(' {
            c1 += 1
        } else {
            if c1 == 0 {
                c2 += 1
            } else if c1 > 0 {
                c1 -= 1
            }
        }
    }
    return c1 + c2
}
```

##### 682. Baseball Game
You're now a baseball game point recorder. Given a list of strings, each string can be one of the 4 following types: 

Integer (one round's score): Directly represents the number of points you get in this round.  
"+" (one round's score): Represents that the points you get in this round are the sum of the last two valid round's points.  
"D" (one round's score): Represents that the points you get in this round are the doubled data of the last valid round's points.  
"C" (an operation, which isn't a round's score): Represents the last valid round's points you get were invalid and should be removed.  
Each round's operation is permanent and could have an impact on the round before and the round after.

You need to return the sum of the points you could get in all the rounds.

Example 1:  
Input: ["5","2","C","D","+"]  
Output: 30   
Explanation:    
Round 1: You could get 5 points. The sum is: 5.    
Round 2: You could get 2 points. The sum is: 7.   
Operation 1: The round 2's data was invalid. The sum is: 5.     
Round 3: You could get 10 points (the round 2's data has been removed). The sum is: 15.    
Round 4: You could get 5 + 10 = 15 points. The sum is: 30.    
Example 2:    
Input: ["5","-2","4","C","D","9","+","+"]     
Output: 27   
Explanation:    
Round 1: You could get 5 points. The sum is: 5.   
Round 2: You could get -2 points. The sum is: 3.    
Round 3: You could get 4 points. The sum is: 7.    
Operation 1: The round 3's data is invalid. The sum is: 3.     
Round 4: You could get -4 points (the round 3's data has been removed). The sum is: -1.    
Round 5: You could get 9 points. The sum is: 8.   
Round 6: You could get -4 + 9 = 5 points. The sum is 13.   
Round 7: You could get 9 + 5 = 14 points. The sum is 27.   
Note:    
The size of the input list will be between 1 and 1000.    
Every integer represented in the list will be between -30000 and 30000.    
```go
func calPoints(ops []string) int {
    points := []int{}
    for _, o := range ops {
        l := len(points)
        if o == "C" && l > 0 {
            points = points[:l-1]
        } else if o == "D" && l > 0 {
            points = append(points, points[l-1]*2)
        } else if o == "+" && l > 1 {
            points = append(points, points[l-1]+points[l-2])
        } else {
            n, _ := strconv.Atoi(o)
            points = append(points, n)
        }
    }
    ans := 0
    for _, n := range points {
        ans += n
    }
    return ans
}
```

### Queue

##### 933.Number of Recent Calls
Write a class RecentCounter to count recent requests. It has only one method: ping(int t), where t represents some time in milliseconds. Return the number of pings that have been made from 3000 milliseconds ago until now. Any ping with time in [t - 3000, t] will count, including the current ping. It is guaranteed that every call to ping uses a strictly larger value of t than before.

Example:  
Input: inputs = ["RecentCounter","ping","ping","ping","ping"], inputs = [[],[1],[100],[3001],[3002]]  
Output: [null,1,2,3,3]  
 
Note:  
Each test case will have at most 10000 calls to ping.   
Each test case will call ping with strictly increasing values of t.   
Each call to ping will have 1 <= t <= 10^9.
```go
type RecentCounter struct {
    pings []int
}

func Constructor() RecentCounter {
    rc := RecentCounter{}
    return rc
}

func (this *RecentCounter) Ping(t int) int {
    this.pings = append(this.pings, t)
    l := len(this.pings)
    for i := l - 1; i >= 0; i-- {
        if this.pings[i] < t-3000 {
            this.pings = this.pings[i+1:]
            break
        }
    }
    return len(this.pings)
}
/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
```

##### 950.Reveal Cards In Increasing Order
In a deck of cards, every card has a unique integer.  You can order the deck in any order you want. Initially, all the cards start face down (unrevealed) in one deck. Now, you do the following steps repeatedly, until all cards are revealed:

Take the top card of the deck, reveal it, and take it out of the deck.
If there are still cards in the deck, put the next top card of the deck at the bottom of the deck.
If there are still unrevealed cards, go back to step 1.  Otherwise, stop.
Return an ordering of the deck that would reveal the cards in increasing order.

The first entry in the answer is considered to be the top of the deck.

Example:
Input: [17,13,11,2,3,5,7]  
Output: [2,13,3,11,5,17,7]  
Explanation:   
We get the deck in the order [17,13,11,2,3,5,7] (this order doesn't matter), and reorder it.  
After reordering, the deck starts as [2,13,3,11,5,17,7], where 2 is the top of the deck.  
We reveal 2, and move 13 to the bottom.  The deck is now [3,11,5,17,7,13].  
We reveal 3, and move 11 to the bottom.  The deck is now [5,17,7,13,11].  
We reveal 5, and move 17 to the bottom.  The deck is now [7,13,11,17].  
We reveal 7, and move 13 to the bottom.  The deck is now [11,17,13].  
We reveal 11, and move 17 to the bottom.  The deck is now [13,17].  
We reveal 13, and move 17 to the bottom.  The deck is now [17].  
We reveal 17.  
Since all the cards revealed are in increasing order, the answer is correct.
 
Note:  
1 <= A.length <= 1000
1 <= A[i] <= 10^6
A[i] != A[j] for all i != j
```go
func deckRevealedIncreasing(deck []int) []int {
    sort.Ints(deck)
    l := len(deck)
    q := []int{}
    ans := []int{}
    for i := l - 1; i >= 0; i-- {
        s := len(q)
        if s > 1 {
            m := q[0]
            q = q[1:s]
            q = append(q, m)
        }
        q = append(q, deck[i])
    }
    for i := l - 1; i >= 0; i-- {
        ans = append(ans, q[i])
    }
    return ans
}
```

---

### LinkedList 链表

##### 876.Middle of the Linked List

Given a non-empty, singly linked list with head node head, return a middle node of linked list. If there are two middle nodes, return the second middle node. 

Example:  
Input: [1,2,3,4,5]  
Output: Node 3 from this list (Serialization: [3,4,5])  
The returned node has value 3.  (The judge's serialization of this node is [3,4,5]).  
Note that we returned a ListNode object ans, such that:  
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next = NULL.  
Input: [1,2,3,4,5,6]  
Output: Node 4 from this list (Serialization: [4,5,6])  
Since the list has two middle nodes with values 3 and 4, we return the second one.
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    l := 1
    h := head
    for h.Next != nil {
        h = h.Next
        l++
    }
    m := l / 2
    for i := 0; i < m; i++ {
        head = head.Next
    }
    return head
}
```

---

### Tree 二叉树

##### 617.Merge Two Binary Trees
Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

Example 1:  
Input: 
    Tree 1                     Tree 2                  
          1                         2                             
         / \                       / \                            
        3   2                     1   3                        
       /                           \   \                      
      5                             4   7                  
Output:  
Merged tree:    
         3  
        / \  
       4   5  
      / \   \   
     5   4   7  
 
Note: The merging process must start from the root nodes of both trees.

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil && t2 == nil {
        return nil
    } else if t1 == nil {
        return t2
    } else if t2 == nil {
        return t1
    } else {
        t := &TreeNode{0, nil, nil}
        t.Val = t1.Val + t2.Val
        t.Left = mergeTrees(t1.Left, t2.Left)
        t.Right = mergeTrees(t1.Right, t2.Right)
        return t
    }
    return nil
}
```

##### 965.Univalued Binary Tree
A binary tree is univalued if every node in the tree has the same value. Return true if and only if the given tree is univalued.

Example:
Input: [1,1,1,1,1,null,1]
Output: true
Input: [2,2,2,5,2]
Output: false

Note:
The number of nodes in the given tree will be in the range [1, 100].
Each node's value will be an integer in the range [0, 99].

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    if root == nil {
        return true
    }
    l, r := true, true
    if root.Left != nil {
        l = root.Left.Val == root.Val && isUnivalTree(root.Left)
    }
    if root.Right != nil {
        r = root.Right.Val == root.Val && isUnivalTree(root.Right)
    }
    return l && r
}
```

##### 700.Search in a Binary Search Tree
Given the root node of a binary search tree (BST) and a value. You need to find the node in the BST that the node's value equals the given value. Return the subtree rooted with that node. If such node doesn't exist, you should return NULL.

For example,   
Given the tree:  
        4   
       / \   
      2   7   
     / \  
    1   3    
And the value to search: 2   
You should return this subtree:   
      2      
     / \   
    1   3      
In the example above, if we want to search the value 5, since there is no node with value 5, we should return NULL. Note that an empty tree is represented by NULL, therefore you would see the expected output (serialized tree format) as [], not null.
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == val {
        return root
    } else if root.Val < val {
        return searchBST(root.Right, val)
    } else if root.Val > val {
        return searchBST(root.Left, val)
    }
    return nil
}
```

##### 897.Increasing Order Search Tree
Given a tree, rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree, and every node has no left child and only 1 right child.

Example 1:  
Input: [5,3,6,2,4,null,8,1,null,null,null,7,9]  
       5   
      / \  
    3    6  
   / \    \  
  2   4    8  
 /        / \   
1        7   9  

Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]   
 1  
  \  
   2  
    \  
     3  
      \  
       4  
        \  
         5  
          \  
           6  
            \  
             7   
              \  
               8  
                \  
                 9   

Note:  
The number of nodes in the given tree will be between 1 and 100.
Each node will have a unique integer value from 0 to 1000.
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    values := dfs(root, []int{})
    l := len(values)
    r := &TreeNode{}
    r.Val = values[0]
    for i, node := 1, r; i < l; i++ {
        node.Right = &TreeNode{values[i], nil, nil}
        node = node.Right
    }
    return r
}

func dfs(node *TreeNode, values []int) []int {
    if node == nil {
        return values
    }
    values = dfs(node.Left, values)
    values = append(values, node.Val)
    values = dfs(node.Right, values)
    return values
}
```

##### 872.Leaf-Similar Trees
Consider all the leaves of a binary tree.  From left to right order, the values of those leaves form a leaf value sequence. For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8). Two binary trees are considered leaf-similar if their leaf value sequence is the same. Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

Note: Both of the given trees will have between 1 and 100 nodes.
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    v1 := dfs(root1, []int{})
    v2 := dfs(root2, []int{})
    l1 := len(v1)
    l2 := len(v2)
    if l1 != l2 {
        return false
    }
    for i := 0; i < l1; i++ {
        if v1[i] != v2[i] {
            return false
        }
    }
    return true
}
func dfs(node *TreeNode, values []int) []int {
    if node == nil {
        return values
    }
    if node.Left == nil && node.Right == nil {
        values = append(values, node.Val)
    }
    values = dfs(node.Left, values)
    values = dfs(node.Right, values)
    return values
}
```

##### 669.Trim a Binary Search Tree
Given a binary search tree and the lowest and highest boundaries as L and R, trim the tree so that all its elements lies in [L, R] (R >= L). You might need to change the root of the tree, so the result should return the new root of the trimmed binary search tree.

Example:  
Input:   
    1  
   / \  
  0   2  

  L = 1
  R = 2

Output:   
    1  
      \  
       2  

Input:   
    3  
   / \  
  0   4  
   \  
    2  
   /  
  1  

  L = 1
  R = 3

Output:   
      3  
     /   
   2     
  /  
 1  
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, L int, R int) *TreeNode {
    if root == nil {
        return root
    }
    if root.Val < L {
        return trimBST(root.Right, L, R)
    } else if root.Val > R {
        return trimBST(root.Left, L, R)
    }
    root.Left = trimBST(root.Left, L, R)
    root.Right = trimBST(root.Right, L, R)
    return root
}
```

##### 637.Average of Levels in Binary Tree
Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array. 

Example:  
Input:  
    3  
   / \  
  9  20  
    /  \  
   15   7  
Output: [3, 14.5, 11]   
Explanation:   
The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].

Note:  
The range of node's value is in the range of 32-bit signed integer.
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    if root == nil {
        return []float64{}
    }
    ans := []float64{}
    q := []*TreeNode{}
    q = append(q, root)
    for len(q) > 0 {
        one := 0.0
        l := len(q)
        for _, n := range q {
            one += float64(n.Val)
            if n.Left != nil {
                q = append(q, n.Left)
            }
            if n.Right != nil {
                q = append(q, n.Right)
            }
        }
        q = q[l:]
        ans = append(ans, one/float64(l))
    }
    return ans
}
```

##### 653. Two Sum IV - Input is a BST
Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

Example 1:

Input:  
    5  
   / \  
  3   6  
 / \   \  
2   4   7  

Target = 9

Output: True

Example 2:

Input:   
    5  
   / \  
  3   6  
 / \   \  
2   4   7  

Target = 28  

Output: False
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    nums := preOrder(root, []int{})
    l, r := 0, len(nums)-1
    for l < r {
        if nums[l]+nums[r] < k {
            l++
        } else if nums[l]+nums[r] > k {
            r--
        } else {
            return true
        }
    }
    return false
}

func preOrder(root *TreeNode, nums []int) []int {
    if root == nil {
        return nums
    }
    nums = preOrder(root.Left, nums)
    nums = append(nums, root.Val)
    nums = preOrder(root.Right, nums)
    return nums
}
```

##### 38.Convert BST to Greater Tree
Given a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.

Example:

Input: The root of a Binary Search Tree like this:  
              5  
            /   \  
           2     13  

Output: The root of a Greater Tree like this:  
             18  
            /   \  
          20     13  
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    preOrder(root, 0)
    return root
}
func preOrder(root *TreeNode, n int) int {
    if root == nil {
        return n
    }
    n = preOrder(root.Right, n)
    n += root.Val
    root.Val = n
    n = preOrder(root.Left, n)
    return n
}
```

##### 938.Range Sum of BST
Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive). The binary search tree is guaranteed to have unique values.

Example 1:

Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
Output: 32
Example 2:

Input: root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
Output: 23
 
Note:

The number of nodes in the tree is at most 10000.
The final answer is guaranteed to be less than 2^31.
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
    return helper(root, L, R, 0)
}

func helper(root *TreeNode, L, R, n int) int {
    if root == nil {
        return n
    }
    if root.Val < L {
        n = helper(root.Right, L, R, n)
    } else if root.Val > R {
        n = helper(root.Left, L, R, n)
    } else {
        n += root.Val
        n = helper(root.Left, L, R, n)
        n = helper(root.Right, L, R, n)
    }
    return n
}
```

##### 701.Insert into a Binary Search Tree
Given the root node of a binary search tree (BST) and a value to be inserted into the tree, insert the value into the BST. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.

Note that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

For example, 

Given the tree:  
        4  
       / \  
      2   7  
     / \  
    1   3  
And the value to insert: 5  
You can return this binary search tree:  
         4    
       /   \  
      2     7  
     / \   /  
    1   3 5  
```go 
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{val, nil, nil}
    }
    if root.Val < val {
        root.Right = insertIntoBST(root.Right, val)
    } else if root.Val > val {
        root.Left = insertIntoBST(root.Left, val)
    }
    return root
}
```

##### 654.Maximum Binary Tree
Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

The root is the maximum number in the array.
The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
Construct the maximum tree by the given array and output the root node of this tree.

Example 1:  
Input: [3,2,1,6,0,5]  
Output: return the tree root node representing the following tree:  
      6  
    /   \  
   3     5  
    \    /   
     2  0   
       \  
        1  

Note:  
The size of the given array will be in the range [1,1000].
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    l := len(nums)
    if l == 0 {
        return nil
    }
    index := 0
    for i, n := range nums {
        if nums[index] < n { //find the max
            index = i
        }
    }
    root := &TreeNode{nums[index], nil, nil}
    root.Left = constructMaximumBinaryTree(nums[:index])
    root.Right = constructMaximumBinaryTree(nums[index+1:])
    return root
}
```

##### 814. Binary Tree Pruning
We are given the head node root of a binary tree, where additionally every node's value is either a 0 or a 1. Return the same tree where every subtree (of the given tree) not containing a 1 has been removed. (Recall that the subtree of a node X is X, plus every node that is a descendant of X.)

Example 1:  
Input: [1,null,0,0,1]  
Output: [1,null,0,null,1]  

Explanation:  
Only the red nodes satisfy the property "every subtree not containing a 1".
The diagram on the right represents the answer.

Example 2:   
Input: [1,0,1,0,0,0,1]   
Output: [1,null,1,null,1]

Example 3:   
Input: [1,1,0,1,1,0,1,0]  
Output: [1,1,0,1,1,null,1]

Note:  
The binary tree will have at most 100 nodes.
The value of each node will only be 0 or 1.
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    root.Left = pruneTree(root.Left)
    root.Right = pruneTree(root.Right)
    if root.Val == 0 && root.Left == nil && root.Right == nil {
        return nil
    }
    return root
}
```

##### 543.Diameter of Binary Tree
Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

Example:  
Given a binary tree   
          1  
         / \  
        2   3  
       / \       
      4   5      
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note: The length of path between two nodes is represented by the number of edges between them.
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func Depth(node *TreeNode, s *int) int {
    if node == nil {
        return 0
    }
    l := Depth(node.Left, s)
    r := Depth(node.Right, s)
    if *s < l+r {
        *s = l + r
    }
    if l > r {
        return l + 1
    } else {
        return r + 1
    }
}
func diameterOfBinaryTree(root *TreeNode) int {
    s := 0
    Depth(root, &s)
    return s
}

//Old method, slower.
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l := dfs(root.Left, 0)
    r := dfs(root.Right, 0)
    n := l + r
    lh := diameterOfBinaryTree(root.Left)
    rh := diameterOfBinaryTree(root.Right)
    ans := math.Max(float64(n), math.Max(float64(lh), float64(rh)))
    return int(ans)
}
func dfs(root *TreeNode, n int) int {
    if root == nil {
        return n
    }
    l := dfs(root.Left, n+1)
    r := dfs(root.Right, n+1)
    if l > r {
        return l
    } else {
        return r
    }
}
```


### Graph

##### 797.All Paths From Source to Target
Given a directed, acyclic graph of N nodes.  Find all possible paths from node 0 to node N-1, and return them in any order.

The graph is given as follows:  the nodes are 0, 1, ..., graph.length - 1.  graph[i] is a list of all nodes j for which the edge (i, j) exists.

Example:  
Input: [[1,2], [3], [3], []]   
Output: [[0,1,3],[0,2,3]]   
Explanation: The graph looks like this:  
0--->1  
|    |  
v    v  
2--->3  
There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.  

Note:
The number of nodes in the graph will be in the range [2, 15].
You can print different paths in any order, but you should keep the order of nodes inside one path.
```go
func allPathsSourceTarget(graph [][]int) [][]int {
    ans := [][]int{}
    dfs(graph, &ans, []int{}, 0)
    return ans
}

func dfs(graph [][]int, ans *[][]int, path []int, cur int) {
    path = append(path, cur)
    if cur == len(graph)-1 {
        temp := make([]int, len(path))
        copy(temp, path)
        *ans = append(*ans, temp)
        return
    }
    for _, nei := range graph[cur] {
        dfs(graph, ans, path, nei)
    }
}
```