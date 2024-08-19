# 1051_高度检查_Height_Checker

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

### 解：

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