# HJ3 明明的随机数

### 较难

明明生成了N个1到500之间的随机整数。请你删去其中重复的数字，即相同的数字只保留一个，把其余相同的数去掉，然后再把这些数从小到大排序，按照排好的顺序输出。

数据范围：1≤n≤1000，

输入的数字大小满足 1≤val≤500 

输入描述：第一行先输入随机整数的个数 N 。 接下来的 N 行每行输入一个整数，代表明明生成的随机数。 具体格式可以参考下面的"示例"。
输出描述：输出多行，表示输入数据处理后的结果

### 示例1
输入：

    3
    2
    2
    1

输出：

    1
    2

### 解：

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	m := make(map[int]struct{})

	var num int
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		m[num] = struct{}{}
	}

	var res []int
	for k, _ := range m {
		res = append(res, k)
	}

	sort.Ints(res)
	for _, a := range res {
		fmt.Println(a)
	}
}
```
