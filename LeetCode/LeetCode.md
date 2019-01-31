# LeetCode

### Two Pass (双指针)

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
---

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

### Bit Manipulation（位运算）

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

### Array 数组

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