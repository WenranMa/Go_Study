# HJ33 整数与IP地址间的转换

### 中等
原理：ip地址的每段可以看成是一个0-255的整数，把每段拆分成一个二进制形式组合起来，然后把这个二进制数转变成
一个长整数。

### 举例：
一个ip地址为10.0.3.193
    
    每段数字             相对应的二进制数
    10                   00001010
    0                    00000000
    3                    00000011
    193                  11000001

组合起来即为：00001010 00000000 00000011 11000001,转换为10进制数就是：167773121，即该IP地址转换后的数字就是它了。

数据范围：保证输入的是合法的 IP 序列

### 输入描述：
输入 
1 输入IP地址
2 输入10进制型的IP地址

### 输出描述：
输出
1 输出转换成10进制的IP地址
2 输出转换后的IP地址

### 示例1
输入：

    10.0.3.193
    167969729

输出：

    167773121
    10.3.3.193

### 解：

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	inputIP := input.Text()
	input.Scan()
	inputInt, _ := strconv.Atoi(input.Text())

	fmt.Println(ipToInt(inputIP))
	fmt.Println(intToIP(inputInt))
}

func ipToInt(ip string) int {
	segments := strings.Split(ip, ".")
	var num int
	for _, seg := range segments {
		num <<= 8
		n, _ := strconv.Atoi(seg)
		num += n
	}
	return num
}

func intToIP(n int) string {
	var segments []int
	for n != 0 {
		segments = append(segments, n&255)
		n >>= 8
	}
	return fmt.Sprintf("%d.%d.%d.%d", segments[3], segments[2], segments[1], segments[0])
}
```