# 697_数组的度_Degree_of_an_Array

Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements. Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

Example:

    Input: [1, 2, 2, 3, 1]  
    Output: 2  
    Explanation:   
    The input array has a degree of 2 because both elements 1 and 2 appear twice.
    Of the subarrays that have the same degree:  
    [1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]   
    The shortest length is 2. So return 2.  
    Input: [1,2,2,3,1,4,2]  
    Output: 6  

Note:
nums.length will be between 1 and 50,000.
nums[i] will be an integer between 0 and 49,999.


### 解：

```go
func findShortestSubArray(nums []int) int {
	//双指针 + Map
	// find degree lists
	m := make(map[int]int)
	for _, n := range nums {
		m[n] += 1
	}
	d := 0
	for _, v := range m {
		if d <= v {
			d = v
		}
	}
	ds := []int{}
	for k, v := range m {
		if d == v {
			ds = append(ds, k)
		}
	}
	// go over degree lists
	var ans int = 50001
	for _, degree := range ds {
		l := 0
		r := len(nums) - 1
		for i, n := range nums {
			if n == degree {
				l = i
				break
			}
		}
		for j := r; j > 0; j-- {
			if nums[j] == degree {
				r = j
				break
			}
		}
		ans = min(ans, r-l+1)
	}
	return ans
}
```

优化，定义一个结构体，一次遍历数组可以存下count, indexL, indexR.

```go
type Degree struct {
    count int
    indexL int
    indexR int
}

func findShortestSubArray(nums []int) int {
    m := make(map[int]*Degree)
    for i, n := range nums {
        if _, ok := m[n]; !ok {
            m[n] = &Degree{1, i, i}
        } else {
            m[n].count += 1
            m[n].indexR = i
        }
    }
    var ans int = 50001
    var maxCnt int = 0
    for _, v := range m {
        if v.count > maxCnt {
            maxCnt = v.count
            ans = v.indexR - v.indexL + 1
        } else if v.count == maxCnt{
            ans = min(ans, v.indexR - v.indexL + 1)
        }
    }
    return ans
}
```