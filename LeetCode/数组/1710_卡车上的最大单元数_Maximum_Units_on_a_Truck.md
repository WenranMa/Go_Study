# 1710. 卡车上的最大单元数 Maximum Units on a Truck

### Easy

You are assigned to put some amount of boxes onto one truck. You are given a 2D array boxTypes, where boxTypes[i] = [numberOfBoxesi, numberOfUnitsPerBoxi]:

numberOfBoxesi is the number of boxes of type i.
numberOfUnitsPerBoxi is the number of units in each box of the type i.
You are also given an integer truckSize, which is the maximum number of boxes that can be put on the truck. You can choose any boxes to put on the truck as long as the number of boxes does not exceed truckSize.

Return the maximum total number of units that can be put on the truck.

### Example 1:

Input: boxTypes = [[1,3],[2,2],[3,1]], truckSize = 4
Output: 8
Explanation: There are:
- 1 box of the first type that contains 3 units.
- 2 boxes of the second type that contain 2 units each.
- 3 boxes of the third type that contain 1 unit each.
You can take all the boxes of the first and second types, and one box of the third type.
The total number of units will be = (1 * 3) + (2 * 2) + (1 * 1) = 8.

### Example 2:

Input: boxTypes = [[5,10],[2,5],[4,7],[3,9]], truckSize = 10
Output: 91

Constraints:

1 <= boxTypes.length <= 1000
1 <= numberOfBoxesi, numberOfUnitsPerBoxi <= 1000
1 <= truckSize <= 106

### 解：

按 UnitPerbox 排序，然后贪心从前往后遍历。

```go
func maximumUnits(boxTypes [][]int, truckSize int) int {
	res := 0
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	for _, b := range boxTypes {
		if b[0] < truckSize {
			res += b[0] * b[1]
			truckSize -= b[0]
		} else {
			res += truckSize * b[1]
            return res
		}
        //fmt.Println(res, " ", truckSize)
	}
	return res
}
```

```go
// 按type排序，优先用type值更高的箱子。
// 可以用counting sort.
func maximumUnits(boxTypes [][]int, truckSize int) int {
	m := make([]int, 1001)
	res := 0
	for _, b := range boxTypes {
		m[b[1]] += b[0]
	}
	for i := 1000; i >= 0; i-- {
		if m[i] > 0 {
			if truckSize > m[i] {
				res += i * m[i]
				truckSize -= m[i]
			} else {
				res += truckSize * i
				break
			}
		}
	}
	return res
}
```