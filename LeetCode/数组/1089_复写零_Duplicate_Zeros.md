# 1089_复写零_Duplicate_Zeros
Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written.

Do the above modifications to the input array in place, do not return anything from your function.

Example:

    Input: [1,0,2,3,0,4,5,0]
    Output: null
    Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

    Input: [1,2,3]
    Output: null
    Explanation: After calling your function, the input array is modified to: [1,2,3]

Note:

    1 <= arr.length <= 10000
    0 <= arr[i] <= 9

### 解：

遍历数组，如果遇到0，则将后面向后移动一位，包括当前的0，然后i跳过下一个0.

```go
func duplicateZeros(arr []int) {
    l := len(arr)
    for i := 0; i < l; i++ {
        if arr[i] == 0 {
            for j := l - 1; j > i; j-- {
                arr[j] = arr[j-1]
            }
            i += 1
        }
    }
}
```