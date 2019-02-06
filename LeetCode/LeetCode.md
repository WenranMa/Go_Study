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



---

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