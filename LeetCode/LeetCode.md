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