# 1046_最后石头重量_Last_Stone_Weight

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

### 解：

和hw银饰那个题类似。

```go
func lastStoneWeight(stones []int) int {
    for len(stones) > 1 {
        sort.Ints(stones)
        l := len(stones)
        m1 := stones[l-1]
        m2 := stones[l-2]
        stones = stones[0 : l-2]
        if m1 > m2 {
            stones = append(stones, m1-m2)
        }  
    }
    if len(stones) == 1 {
        return stones[0]
    }
    return 0
}
```