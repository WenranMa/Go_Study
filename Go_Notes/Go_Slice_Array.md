
# Range 相关

## 2. 下面这段代码输出什么？说明原因。--- TBD

```go
package main

import "fmt"

func main() {
    slice := []int{0, 1, 2, 3}
    m := make(map[int]*int)

    for key, val := range slice {
        m[key] = &val
    }

    for k, v := range m {
        fmt.Println(k, "->", *v)
    }
}
```

**答：**

```shell
// 注：key的顺序无法确定
// go 1.22输出为：
0 -> 0
1 -> 1
2 -> 2
3 -> 3

// 这题不对，之前的答案
0 -> 3
1 -> 3
2 -> 3
3 -> 3
```

**解析：**

~~`for range` 循环的时候会创建每个元素的副本，而不是每个元素的引用，所以 `m[key] = &val` 取的都是变量val的地址，所以最后 `map` 中的所有元素的值都是变量 `val` 的地址，因为最后 `val` 被赋值为3，所有输出的都是3。~~


## 91. 下面代码输出什么？

```go
func main() {
    x := []string{"a", "b", "c"}
    for v := range x {
        fmt.Print(v)
    }
}
```

**答：012**

**解析：**

注意区别下面代码段：

```go
func main() {
    x := []string{"a", "b", "c"}
    for _, v := range x {
        fmt.Print(v)     //输出 abc
    }
}
```


## 70. 下面这段代码输出什么？

```go
func main() {
    var a = [5]int{1, 2, 3, 4, 5}
    var r [5]int

    for i, v := range a {
        if i == 0 {
            a[1] = 12
            a[2] = 13
        }
        r[i] = v
    }
    fmt.Println("r = ", r)
    fmt.Println("a = ", a)
}
```

**答：**

```shell
r =  [1 2 3 4 5]
a =  [1 12 13 4 5]
```

**解析：**

range 表达式是副本参与循环，就是说例子中参与循环的是 a 的副本，而不是真正的 a。就这个例子来说，假设 b 是 a 的副本，则 range 循环代码是这样的：

```go
for i, v := range b {
    if i == 0 {
        a[1] = 12
        a[2] = 13
    }
    r[i] = v
}
```

因此无论 a 被如何修改，其副本 b 依旧保持原值，并且参与循环的是 b，因此 v 从 b 中取出的仍旧是 a 的原值，而非修改后的值。

如果想要 r 和 a 一样输出，修复办法：

```go
func main() {
    var a = [5]int{1, 2, 3, 4, 5}
    var r [5]int

    for i, v := range &a {
        if i == 0 {
            a[1] = 12
            a[2] = 13
        }
        r[i] = v
    }
    fmt.Println("r = ", r)
    fmt.Println("a = ", a)
}
```

输出：

```shell
r =  [1 12 13 4 5]
a =  [1 12 13 4 5]
```

修复代码中，使用 *[5]int 作为 range 表达式，其副本依旧是一个指向原数组 a 的指针，因此后续所有循环中均是 &a 指向的原数组亲自参与的，因此 v 能从 &a 指向的原数组中取出 a 修改后的值。

## 67. 下面这段代码能否正常结束？

```go
func main() {
    v := []int{1, 2, 3}
    for i:= range v {
        v = append(v, i)
    }
}
```

**答：不会出现死循环，能正常结束**

**解析：**

循环次数在循环开始前就已经确定，循环内改变切片的长度，不影响循环次数。


## 169. 下面代码输出什么？ --- TBD

```go
func test() []func() {
    var funs []func()
    for i := 0; i < 2; i++ {
        funs = append(funs, func() {
            println(&i, i)
        })
    }
    return funs
}

func main() {
    funs := test()
    for _, f := range funs {
        f()
    }
}
```

**答：**

```shell
0xc000012028 0
0xc000012040 1


10xc000018058 2
20xc000018058 2
```

**解析：**

~~知识点：闭包延迟求值。for 循环局部变量 i，匿名函数每一次使用的都是同一个变量。（说明：i 的地址，输出可能与上面的不一样）。~~


# 数组，Slice 相关

## 3. 下面两段代码输出什么？

```go
// 1.
func main() {
    s := make([]int, 5)
    s = append(s, 1, 2, 3)
    fmt.Println(s)
}

// 2.
func main() {
    s := make([]int, 0)
    s = append(s, 1, 2, 3, 4)
    fmt.Println(s)
}
```

**答：**

```shell
// 1.
[0 0 0 0 0 1 2 3]

// 2.
[1 2 3 4]
```

**解析：**

使用 `append` 向 `slice` 中添加元素，第一题中slice长度为5，容量为5的零值，第二题为长度为0。


## 7. 下面这段代码能否通过编译，不能的话原因是什么；如果可以，输出什么？

```go
func main() {
    s1 := []int{1, 2, 3}
    s2 := []int{4, 5}
    s1 = append(s1, s2)
    fmt.Println(s1)
}
```

**答：不能通过**

**解析：**

`append()` 的第二个参数不能直接使用 `slice` ，需使用 `...` 操作符，将一个切片追加到另一个切片上： `append(s1, s2...)` 。或者直接跟上元素，形如： `append(s1, 1, 2, 3)`  。



## 12. 以下代码输出什么？

```go
func main() {
    a := []int{7, 8, 9}
    fmt.Printf("%+v\n", a)
    ap(a)
    fmt.Printf("%+v\n", a)
    app(a)
    fmt.Printf("%+v\n", a)
}

func ap(a []int) {
    a = append(a, 10)
}

func app(a []int) {
    a[0] = 1
}
```

**答：输出内容为：**

```shell
[7 8 9]
[7 8 9]
[1 8 9]
```

**解析：**

因为append导致slice的capacity发生变化，所以底层数组重新分配内存了，append中的a这个alice的底层数组和外面不是一个，并没有改变外面的。

## 21. 下面这段代码输出什么？

```go
func hello(num ...int) {
    num[0] = 18
}

func main() {
    i := []int{5, 6, 7}
    hello(i...)
    fmt.Println(i[0])
}
```
- A. 18
- B. 5
- C. Compilation error

**答：A**

**解析：**

可变参数传递过去，改变了第一个值。


## 23. 下面这段代码输出什么？

```go
package main

import (  
    "fmt"
)

func main() {  
    a := [5]int{1, 2, 3, 4, 5}
    t := a[3:4:4]
    fmt.Println(t[0])
}
```

- A. 3
- B. 4
- C. compilation error  

**答：B**

**解析：**

- 知识点：操作符 `[i, j]`。基于数组（切片）可以使用操作符 `[i, j]`创建新的切片，从索引 `i` ，到索引 `j` 结束，截取已有数组（切片）的任意部分，返回新的切片，新切片的值包含原数组（切片）的 `i` 索引的值，但是不包含 `j` 索引的值。`i` 、`j` 都是可选的，`i` 如果省略，默认是0，`j` 如果省略，默认是原数组（切片）的长度。`i` 、`j` 都不能超过这个长度值。

- 假如底层数组的大小为 k，截取之后获得的切片的长度和容量的计算方法：**长度：j-i，容量：k-i**。

- 截取操作符还可以有第三个参数，形如 [i,j,k]，第三个参数 k 用来限制新切片的容量，但 k 不能小于i，j，也不能超过原数组（切片）的底层数组大小。截取获得的切片的长度和容量分别是：**j-i、k-i**。所以例子中，切片 t 为 [4]，长度和容量都是 1。如果`t := a[3:4]`, 则切片长度为1，容量是2。


## 25. 关于 cap() 函数的适用类型，下面说法正确的是()

- A. array
- B. slice
- C. map
- D. channel

**答：A、B、D**

**解析：**

cap() 函数不适用 map


## 24. 下面这段代码输出什么？

```go
func main() {
    a := [2]int{5, 6}
    b := [3]int{5, 6}
    if a == b {
        fmt.Println("equal")
    } else {
        fmt.Println("not equal")
    }
}
```

- A. compilation error  
- B. equal  
- C. not equal 

**答：A**

**解析：**

`invalid operation: a == b (mismatched types [2]int and [3]int)`

Go中的数组是值类型，可比较，另外一方面，数组的长度也是数组类型的组成部分，所以 `a` 和 `b` 是不同的类型，是不能比较的，所以编译错误。


## 71. 下面这段代码输出什么？

```go
func change(s ...int) {
    s = append(s,3)
}

func main() {
    slice := make([]int,5,5)
    slice[0] = 1
    slice[1] = 2
    change(slice...)
    fmt.Println(slice)
    change(slice[0:2]...)
    fmt.Println(slice)
}
```

**答：**

```shell
[1 2 0 0 0]
[1 2 3 0 0]
```

**解析：**

知识点：可变函数、append()操作。

Go 提供的语法糖`...`，可以将 slice 传进可变函数，不会创建新的切片。第一次调用 change() 时，append() 操作使切片底层数组发生了扩容，change函数内部s的地址发生了变化，原 slice 的底层数组不会改变；第二次调用change() 函数时，使用了操作符`[i,j]`获得一个新的切片，假定为 slice1，它的底层数组和原切片底层数组是重合的，不过 slice1 的长度、容量分别是 2、5，所以在 change() 函数中对 slice1 底层数组的修改会影响到原切片。

```go
// 这个代码帮助理解
func change(s ...int) {
	fmt.Println("IN: ", s, &s, &s[0])
	s = append(s, 3)
	fmt.Println("IN: ", s, &s, &s[0])
}

func main() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2

	fmt.Println(slice, len(slice), cap(slice))
	change(slice...)
	fmt.Println(slice, len(slice), cap(slice))
	change(slice[0:2]...)
	fmt.Println(slice, len(slice), cap(slice))
}

/* 结果：
[1 2 0 0 0] 5 5
IN:  [1 2 0 0 0] &[1 2 0 0 0] 0xc00011a000
IN:  [1 2 0 0 0 3] &[1 2 0 0 0 3] 0xc000124000
[1 2 0 0 0] 5 5
IN:  [1 2] &[1 2] 0xc00011a000
IN:  [1 2 3] &[1 2 3] 0xc00011a000
[1 2 3 0 0] 5 5
*/
```

![slice](../file/img/slice_01.webp)

## 135. 下面的代码有什么问题？

```go
package main

import "fmt"

func main() {
    s := make([]int, 3, 9)
    fmt.Println(len(s)) 
    s2 := s[4:8]
    fmt.Println(len(s2)) 
}
```

**答：代码没问题，输出 3 4**

**解析：**

**从一个基础切片派生出的子切片的长度可能大于基础切片的长度**。假设基础切片是 baseSlice，使用操作符 [low,high]，有如下规则：0 <= low <= high <= cap(baseSlice)，只要上述满足这个关系，下标 low 和 high 都可以大于 len(baseSlice)。


## 192. 下面代码输出什么？

```go
func main() {
    a := [3]int{0, 1, 2}
    s := a[1:2]

    s[0] = 11
    s = append(s, 12)
    s = append(s, 13)
    s[0] = 21

    fmt.Println(a)
    fmt.Println(s)
}
```

**答：**

```go
[0 11 12]
[21 12 13]
```


## 194. 下面代码输出什么？

```go
func main() {
    var src, dst []int
    src = []int{1, 2, 3}
    copy(dst, src) 
    fmt.Println(dst)
}
```

**答：输出结果为[]**

**解析：**

copy函数实际上会返回一个int值，这个int是一个size，计算逻辑为size = min(len(dst), len(src))，这个size的大小，
决定了src要copy几个元素给dst，由于题目中，dst声明了，但是没有进行初始化，所以dst的len是0，因此实际没有从src上copy到任何元素给dst





## 155. 下面的代码输出什么？

```go
type T struct {
    n int
}

func main() {
    ts := [2]T{}
    for i, t := range ts {
        switch i {
        case 0:
            t.n = 3
            ts[1].n = 9
        case 1:
            fmt.Print(t.n, " ")
        }
    }
    fmt.Print(ts)
}
```

**答：0 [{0} {9}]**

**解析：**

知识点：for-range 循环数组。

此时使用的是数组 ts 的副本，所以 t.n = 3 的赋值操作不会影响原数组。

## 156. 下面的代码输出什么？

```go
type T struct {
    n int
}

func main() {
    ts := [2]T{}
    for i, t := range &ts {
        switch i {
        case 0:
            t.n = 3
            ts[1].n = 9
        case 1:
            fmt.Print(t.n, " ")
        }
    }
    fmt.Print(ts)
}
```

**答：9 [{0} {9}]**

**解析：**

知识点：for-range 数组指针。

for-range 循环中的循环变量 t 是原数组元素的副本。如果数组元素是结构体值，则副本的字段和原数组字段是两个不同的值。

## 157. 下面的代码输出什么？

```go
type T struct {
    n int
}

func main() {
    ts := [2]T{}
    for i := range ts[:] {
        switch i {
        case 0:
            ts[1].n = 9
        case 1:
            fmt.Print(ts[i].n, " ")
        }
    }
    fmt.Print(ts)
}
```

**答：9 [{0} {9}]**

**解析：**

知识点：for-range 切片。

for-range 切片时使用的是切片的副本，但不会复制底层数组，换句话说，此副本切片与原数组共享底层数组。



## 158. 下面的代码输出什么？

```go
type T struct {
    n int
}

func main() {
    ts := [2]T{}
    for i := range ts[:] {
        switch t := &ts[i]; i {
        case 0:
            t.n = 3;
            ts[1].n = 9
        case 1:
            fmt.Print(t.n, " ")
        }
    }
    fmt.Print(ts)
}
```

**答：9 [{3} {9}]**


## 104. 下面代码输出什么？

```go
func main() {
    var a = []int{1, 2, 3, 4, 5}
    var r = make([]int, 0)

    for i, v := range a {
        if i == 0 {
            a = append(a, 6, 7)
        }

        r = append(r, v)
    }

    fmt.Println(r)
}
```

**答：[1 2 3 4 5]**

**解析：**

a 在 for range 过程中增加了两个元素，len 由 5 增加到 7，但 for range 时会使用 a 的副本 a' 参与循环，副本的 len 依旧是 5，因此 for range 只会循环 5 次，也就只获取 a 对应的底层数组的前 5 个元素。


## 100. 下面这段代码输出什么？

```go
var x = []int{2: 2, 3, 0: 1}

func main() {
    fmt.Println(x)
}
```

**答：[1 0 2 3]**

**解析：**

字面量初始化切片时候，可以指定索引，没有指定索引的元素会在前一个索引基础之上加一，所以输出`[1 0 2 3]`，而不是`[1 3 2]`。


## 37. 下面代码下划线处可以填入哪个选项以输出yes nil？

```go
func main() {
    var s1 []int
    var s2 = []int{}
    if __ == nil {
        fmt.Println("yes nil")
    }else{
        fmt.Println("no nil")
    }
    fmt.Println(s1, s2) // for bypass declare but not use.
}
```

- A. s1
- B. s2
- C. s1、s2 都可以

**答：A**

**解析：**

知识点：nil 切片和空切片。

nil 切片和 nil 相等，一般用来表示一个不存在的切片；空切片和 nil 不相等，表示一个空的集合。


## 40. 切片 a、b、c 的长度和容量分别是多少？

```go
func main() {

    s := [3]int{1, 2, 3}
    a := s[:0]
    b := s[:2]
    c := s[1:2:cap(s)]
}
```

**答：0  3、2  3、1  2**

**解析：**

知识点：数组或切片的截取操作。

结合第23题。截取操作有带 2 个或者 3 个参数，形如：[i:j] 和 [i:j:k]，假设截取对象的底层数组长度为 l。在操作符 [i:j] 中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的长度，截取得到的**切片长度和容量计算方法是 j-i、l-i**。操作符 [i:j:k]，k 主要是用来限制切片的容量，但是不能大于数组的长度 l，截取得到的**切片长度和容量计算方法是 j-i、k-i**。


## 50. 下面的两个切片声明中有什么区别？哪个更可取？

```go
A. var a []int
B. a := []int{}
```

**解析：**

A 声明的是 nil 切片；B 声明的是长度和容量都为 0 的空切片。第一种切片声明不会分配内存，优先选择。


## 55. 下面这段代码输出什么？为什么？--- TBD

```go
func main() {

    s1 := []int{1, 2, 3}
    s2 := s1[1:]
    s2[1] = 4
    fmt.Println(s1)
    s2 = append(s2, 5, 6, 7)
    fmt.Println(s1)
}
```

**答：**

```shell
[1 2 4]
[1 2 4]
```

**解析：**

golang 中切片底层的数据结构是数组。当使用 s1[1:] 获得切片 s2，和 s1 共享同一个底层数组，这会导致 s2[1] = 4 语句影响 s1。

而 append 操作会导致底层数组扩容，生成新的数组，因此追加数据后的 s2 不会影响 s1。

这里注意，**如果分三次append 5,6,7, cap(s2)是8，但一次append(5,6,7), cap(s2)是6。??????????**

![slice](../file/img/slice.webp)


## 65. 下面的代码有什么问题？

```go
func main() {
    fmt.Println([...]int{1} == [2]int{1})
    fmt.Println([]int{1} == []int{1})
}
```

**答：有两处错误**

**解析：**

- go 中不同类型是不能比较的，而数组长度是数组类型的一部分，所以 `[…]int{1}` 和 `[2]int{1}` 是两种不同的类型，不能比较；
- 切片是不能比较的；


## 68. 下面这段代码输出什么？为什么？  --- TBD

```go
func main() {

    var m = [...]int{1, 2, 3}

    for i, v := range m {
        go func() {
            fmt.Println(i, v)
        }()
    }

    time.Sleep(time.Second * 3)
}
```

**答：**

```shell
2 3
2 3
2 3
```

**解析：**

for range 使用短变量声明(:=)的形式迭代变量，需要注意的是，变量 i、v 在每次循环体中都会被重用，而不是重新声明。

各个 goroutine 中输出的 i、v 值都是 for range 循环结束后的 i、v 最终值，而不是各个goroutine启动时的i, v值。可以理解为闭包引用，使用的是上下文环境的值。

两种可行的 fix 方法:

1. **使用函数传递**

   ```go
   for i, v := range m {
       go func(i,v int) {
           fmt.Println(i, v)
       }(i,v)
   }
   ```

2. **使用临时变量保留当前值**

   ```go
   for i, v := range m {
       i := i           // 这里的 := 会重新声明变量，而不是重用
       v := v
       go func() {
           fmt.Println(i, v)
       }()
   }
   ```





## 72. 下面这段代码输出什么？

```go
func main() {
    var a = []int{1, 2, 3, 4, 5}
    var r [5]int

    for i, v := range a {
        if i == 0 {
            a[1] = 12
            a[2] = 13
        }
        r[i] = v
    }
    fmt.Println("r = ", r)
    fmt.Println("a = ", a)
}
```

**答：**

```shell
r =  [1 12 13 4 5]
a =  [1 12 13 4 5]
```

**解析：**

结合70题，70题中a是数组。

此题为切片，切片在 go 的内部结构有一个指向底层数组的指针，当 range 表达式发生复制时，副本的指针依旧指向原底层数组，所以对切片的修改都会反应到底层数组上，所以通过 v 可以获得修改后的数组元素。


## 73. 下面这段代码输出结果正确正确吗？ --- TBD

```go
type Foo struct {
    bar string
}
func main() {
    s1 := []Foo{
        {"A"},
        {"B"},
        {"C"},
    }
    s2 := make([]*Foo, len(s1))
    for i, value := range s1 {
        s2[i] = &value
    }
    fmt.Println(s1[0], s1[1], s1[2])
    fmt.Println(s2[0], s2[1], s2[2])
}
输出：
{A} {B} {C}
&{A} &{B} &{C}
```

**答：结果没有问题 --- ~~s2 的输出结果错误~~**

**解析：**

结合第2题。

~~s2 的输出是 `&{C} &{C} &{C}`，for range 使用短变量声明(:=)的形式迭代变量时，变量 i、value 在每次循环体中都会被重用，而不是重新声明。所以 s2 每次填充的都是临时变量 value 的地址，而在最后一次循环中，value 被赋值为{c}。因此，s2 输出的时候显示出了三个 &{c}。~~

可行的解决办法如下：

```go
for i := range s1 {
    s2[i] = &s1[i]
}
```


## 85. 关于整型切片的初始化，下面正确的是？

- A. s := make([]int)
- B. s := make([]int, 0)
- C. s := make([]int, 5, 10)
- D. s := []int{1, 2, 3, 4, 5}

**答：B C D**

## 140. 下面哪一行代码会 panic，请说明原因？

```go
package main

func main() {
    x := make([]int, 2, 10)
    _ = x[6:10]
    _ = x[6:]
    _ = x[2:]
}
```

**答：第 6 行**

**解析：**

第 6 行，截取符号 [i:j]，如果 j 省略，默认是原切片或者数组的长度，x 的长度是 2，小于起始下标 6 ，所以 panic。

## 142. 下面哪一行代码会 panic，请说明原因？ --- TBD

```go
package main

func main() {
    var m map[int]bool // nil
    _ = m[123]
    var p *[5]string // nil
    for range p {
        _ = len(p)
    }
    var s []int // nil
    _ = s[:]
    s, s[0] = []int{1, 2}, 9
}
```

**答：第 12 行**

**解析：**

因为左侧的 s[0] 中的 s 为 nil。

## 146. 下面代码输出什么？--- TBD
```go
func main() {
    var k = 9
    for k = range []int{} {}
    fmt.Println(k)

    for k = 0; k < 3; k++ {
    }
    fmt.Println(k)

    for k = range (*[3]int)(nil) {
    }
    fmt.Println(k)
}
```

**答：932**
