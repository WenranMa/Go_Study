# 739_每日温度_Daily_Temperatures

Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put 0 instead.

For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].

Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].

### 解：

单调栈

```
当 i=0 时，单调栈为空，因此将 0 进栈。
stack=[0(73)]
ans=[0,0,0,0,0,0,0,0]

当 i=1 时，由于 74 大于 73，因此移除栈顶元素 0，赋值 ans[0]:=1−0，将 1 进栈。
stack=[1(74)]
ans=[1,0,0,0,0,0,0,0]

当 i=2 时，由于 75 大于 74，因此移除栈顶元素 1，赋值 ans[1]:=2−1，将 2 进栈。
stack=[2(75)]
ans=[1,1,0,0,0,0,0,0]

当 i=3 时，由于 71 小于 75，因此将 3 进栈。
stack=[2(75),3(71)]
ans=[1,1,0,0,0,0,0,0]

当 i=4 时，由于 69 小于 71，因此将 4 进栈。
stack=[2(75),3(71),4(69)]
ans=[1,1,0,0,0,0,0,0]

当 i=5 时，由于 72 大于 69 和 71，因此依次移除栈顶元素 4 和 3，赋值 ans[4]:=5−4 和 ans[3]:=5−3，将 5 进栈。
stack=[2(75),5(72)]
ans=[1,1,0,2,1,0,0,0]

当 i=6 时，由于 76 大于 72 和 75，因此依次移除栈顶元素 5 和 2，赋值 ans[5]:=6−5 和 ans[2]:=6−2，将 6 进栈。
stack=[6(76)]
ans=[1,1,4,2,1,1,0,0]

当 i=7 时，由于 73 小于 76，因此将 7 进栈。
stack=[6(76),7(73)]
ans=[1,1,4,2,1,1,0,0]
```

```go
func dailyTemperatures(temperatures []int) []int {
	var dstack []int
	var res []int = make([]int, len(temperatures))
	for i, t := range temperatures {
		for len(dstack) > 0 && t > temperatures[dstack[len(dstack)-1]] {
			preIndex := dstack[len(dstack)-1]
			dstack = dstack[:len(dstack)-1]
			res[preIndex] = i - preIndex
		}
		dstack = append(dstack, i)
	}
	return res
}
```

暴力

```go
func dailyTemperatures(T []int) []int {
    res := []int{}
    l := len(T)
    for i := 0; i < l; i++ {
        res = append(res, 0)
    }
    for i, t := range T {
        for j := i; j < l; j++ {
            if T[j] > t {
                res[i] = j - i
                break
            }
        }
    }
    return res
}
```