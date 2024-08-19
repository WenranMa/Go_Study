# 1630. 等差子数组 Arithmetic Subarrays

### Medium

A sequence of numbers is called arithmetic if it consists of at least two elements, and the difference between every two consecutive elements is the same. More formally, a sequence s is arithmetic if and only if s[i+1] - s[i] == s[1] - s[0] for all valid i.

For example, these are arithmetic sequences:

1, 3, 5, 7, 9
7, 7, 7, 7
3, -1, -5, -9

The following sequence is not arithmetic:

1, 1, 2, 5, 7
You are given an array of n integers, nums, and two arrays of m integers each, l and r, representing the m range queries, where the ith query is the range [l[i], r[i]]. All the arrays are 0-indexed.

Return a list of boolean elements answer, where answer[i] is true if the subarray nums[l[i]], nums[l[i]+1], ... , nums[r[i]] can be rearranged to form an arithmetic sequence, and false otherwise.

### Example 1:

Input: nums = [4,6,5,9,3,7], l = [0,0,2], r = [2,3,5]
Output: [true,false,true]
Explanation:
In the 0th query, the subarray is [4,6,5]. This can be rearranged as [6,5,4], which is an arithmetic sequence.
In the 1st query, the subarray is [4,6,5,9]. This cannot be rearranged as an arithmetic sequence.
In the 2nd query, the subarray is [5,9,3,7]. This can be rearranged as [3,5,7,9], which is an arithmetic sequence.

### Example 2:

Input: nums = [-12,-9,-3,-12,-6,15,20,-25,-20,-15,-10], l = [0,1,6,4,8,7], r = [4,4,9,7,9,10]
Output: [false,true,false,false,true,true]

Constraints:

n == nums.length
m == l.length
m == r.length
2 <= n <= 500
1 <= m <= 500
0 <= l[i] < r[i] < n
-10^5 <= nums[i] <= 10^5

### 解：

排序后比较，或者用map也可以比较

```go
// O(m * nlogn) time, m is length of l.
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool { 
    s:= len(l)
    res:= []bool{}
    for i:= 0; i<s; i++ {
        t:= []int{}
        t= append(t, nums[l[i]:r[i]+1]...)
        sort.Ints(t)
        st:= len(t)
        df:= 0
        if st >= 2 {
            df= t[1] - t[0]
        }
        r:= true
        for j:= 2; j< st; j++ {
            if t[j] - t[j-1] != df {
                r = false
            }
        }
        res =append(res, r)
    }  
    return res
}

// 改进，用map，求差。
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	s := len(l)
	res := []bool{}
	for i := 0; i < s; i++ {
		st := r[i] - l[i] + 1
		max := -100001
		min := 100001
		m := make(map[int]bool)
		for j := l[i]; j <= r[i]; j++ {
			if nums[j] > max {
				max = nums[j]
			}
			if nums[j] < min {
				min = nums[j]
			}
			m[nums[j]] = true
		}
		if max == min {
			res = append(res, true)
		} else if (max-min)%(r[i]-l[i]) != 0 {
			res = append(res, false)
		} else {
			df := (max - min) / (r[i] - l[i])
			r := true
			for j := 0; j < st; j++ {
				if _, ok := m[min+df*j]; !ok {
					r = false
					break
				}
			}
			res = append(res, r)
		}
	}
	return res
}
//第二个判断要加，[0,0,0,0,0,0,0,0,0,3] 这是个反例。

```