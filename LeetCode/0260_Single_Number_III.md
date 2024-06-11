# 260. 只出现一次的数字 III

### 中等

给你一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。

你必须设计并实现线性时间复杂度的算法且仅使用常量额外空间来解决此问题。

### 示例 1：

    输入：nums = [1,2,1,3,2,5]
    输出：[3,5]
    解释：[5, 3] 也是有效的答案。

### 示例 2：

    输入：nums = [-1,0]
    输出：[-1,0]

### 示例 3：

    输入：nums = [0,1]
    输出：[1,0]

### 解：
还是XOR的思想，设法将数组分成两个部分，每个部分只有一个只出现一次的数；

从头到尾异或数组，最终得到结果是两个只出现一次的数字的异或结果(resXOR)。因为其他数字都出现过两次，全部异或抵消。由于这两个数字不一样，所以resXOR的二进制中至少有一位是1。找个第一个1的位置(firstBitIndex)，通过判断数组中的每个数在这个位上是不是1，可以把数组分成两部分，而只出现一次的两个数肯定被分开到这两部分中。然后在通过XOR遍历数组即可；O(n) time. O(1) space.

```go
func singleNumber(nums []int) []int {
	res := []int{0, 0}
	if len(nums) < 2 {
		return res
	}
	resXOR := 0
	for _, n := range nums {
		resXOR ^= n
	}
	firstBitIndex := findFirstBit(resXOR)
	for _, n := range nums {
		if isBitOne(n, firstBitIndex) {
			res[0] ^= n
		} else {
			res[1] ^= n
		}
	}
	return res
}

func findFirstBit(n int) int { //find the position of first 1
	indexBit := 0
	for n&1 == 0 && indexBit < 32 {
		n = n >> 1
		indexBit++
	}
	return indexBit
}
func isBitOne(n, indexBit int) bool {
	n = n >> indexBit
	return (n & 1) == 1
}
```