# 781_森林中的兔子_Rabbits_in_Forest

In a forest, each rabbit has some color. Some subset of rabbits (possibly all of them) tell you how many other rabbits have the same color as them. Those answers are placed in an array.

Return the minimum number of rabbits that could be in the forest.

Examples:

    Input: answers = [1, 1, 2]
    Output: 5
    Explanation:
    The two rabbits that answered "1" could both be the same color, say red.
    The rabbit than answered "2" can't be red or the answers would be inconsistent.
    Say the rabbit that answered "2" was blue.
    Then there should be 2 other blue rabbits in the forest that didn't answer into the array.
    The smallest possible number of rabbits in the forest is therefore 5: 3 that answered plus 2 that didn't.

    Input: answers = [10, 10, 10]
    Output: 11

    Input: answers = []
    Output: 0

### 解：

Map and math.
1. 用map记录每个数字出现次数。
2. 如果数字N出现少于或等于N+1次，则这个颜色的兔子个数为N+1。  比如[10, 10, 10] 这种情况， N = 10， 出现3次，结果就是11。
3. 如果数字N出现大于N+1次，则要算出出现次数对N+1的倍数b，结果就是b*(N+1)。比如[5,5,...5] 13个5的情况，前6个5为一共颜色，中间6个5为一共颜色，还剩1个5，肯定是6只，所以一共6*3 = 18只。
4. 代码中(v-1)与(n+1)都是编程技巧，例如N=3, N+1=4, 则5~8的倍数都是2.

```go
func numRabbits(answers []int) int {
    m := make(map[int]int)
    for _, a := range answers {
        m[a] += 1
    }
    res := 0
    for k, v := range m {
        if v <= k+1 {
            res += (k + 1)
        } else {
            n := (v - 1) / (k + 1)
            res += (k + 1) * (n + 1)
        }
    }
    return res
}
```