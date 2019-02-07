# LeetCode

### String

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


---

### Map 

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

---

### Array 数组 and Two Pass 双指针
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


---

### Stack

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




