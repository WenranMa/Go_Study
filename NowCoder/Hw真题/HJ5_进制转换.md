# HJ5 进制转换

### 简单 
写出一个程序，接受一个十六进制的数，输出该数值的十进制表示。

保证结果在 1 ≤ n ≤ 2^31 -1 

输入描述：输入一个十六进制的数值字符串。

输出描述：输出该数值的十进制字符串。不同组的测试用例用\n隔开。

### 示例1

    输入：0xAA
    输出：170

### 解：

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
    var input string
    for {
        n, _ := fmt.Scan(&input)
        if n == 0 {
            break
        } else {
            fmt.Printf("%d\n", hexToDec(input))
        }
    }
}

func hexToDec(str string) int {
    res:= 0
    for i:= 2; i< len(str); i++ {
        res *= 16
        var n int
        if str[i] == 'A' || str[i] == 'B' || str[i] == 'C' || str[i] == 'D' || str[i] == 'E' || str[i] == 'F'{
            n = int(str[i] - 'A' + 10)
        } else {
            n, _ =strconv.Atoi(string(str[i]))
        }
        res += n
    }
    return res
}
```