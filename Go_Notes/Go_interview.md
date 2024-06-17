# 基础

## 14. 下面这段代码能否编译通过？如果可以，输出什么？-- TBD

```go
const (
    x = iota
    _
    y
    z = "zz"
    k
    p = iota
)

func main() {
    fmt.Println(x, y, z, k, p)
}
```

**答：编译通过，输出：`0 2 zz zz 5`** 

**解析：**

iota初始值为0，所以x为0，_表示不赋值，但是iota是从上往下加1的，所以y是2，z是“zz”,k和上面一个同值也是“zz”,p是iota,从上0开始数他是5。


## 22. 下面这段代码输出什么？

```go
func main() {  
    a := 5
    b := 8.1
    fmt.Println(a + b)
}
```

- A. 13.1  
- B. 13
- C. compilation error  

**答：C**

**解析：**

`invalid operation: a + b (mismatched types int and float64)`

`a` 的类型是`int` ，`b` 的类型是`float` ，两个不同类型的数值不能相加，编译报错。


## 29. 下面这段代码输出什么？

```go
func main() {  
    i := -5
    j := +5
    fmt.Printf("%+d %+d", i, j)
}
```

- A. -5 +5
- B. +5 +5
- C. 0  0

**答：A**

**解析：**

`%d`表示输出十进制数字，`+`表示输出数值的符号。这里不表示取反。

## 61. 下面这段代码输出什么？

```go
const (
    a = iota
    b = iota
)
const (
    name = "name"
    c    = iota
    d    = iota
)
func main() {
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
}
```

**答：0 1 1 2**

**解析：**

知识点：iota 的用法。

iota 是 golang 语言的常量计数器，只能在常量的表达式中使用。

iota 在 const 关键字出现时将被重置为0，const中每新增一行常量声明将使 iota 计数一次。



## 63. 下面这段代码输出什么？

```go
type Direction int

const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) String() string {
    return [...]string{"North", "East", "South", "West"}[d]
}

func main() {
    fmt.Println(South)
}
```

**答：South**

**解析：**

知识点：iota 的用法、类型的 String() 方法。

分别对应0，1，2，3。如果类型实现 String() 方法，当格式化输出时会自动使用 String() 方法。



## 76. 关于循环语句，下面说法正确的有（）

- A. 循环语句既支持 for 关键字，也支持 while 和 do-while；
- B. 关键字 for 的基本使用方法与 C/C++ 中没有任何差异；
- C. for 循环支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break，可以选择中断哪一个循环；
- D. for 循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量；

**答：C D**


`
## 79. 关于switch语句，下面说法正确的有?  --- TBD

- A. 条件表达式必须为常量或者整数；
- B. 单个case中，可以出现多个结果选项；
- C. 需要用break来明确退出一个case；
- D. 只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case；

**答：B D**

## 80. 已知 Add() 函数的调用代码，则Add函数定义正确的是() --- TBD

```go
func main() {
    var a Integer = 1
    var b Integer = 2
    var i interface{} = &a
    sum := i.(*Integer).Add(b)
    fmt.Println(sum)
}
```

```go
A.
type Integer int
func (a Integer) Add(b Integer) Integer {
        return a + b
}

B.
type Integer int
func (a Integer) Add(b *Integer) Integer {
        return a + *b
}

C.
type Integer int
func (a *Integer) Add(b Integer) Integer {
        return *a + b
}

D.
type Integer int
func (a *Integer) Add(b *Integer) Integer {
        return *a + *b
}
```

**答：A C**

**解析：**

知识点：类型断言、方法集。



## 83. 关于GetPodAction定义，下面赋值正确的是

```go
type Fragment interface {
    Exec(transInfo *TransInfo) error
}
type GetPodAction struct {
}
func (g GetPodAction) Exec(transInfo *TransInfo) error {
    ...
    return nil
}
```

- A. var fragment Fragment = new(GetPodAction)
- B. var fragment Fragment = GetPodAction
- C. var fragment Fragment = &GetPodAction{}
- D. var fragment Fragment = GetPodAction{}

**答：A C D**

**解析**

结合第5题。


## 90. 关于异常的触发，下面说法正确的是？

- A. 空指针解析；
- B. 下标越界；
- C. 除数为0；
- D. 调用panic函数；

**答：A B C D**






## 99. 下面代码编译能通过吗？

```go
func main()  
{ 
    fmt.Println("hello world")
}
```

**答：编译错误**

```shell
syntax error: unexpected semicolon or newline before {
```

**解析：**

Go 语言中，大括号不能放在单独的一行。

正确的代码如下：

```go
func main() {
    fmt.Println("works")
}
```


## 102. 请指出下面代码的错误？

```go
package main

var gvar int 

func main() {  
    var one int   
    two := 2      
    var three int 
    three = 3

    func(unused string) {
        fmt.Println("Unused arg. No compile error")
    }("what?")
}
```

**答：变量 one、two 和 three 声明未使用**

**解析：**

知识点：**未使用变量**。

如果有未使用的变量代码将编译失败。但也有例外，函数中声明的变量必须要使用，但可以有未使用的全局变量。函数的参数未使用也是可以的。

如果你给未使用的变量分配了一个新值，代码也还是会编译失败。你需要在某个地方使用这个变量，才能让编译器编译。

修复代码：

```go
func main() {
    var one int
    _ = one

    two := 2
    fmt.Println(two)

    var three int
    three = 3
    one = three

    var four int
    four = four
}
```

另一个选择是注释掉或者移除未使用的变量 。


## 113. 下面代码输出什么？

```go
func test(x byte)  {
    fmt.Println(x)
}

func main() {
    var a byte = 0x11 
    var b uint8 = a
    var c uint8 = a + b
    test(c)
}
```

**答：34**

**解析：**

0x11是16进制，相当于10进制17。

与 rune 是 int32 的别名一样，byte 是 uint8 的别名，别名类型无序转换，可直接转换。

## 114. 下面的代码有什么问题？

```go
func main() {
    const x = 123
    const y = 1.23
    fmt.Println(x)
}
```

**答：编译可以通过**

**解析：**

知识点：常量。

常量是一个简单值的标识符，在程序运行时，不会被修改的量。不像变量，常量未使用是能编译通过的。

## 115. 下面代码输出什么？

```go
const (
    x uint16 = 120
    y
    s = "abc"
    z
)

func main() {
    fmt.Printf("%T %v\n", y, y)
    fmt.Printf("%T %v\n", z, z)
}
```

**答：**

```shell
uint16 120
string abc
```

**解析：**

常量组中如不指定类型和初始化值，则与上一行非空常量右值相同

## 116. 下面代码有什么问题？

```go
func main() {  
    var x string = nil 

    if x == nil { 
        x = "default"
    }
}
```

**答：将 nil 分配给 string 类型的变量**

**解析：**

修复代码：

```go
func main() {  
    var x string //defaults to "" (zero value)

    if x == "" {
        x = "default"
    }
}
```

## 117. 下面的代码有什么问题？

```go
func main() {
    data := []int{1,2,3}
    i := 0
    ++i
    fmt.Println(data[i++])
}
```

**解析：**

对于自增、自减，需要注意：

- 自增、自减不在是运算符，只能作为独立语句，而不是表达式；
- 不像其他语言，Go 语言中不支持 ++i 和 --i 操作；

表达式通常是求值代码，可作为右值或参数使用。而语句表示完成一个任务，比如 if、for 语句等。表达式可作为语句使用，但语句不能当做表达式。

修复代码：

```go
func main() {  
    data := []int{1,2,3}
    i := 0
    i++
    fmt.Println(data[i])
}
```

## 118. 下面代码最后一行输出什么？请说明原因。

```go
func main() {
    x := 1
    fmt.Println(x)
    {
        fmt.Println(x)
        i,x := 2,2
        fmt.Println(i,x)
    }
    fmt.Println(x)  // print ?
}
```

**答：输出`1`**

**解析：**

知识点：变量隐藏。

使用变量简短声明符号 := 时，如果符号左边有多个变量，只需要保证至少有一个变量是新声明的，并对已定义的变量尽进行赋值操作。但如果出现作用域之后，就会导致变量隐藏的问题，就像这个例子一样。

## 119. 下面代码有什么问题？

```go
type foo struct {
    bar int
}

func main() {
    var f foo
    f.bar, tmp := 1, 2
}
```

**答：编译错误**

```shell
non-name f.bar on left side of :=
```

**解析：**

结合110题。

`:=` 操作符不能用于结构体字段赋值。

## 120. 下面的代码输出什么？ --- TBD

```go
func main() {  
    fmt.Println(~2) 
}
```

**答：编译错误**

```shell
cannot use ~ outside of interface or type constraint (use ^ for bitwise complement)
```

**解析：**
位取反运算符，Go 里面采用的是` ^` 。按位取反之后返回一个每个 bit 位都取反的数，对于有符号的整数来说，是按照补码进行取反操作的（快速计算方法：对数 a 取反，结果为 -(a+1) ），对于无符号整数来说就是按位取反。例如：

```go
func main() {
    var a int8 = 3
    var b uint8 = 3
    var c int8 = -3

    fmt.Printf("^%b=%b %d\n", a, ^a, ^a) // ^11=-100 -4
    fmt.Printf("^%b=%b %d\n", b, ^b, ^b) // ^11=11111100 252
    fmt.Printf("^%b=%b %d\n", c, ^c, ^c) // ^-11=10 2
}
```

另外需要注意的是，如果作为二元运算符，^ 表示按位异或，即：对应位相同为 0，相异为 1。例如：

```go
func main() {
    var a int8 = 3
    var c int8 = 5

    fmt.Printf("a: %08b\n",a)
    fmt.Printf("c: %08b\n",c)
    fmt.Printf("a^c: %08b\n",a ^ c)
}
```

给大家重点介绍下这个操作符 &^，按位置零，例如：z = x &^ y，表示如果 y 中的 bit 位为 1，则 z 对应 bit 位为 0，否则 z 对应 bit 位等于 x 中相应的 bit 位的值。

不知道大家发现没有，我们还可以这样理解或操作符 | ，表达式 z = x | y，如果 y 中的 bit 位为 1，则 z 对应 bit 位为 1，否则 z 对应 bit 位等于 x 中相应的 bit 位的值，与 &^ 完全相反。

```go
var x uint8 = 214
var y uint8 = 92
fmt.Printf("x: %08b\n",x)     
fmt.Printf("y: %08b\n",y)       
fmt.Printf("x | y: %08b\n",x | y)     
fmt.Printf("x &^ y: %08b\n",x &^ y)
```

输出：

```shell
x: 11010110
y: 01011100
x | y: 11011110
x &^ y: 10000010
```



## 123. 下面这段代码输出什么？

```GO
type T struct {
    ls []int
}

func foo(t T) {
    t.ls[0] = 100
}

func main() {
    var t = T{
        ls: []int{1, 2, 3},
    }

    foo(t)
    fmt.Println(t.ls[0])
}
```

- A. 1
- B. 100
- C. compilation error

**答：输出 B**

**解析：**

调用 foo() 函数时虽然是传值，但 foo() 函数中，字段 ls 依旧可以看成是指向底层数组的指针。



## 139. 下面的代码输出什么？

```go
type T struct {
    x int
    y *int
}

func main() {

    i := 20
    t := T{10,&i}

    p := &t.x

    *p++
    *p--

    t.y = p

    fmt.Println(*t.y)
}
```

**答：10**

**解析：**

知识点：运算符优先级。

如下规则：递增运算符 ++ 和递减运算符 -- 的优先级低于解引用运算符 * 和取址运算符 &，解引用运算符和取址运算符的优先级低于选择器 . 中的属性选择操作符。




## 148. 下面代码输出什么？

```go
func main() {
    var x int8 = -128
    var y = x/-1
    fmt.Println(y)
}
```

**答：-128**

**解析：**

溢出

## 149. 下面选项正确的是？

- A. 类型可以声明的函数体内；
- B. Go 语言支持 ++i 或者 --i 操作；
- C. nil 是关键字；
- D. 匿名函数可以直接赋值给一个变量或者直接执行；

**答：A D**


## 161. 关于 slice 或 map 操作，下面正确的是？

* A

  ```go
  var s []int
  s = append(s,1)
  ```

* B

  ```go
  var m map[string]int
  m["one"] = 1 
  ```

* C

  ```go
  var s []int
  s = make([]int, 0)
  s = append(s,1)
  ```

* D

  ```go
  var m map[string]int
  m = make(map[string]int)
  m["one"] = 1 
  ```

**答：A C D**

## 162. 下面代码输出什么？

```go
func test(x int) (func(), func()) {
    return func() {
        println(x)
        x += 10
    }, func() {
        println(x)
    }
}

func main() {
    a, b := test(100)
    a()
    b()
}
```

**答：100 110**

**解析：**

知识点：闭包引用相同变量。

## 163. 关于字符串连接，下面语法正确的是？

- A. str := 'abc' + '123'
- B. str := "abc" + "123"
- C. str ：= '123' + "abc"
- D. fmt.Sprintf("abc%d", 123)

**答：B D**

**解析：**

知识点：单引号、双引号和字符串连接。

在 Go 语言中，双引号用来表示字符串 string，其实质是一个 byte 类型的数组，单引号表示 rune 类型。



## 165. 判断题：对变量x的取反操作是 ~x？

**答：错**

**解析：**

Go 语言的取反操作是 `^`，它返回一个每个 bit 位都取反的数。作用类似在 C、C#、Java 语言中中符号 ~，对于有符号的整数来说，是按照补码进行取反操作的（快速计算方法：对数 a 取反，结果为 -(a+1) ），对于无符号整数来说就是按位取反。


----

## 187. 下面这段代码输出什么？

```go
func main() {
    count := 0
    for i := range [256]struct{}{} {
        m, n := byte(i), int8(i)
        if n == -n {
            count++
        }
        if m == -m {
            count++
        }
    }
    fmt.Println(count)
}
```

**解析：**

知识点：数值溢出。当 i 的值为 0、128 是会发生相等情况，注意 byte 是 uint8 的别名。

byte = uint8. 所以取值范围为: [0, 255]. 所以-m为负数就溢出了byte的表示范围. 那么 对 0~255 取反什么时候相等呢? 0是相等的就不用说来. 为啥128与-128相等呢?

我们知道负数是使用补位计数表示的. 所以-128,

先对128取反(^1000 0000 = 0111 1111)
然后进行加1操作, 即0111 1111 + 1 = 1000 0000
所以-128是与正的128(1000 0000)一样.

int8 表示范围是[-128, 127]. 所以当循环中i >= 128时, 超出了int8的表示范围, 就溢出了.

那么当循环 i = 128时. 128 = 1000 0000. n = int8(128) 强制转换后时多少呢?

我们知道补位计数法中, 正整数, 0, 最高为0, 而负数最高为1.
所以1000 0000 一定是一个负数, 我们用补位计数法逆推. 就是上面步骤反向

减1操作: 1000 0000 - 1 = 0111 1111
取反操作: ^(0111 1111) = 1000 0000
int8(128) = 1000 0000 对与int8表示 -128, 而-128 = 1000 0000 对于 int8就是-128.

所以当i= 128时, n = int8(128) = -128, 而-n = 128, 128刚好溢出了int8的最大值. 对于int8 就是-128.

注意:

补位计数法中, 正整数与0, 符号位和值部是分开的. 对int8来说, 符号位(左侧最高位0)表示正数, 剩余7位用来表示正整数. 所以int8的最大值为 0111 1111 = 127.
对于负数, 符号位与值部是在一起的. 左侧最高位1即是表示负数, 也是值的一部分. 如1000 0000 = -128

## 188. 下面代码输出什么？

```go
const (
    azero = iota
    aone  = iota
)

const (
    info  = "msg"
    bzero = iota
    bone  = iota
)

func main() {
    fmt.Println(azero, aone)
    fmt.Println(bzero, bone)
}
```

**答：0 1 1 2**

**解析：**

知识点：iota 的使用。这道题易错点在 bzero、bone 的值，在一个常量声明代码块中，如果 iota 没出现在第一行，则常量的初始值就是非 0 值。

