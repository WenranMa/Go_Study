# 40. 组合总和 II

### 中等

给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

注意：解集不能包含重复的组合。 

### 示例 1:

    输入: candidates = [10,1,2,7,6,1,5], target = 8,
    输出:
    [
    [1,1,6],
    [1,2,5],
    [1,7],
    [2,6]
    ]

### 示例 2:

    输入: candidates = [2,5,2,1,2], target = 5,
    输出:
    [
    [1,2,2],
    [5]
    ]
 
### 提示:
- 1 <= candidates.length <= 100
- 1 <= candidates[i] <= 50
- 1 <= target <= 30

### 解：

`if i > start && candi[i] == candi[i-1]` 这句判断很重要。特变是 `i > start`用于去重 

```go
var res [][]int

func combinationSum2(candidates []int, target int) [][]int {
	res = [][]int{}
	sort.Ints(candidates)
	comb(candidates, 0, target, []int{})
	return res
}

func comb(candi []int, start int, target int, row []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		r := make([]int, len(row))
		copy(r, row)
		res = append(res, r)
	}
	for i := start; i < len(candi); i++ {
		if i > start && candi[i] == candi[i-1] {
			continue
		}
		row = append(row, candi[i])
		comb(candi, i+1, target-candi[i], row)
		row = row[:len(row)-1]
	}
}
```