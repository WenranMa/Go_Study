# golang 处理ACM模式输入输出

## 1. 用 fmt 包实现简单读取数字
### 1.1 每行数字数量固定

#### 1.1.1 知道行数
输入描述: 输入第一行包括一个数据组数t(1 <= t <= 100)
接下来每行包括两个正整数a,b(1 <= a, b <= 1000)

输出描述: 输出a+b的结果

输入例子1:

    2
    1 5
    10 20

输出例子1:

    6
    30

```go
package main

import(
    "fmt"
)

func main(){
    var t,a,b int
    fmt.Scanln(&t)
    for i:=0;i<t;i++{
        // or put a, b here
        fmt.Scanln(&a,&b)
        fmt.Println(a+b)
    }
}
```

#### 1.1.2 不知道行数，没有标志位
输入描述: 输入包括两个正整数 a,b(1 <= a, b <= 1000),输入数据包括多组。

输出描述: 输出a+b的结果

输入例子1:

    1 5
    10 20

输出例子1:

    6
    30

```go
// 代码：使用 fmt.Scanln 或 fmt.Scan

package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scanln(&a, &b)  //也可以用 fmt.Scan
		if n == 0 {
			break
		} else {
			fmt.Printf("%d\n", a+b)
		}
	}
}
```

#### 1.1.3 不知道行数，有标志位
输入例子1:

    1 5
    10 20
    0 0

输出例子1:

    6
    30

```go
package main
 
import (
    "fmt"
)
 
func main() {
    var a, b int
    for {
        fmt.Scan(&a, &b)
        if a == 0 && b == 0 {
            break
        }
        fmt.Printf("%d\n", a+b)
    }
}
```


### 1.2. 每行数字数量不固定，但知道数量
#### 1.2.1 不知道行数，但有结束标志
输入描述: 输入数据包括多组。
每组数据一行,每行的第一个整数为整数的个数n(1 <= n <= 100), n为0的时候结束输入。接下来n个正整数,即需要求和的每个正整数。

输出描述: 每组数据输出求和的结果

输入例子1:

    4 1 2 3 4
    5 1 2 3 4 5
    0

输出例子1:

    10
    15

```go
package main

import (
	"fmt"
)

func main() {
	var t int
	for {
        var sum int
		fmt.Scan(&t)
		if t == 0 {
			break
		} else {
			a := make([]int, t)
			for i := 0; i < t; i++ {
				fmt.Scan(&a[i])
			}
			for i := 0; i < t; i++ {
				sum += a[i]
			}
		}
		fmt.Println(sum)
	}
}
```
或者
```go
package main

import(
    "fmt"
)

func main(){
    var t, crr, sum int
    for{
        fmt.Scan(&t)
        sum = 0
        if t==0{
            break
        }else{
            for i:=0;i<t;i++{
                fmt.Scan(&crr)
                sum += crr
            }
            fmt.Println(sum)
        }
    }
}
```

#### 1.2.2 不知道行数，没有结束标志（没得读了就结束）
输入描述: 输入数据有多组, 每行表示一组输入数据。
每行的第一个整数为整数的个数n(1 <= n <= 100)。
接下来n个正整数, 即需要求和的每个正整数。

输出描述: 每组数据输出求和的结果

输入例子1:

    4 1 2 3 4
    5 1 2 3 4 5

输出例子1:

    10
    15

```go
package main

import (
	"fmt"
)

func main() {
	var t int
	for {
        var sum int
        n,_ := fmt.Scan(&t)
		if n == 0 {
			break
		} else {
			a := make([]int, t)
			for i := 0; i < t; i++ {
				fmt.Scan(&a[i])
			}
			for i := 0; i < t; i++ {
				sum += a[i]
			}
		}
		fmt.Println(sum)
	}
}
```
或者
```go
package main

import(
    "fmt"
)

func main(){
    var t,crr,sum int
    for {
        n,_ := fmt.Scan(&t)
        if n == 0{
            break
        }else{
            sum = 0
            for i:=0;i<t;i++{
                fmt.Scan(&crr)
                sum += crr
            }
            fmt.Println(sum)
        }
    }  
}
```

#### 1.2.3 知道行数
输入描述:
输入的第一行包括一个正整数t(1 <= t <= 100), 表示数据组数。
接下来t行, 每行一组数据。
每行的第一个整数为整数的个数n(1 <= n <= 100)。
接下来n个正整数, 即需要求和的每个正整数。

输出描述:
每组数据输出求和的结果

输入例子1:

    2
    4 1 2 3 4
    5 1 2 3 4 5

输出例子1:

    10
    15

```go
package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var num, sum, a int
		fmt.Scan(&num)
		for i := 0; i < num; i++ {
			fmt.Scan(&a)
			sum += a
		}
		fmt.Println(sum)
	}
}
```

## 2. 必须整行读（用 fmt、os、bufio、strconv、strings 实现）
### 2.1 每行数字数量不固定，且不知道数量

输入描述: 输入数据有多组, 每行表示一组输入数据。
每行不定有n个整数，空格隔开。(1 <= n <= 100)。

输出描述: 每组数据输出求和的结果

输入例子1:

    1 2 3
    4 5
    0 0 0 0 0

输出例子1:

    6
    9
    0

这个时候我们需要一整行一整行地读，这时需要用到 bufio 包：

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
    inputs := bufio.NewScanner(os.Stdin)
    for inputs.Scan() {  //每次读入一行
        data := strings.Split(inputs.Text(), " ")  //通过空格将他们分割，并存入一个字符串切片
        var sum int
        for i := range data {
            val, _ := strconv.Atoi(data[i])   //将字符串转换为int
            sum += val
        }
        fmt.Println(sum)
    }
}
```

## 3. 读字符串
### 3.1. 读一行字符串

输入描述: 输入有两行，第一行n
第二行是n个字符串，字符串之间用空格隔开

输出描述:
输出一行排序后的字符串，空格隔开，无结尾空格

输入例子1:

    5
    c d a bb e

输出例子1:

    a bb c d e

```go
package main
import(
    "fmt"
    "os"
    "bufio"
    "sort"
    "strings"
)
 
func main(){
    in := bufio.NewScanner(os.Stdin)
    in.Scan()
    for in.Scan(){
        str := in.Text()
        s := strings.Split(str, " ")
        sort.Strings(s)  //排序
        fmt.Println(strings.Join(s," "))  //将切片连接成字符串
    }
}
```

### 3.2. 读多行字符串

输入描述: 多个测试用例，每个测试用例一行。
每行通过空格隔开，有n个字符，n＜100

输出描述: 对于每组测试用例，输出一行排序过的字符串，每个字符串通过空格隔开

输入例子1:

    a c bb
    f dddd
    nowcoder

输出例子1:

    a bb c
    dddd f
    nowcoder

```go
package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "sort"
)

func main(){
    input := bufio.NewScanner(os.Stdin)
    for input.Scan(){
        data := strings.Split(input.Text()," ")
        sort.Strings(data)
        fmt.Println(strings.Join(data, " "))
	}
}
```
