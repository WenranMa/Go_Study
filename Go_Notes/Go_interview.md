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


## 66. 下面这段代码输出什么？如果编译错误的话，为什么？

```go
var p *int

func foo() (*int, error) {
    var i int = 5
    return &i, nil
}

func bar() {
    //use p
    fmt.Println(*p)
}

func main() {
    p, err := foo()
    if err != nil {
        fmt.Println(err)
        return
    }
    bar()
    fmt.Println(*p)
}
```

- A. 5 5
- B. runtime error

**答：B**

**解析：**

知识点：变量作用域。

问题出在操作符`:=`，对于使用`:=`定义的变量，如果新变量与同名已定义的变量不在同一个作用域中，那么 Go 会新定义这个变量。对于本例来说，main() 函数里的 p 是新定义的变量，会遮住全局变量 p，导致执行到`bar()`时程序，全局变量 p 依然还是 nil，程序随即 Crash。

正确的做法是将 main() 函数修改为：

```go
func main() {
    var err error
    p, err = foo()
    if err != nil {
        fmt.Println(err)
        return
    }
    bar()
    fmt.Println(*p)
}
```

## 74. 下面代码里的 counter 的输出值？--- TBD

```go
func main() {
    var m = map[string]int{
        "A": 21,
        "B": 22,
        "C": 23,
    }
    counter := 0
    for k, v := range m {
        if counter == 0 {
            delete(m, "A")
        }
        counter++
        fmt.Println(k, v)
    }
    fmt.Println("counter is ", counter)
}
```

- A. 2
- B. 3
- C. 2 或 3

**答：C**

**解析：**

for range map 是无序的，如果第一次循环到 A，则输出 3；否则输出 2。



## 76. 关于循环语句，下面说法正确的有（）

- A. 循环语句既支持 for 关键字，也支持 while 和 do-while；
- B. 关键字 for 的基本使用方法与 C/C++ 中没有任何差异；
- C. for 循环支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break，可以选择中断哪一个循环；
- D. for 循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量；

**答：C D**

## 77. 下面代码输出正确的是？

```go
func main() {
    i := 1
    s := []string{"A", "B", "C"}
    i, s[i-1] = 2, "Z"
    fmt.Printf("s: %v \n", s)
}
```

- A. s: [Z,B,C]
- B. s: [A,Z,C]

**答：A**

**解析：**

知识点：多重赋值。

多重赋值分为两个步骤，有先后顺序：

- 计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；
- 赋值；

所以本例，会先计算 s[i-1]，等号右边是两个表达式是常量，所以赋值运算等同于 `i, s[0] = 2, "Z"`。

## 78. 关于类型转化，下面选项正确的是？-- TBD

```go
A.
type MyInt int
var i int = 1
var j MyInt = i

B.
type MyInt int
var i int = 1
var j MyInt = (MyInt)i

C.
type MyInt int
var i int = 1
var j MyInt = MyInt(i)

D.
type MyInt int
var i int = 1
var j MyInt = i.(MyInt)
```

**答：C**

**解析：**

知识点：强制类型转化

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

## 81. 关于 bool 变量 b 的赋值，下面错误的用法是？

- A. b = true
- B. b = 1
- C. b = bool(1)
- D. b = (1 == 2)

**答：B C**

**解析**

B: `cannot use 1 (untyped int constant) as bool value in assignment`

C: `cannot convert 1 (untyped int constant) to type bool`

## 82. 关于变量的自增和自减操作，下面语句正确的是？

```go
A.
i := 1
i++

B.
i := 1
j = i++

C.
i := 1
++i

D.
i := 1
i--
```

**答：A D**

**解析：**

知识点：自增自减操作。

i++ 和 i-- 在 Go 语言中是语句，不是表达式，因此不能赋值给另外的变量。此外没有 ++i 和 --i。

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

## 84. 关于函数声明，下面语法正确的是？

- A. func f(a, b int) (value int, err error)
- B. func f(a int, b int) (value int, err error)
- C. func f(a, b int) (value int, error)
- D. func f(a int, b int) (int, int, error)

**答：A B D**

**解析**

结合第4题。


## 88. 下面代码有什么问题？

```go
const i = 100
var j = 123

func main() {
    fmt.Println(&j, j)
    fmt.Println(&i, i)
}
```

**答：编译报错**

**解析：**

编译报错`cannot take the address of i`。知识点：常量。常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，所以常量无法寻址。

## 89. 下面代码能否编译通过？如果通过，输出什么？

```go
func GetValue(m map[int]string, id int) (string, bool) {

    if _, exist := m[id]; exist {
        return "exist", true
    }
    return nil, false
}
func main() {
    intmap := map[int]string{
        1: "a",
        2: "b",
        3: "c",
    }

    v, err := GetValue(intmap, 3)
    fmt.Println(v, err)
}
```

**答：不能通过编译**

**解析：**

知识点：函数返回值类型。nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。但是如果不特别指定的话，Go 语言不能识别类型，所以会报错:`cannot use nil as type string in return argument`

## 90. 关于异常的触发，下面说法正确的是？

- A. 空指针解析；
- B. 下标越界；
- C. 除数为0；
- D. 调用panic函数；

**答：A B C D**



## 92. 下面这段代码能否编译通过？如果通过，输出什么？

```go
type User struct{}
type User1 User
type User2 = User

func (i User) m1() {
    fmt.Println("m1")
}
func (i User) m2() {
    fmt.Println("m2")
}

func main() {
    var i1 User1
    var i2 User2
    i1.m1()
    i2.m2()
}
```

**答：不能，报错`i1.m1 undefined (type User1 has no field or method m1)`**

**解析：**

第 2 行代码基于类型 User 创建了新类型 User1，第 3 行代码是创建了 User 的类型别名 User2，注意使用 = 定义类型别名。因为 User2 是别名，完全等价于 User，所以 User2 具有 User 所有的方法。但是 i1.m1() 是不能执行的，因为 User1 没有定义该方法。




## 98. 下面这段代码存在什么问题？

```go
type Param map[string]interface{}
 
type Show struct {
    *Param
}

func main() {
    s := new(Show)
    s.Param["day"] = 2
}
```

**答：存在两个问题**

**解析：**

1. map 需要初始化才能使用；

2. 指针不支持索引。修复代码如下：

   ```go
   func main() {
       s := new(Show)
       // 修复代码
       p := make(Param)
       p["day"] = 2
       s.Param = &p
       tmp := *s.Param //指针不能索引
       fmt.Println(tmp["day"])
   }
   ```

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




## 105. 下面的代码有什么问题？

```go
import (  
    "fmt"
    "log"
    "time"
)
func main() {  
}
```

**答：导入的包没有被使用**

**解析：**

如果引入一个包，但是未使用其中如何函数、接口、结构体或变量的话，代码将编译失败。

如果你真的需要引入包，可以使用下划线操作符，`_`，来作为这个包的名字，从而避免失败。下划线操作符用于引入，但不使用。

我们还可以注释或者移除未使用的包。

修复代码：

```go
import (  
    _ "fmt"
    "log"
    "time"
)
var _ = log.Println
func main() {  
    _ = time.Now
}
```



## 107. 下面代码有几处错误的地方？请说明原因。

```go
func main() {
    var s []int
    s = append(s,1)

    var m map[string]int
    m["one"] = 1 
}
```

**答：有 1 处错误**

**解析：**

有 1 处错误，不能对 nil 的 map 直接赋值，需要使用 make() 初始化。但可以使用 append() 函数对为 nil 的 slice 增加元素。

修复代码：

```go
func main() {
    var m map[string]int
    m = make(map[string]int)
    m["one"] = 1
}
```

## 108. 下面代码有什么问题？

```go
func main() {
    m := make(map[string]int,2)
    cap(m) 
}
```

**答：使用 cap() 获取 map 的容量**

**解析：**

1. 使用 make 创建 map 变量时可以指定第二个参数，不过会被忽略。

2. cap() 函数适用于数组、数组指针、slice 和 channel，不适用于 map，可以使用 len() 返回 map 的元素个数。



## 110. 下面代码能编译通过吗？

```go
type info struct {
    result int
}

func work() (int,error) {
    return 13,nil
}

func main() {
    var data info

    data.result, err := work() 
    fmt.Printf("info: %+v\n",data)
}
```

**答：编译失败**

```shell
non-name data.result on left side of :=
```

**解析：**

不能使用短变量声明设置结构体字段值，修复代码：

```go
func main() {
    var data info

    var err error
    data.result, err = work() //ok
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(data)   
}
```

## 111. 下面代码有什么错误？

```go
func main() {
    one := 0
    one := 1 
}
```

**答：变量重复声明**

**解析：**

不能在单独的声明中重复声明一个变量，但在多变量声明的时候是可以的，但必须保证至少有一个变量是新声明的。

修复代码：

```go
func main() {  
    one := 0
    one, two := 1,2
    one,two = two,one
}
```

## 112. 下面代码有什么问题？

```go
func main() {
    x := []int{
        1,
        2
    }
    _ = x
}
```

**答：编译错误**

**解析：**

第四行代码没有逗号。用字面量初始化数组、slice 和 map 时，最好是在每个元素后面加上逗号，即使是声明在一行或者多行都不会出错。

修复代码：

```go
func main() {
	x := []int{ // 多行
		1,
		2,
	}
	_ = x

	y := []int{3, 4,} // 一行 no error
	_ = y
}
```

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



## 122. 下面这段代码输出什么？

```go
type People struct {
    name string `json:"name"`
}

func main() {
    js := `{
        "name":"seekload"
    }`
    var p People
    err := json.Unmarshal([]byte(js), &p)
    if err != nil {
        fmt.Println("err: ", err)
        return
    }
    fmt.Println(p)
}
```

**答：输出 {}**

**解析：**

知识点：结构体访问控制，因为 name 首字母是小写，导致其他包不能访问，所以输出为空结构体。

修复代码：

```go
type People struct {
    Name string `json:"name"`
}
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

## 124. 下面代码输出什么？

```go
func main() {
    isMatch := func(i int) bool {
        switch(i) {
        case 1:
        case 2:
            return true
        }
        return false
    }

    fmt.Println(isMatch(1))
    fmt.Println(isMatch(2))
}
```

**答：false true**

**解析：**

Go 语言的 switch 语句虽然没有"break"，但如果 case 完成程序会默认 break，可以在 case 语句后面加上关键字 fallthrough，这样就会接着走下一个 case 语句（不用匹配后续条件表达式）。或者，利用 case 可以匹配多个值的特性。

修复代码：

```go
func main() {
    isMatch := func(i int) bool {
        switch(i) {
        case 1:
            fallthrough
        case 2:
            return true
        }
        return false
    }

    fmt.Println(isMatch(1))     // true
    fmt.Println(isMatch(2))     // true

    match := func(i int) bool {
        switch(i) {
        case 1,2:
            return true
        }
        return false
    }

    fmt.Println(match(1))       // true
    fmt.Println(match(2))       // true
}
```

## 125. 下面的代码能否正确输出？

```go
func main() {
    var fn1 = func() {}
    var fn2 = func() {}

    if fn1 != fn2 {
        println("fn1 not equal fn2")
    }
}
```

**答：编译错误**

```shell
invalid operation: fn1 != fn2 (func can only be compared to nil)
```

**解析：**

函数只能与 nil 比较。

## 126. 下面代码输出什么？

```go
type T struct {
    n int
}

func main() {
    m := make(map[int]T)
    m[0].n = 1
    fmt.Println(m[0].n)
}
```

- A. 1
- B. compilation error

**答：B**

```shell
cannot assign to struct field m[0].n in map
```

**解析：**

map[key]struct 中 struct 是不可寻址的，所以无法直接赋值。

修复代码：

```go
type T struct {
    n int
}

func main() {
    m := make(map[int]T)

    t := T{1}
    m[0] = t
    fmt.Println(m[0].n)
}
```

## 127. 下面的代码有什么问题？---- TBD

```go
type X struct {}

func (x *X) test()  {
    println(x)  // 这个为啥可以直接调用？
}

func main() {

    var a *X
    a.test()

    X{}.test()
}
```

**答：X{} 是不可寻址的，不能直接调用方法**

**解析：**

`cannot call pointer method test on X`

知识点：在方法中，指针类型的接收者必须是合法指针（包括 nil）,或能获取实例地址。

修复代码：

```go
func main() {

    var a *X
    a.test()    // 相当于 test(nil)

    var x = X{}
    x.test()
}
```

## 128. 下面代码有什么不规范的地方吗？

```go
func main() {
    x := map[string]string{"one":"a","two":"","three":"c"}

    if v := x["two"]; v == "" { 
        fmt.Println("no entry")
    }
}
```

**解析：**

检查 map 是否含有某一元素，直接判断元素的值并不是一种合适的方式。最可靠的操作是使用访问 map 时返回的第二个值。

修复代码如下：

```go
func main() {  
    x := map[string]string{"one":"a","two":"","three":"c"}

    if _,ok := x["two"]; !ok {
        fmt.Println("no entry")
    }
}
```



## 130. 下面的代码有几处问题？请详细说明。

```go
type T struct {
    n int
}

func (t *T) Set(n int) {
    t.n = n
}

func getT() T {
    return T{}
}

func main() {
    getT().Set(1)
}
```

**答：有两处问题**

**解析：**

- 1.直接返回的 T{} 不可寻址；
- 2.不可寻址的结构体不能调用带结构体指针接收者的方法；

修复代码：

```go
type T struct {
    n int
}

func (t *T) Set(n int) {
    t.n = n
}

func getT() T {
    return T{}
}

func main() {
    t := getT()
    t.Set(2)
    fmt.Println(t.n)
}
```

## 131. 下面的代码有什么问题？

```go
type N int

func (n N) value(){
    n++
    fmt.Printf("v:%p,%v\n",&n,n)
}

func (n *N) pointer(){
    *n++
    fmt.Printf("v:%p,%v\n",n,*n)
}

func main() {
    var a N = 25

    p := &a
    p1 := &p

    p1.value()
    p1.pointer()
}
```

**答：编译错误**

```shell
calling method value with receiver p1 (type **N) requires explicit dereference
calling method pointer with receiver p1 (type **N) requires explicit dereference
```

**解析：**

不能使用多级指针调用方法。

正确做法：
```go
type N int

func (n N) value() {
	n++
	fmt.Printf("v:%p,%v\n", &n, n)
}

func (n *N) pointer() {
	*n++
	fmt.Printf("v:%p,%v\n", n, *n)
}

func main() {
	var a N = 25

	p := &a
	p.value()
	p.pointer()
}
// 注意输出：
// v:0xc0000aa018,26
// v:0xc0000aa010,26
```

## 132. 下面的代码输出什么？

```go
type N int

func (n N) test(){
    fmt.Println(n)
}

func main()  {
    var n N = 10
    fmt.Println(n)

    n++
    f1 := N.test
    f1(n)

    n++
    f2 := (*N).test
    f2(&n)
}
```

**答：10 11 12**

**解析：**

知识点：方法表达式。

通过类型引用的方法表达式会被还原成普通函数样式，接收者是第一个参数，调用时显示传参。类型可以是 T 或 *T，只要目标方法存在于该类型的方法集中就可以。

还可以直接使用方法表达式调用：

```go
func main()  {
    var n N = 10

    fmt.Println(n)

    n++
    N.test(n)

    n++
    (*N).test(&n)
}
```



## 134. 下面的代码有什么问题？

```go
type T struct {
    n int
}

func getT() T {
    return T{}
}

func main() {
    getT().n = 1
}
```

**答：编译错误**

```shell
cannot assign to getT().n
```

**解析：**

结合130题。直接返回的 T{} 无法寻址，不可直接赋值。

修复代码：

```go
type T struct {
    n int
}

func getT() T {
    return T{}
}

func main() {
    t := getT()
    p := &t.n    // <=> p = &(t.n)
    *p = 1
    fmt.Println(t.n)
}
```


## 136. 下面代码输出什么？

```go
type N int

func (n N) test(){
    fmt.Println(n)
}

func main()  {
    var n N = 10
    p := &n

    n++
    f1 := n.test

    n++
    f2 := p.test

    n++
    fmt.Println(n)

    f1()
    f2()
}
```

**答：13 11 12**

**解析：**

知识点：方法值。结合132题。

当指针值赋值给变量或者作为函数参数传递时，会立即计算并复制该方法执行所需的接收者对象，与其绑定，以便在稍后执行时，能隐式第传入接收者参数。


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


## 141. 下面的代码输出什么？

```go
type N int

func (n *N) test(){
    fmt.Println(*n)
}

func main()  {
    var n N = 10
    p := &n

    n++
    f1 := n.test

    n++
    f2 := p.test

    n++
    fmt.Println(n)

    f1()
    f2()
}
```

**答：13 13 13**

**解析：**

知识点：方法值。结合136题。

当目标方法的接收者是指针类型时，那么被复制的就是指针。



## 143. 下面哪一行代码会 panic，请说明原因？

```go
package main

type T struct{}

func (*T) foo() {
}

func (T) bar() {
}

type S struct {
  *T
}

func main() {
  s := S{}
  _ = s.foo
  s.foo()
  _ = s.bar
}
```

**答：第 19 行**

**解析：**

因为 s.bar 将被展开为 (*s.T).bar，而 s.T 是个空指针，解引用会 panic。

可以使用下面代码输出 s：

```go
func main() {
    s := S{}
    fmt.Printf("%#v",s)   // 输出：main.S{T:(*main.T)(nil)}
}
```



## 145. 下面这段代码输出什么？

```go
func main() {
    var k = 1
    var s = []int{1, 2}
    k, s[k] = 0, 3
    fmt.Println(s[0] + s[1])
}
```

**答：4**

**解析：**

知识点：多重赋值。

多重赋值分为两个步骤，有先后顺序：

- 计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；
- 赋值；

所以本例，会先计算 s[k]，等号右边是两个表达式是常量，所以赋值运算等同于 `k, s[1] = 0, 3`。


## 147. 下面哪一行代码会 panic，请说明。

```go
func main() {
    nil := 123
    fmt.Println(nil)
    var _ map[string]int = nil
}
```

**答：第 4 行**

**解析：**

当前作用域中，预定义的 nil 被覆盖，此时 nil 是 int 类型值，不能赋值给 map 类型。

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


## 153. flag 是 bool 型变量，下面 if 表达式符合编码规范的是？

- A. if flag == 1
- B. if flag
- C. if flag == false
- D. if !flag

**答：B C D**

## 159. 下面代码有什么问题吗？

```go
func main()  {
    for i:=0;i<10 ;i++  {
    loop:
        println(i)
    }
    goto loop
}
```

**解析：**

goto 不能跳转到其他函数或者内层代码。编译报错：

```shell
goto loop jumps into block starting at
```

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




## 170. 下面的代码能编译通过吗？可以的话输出什么，请说明？

```go
var f = func(i int) {
    print("x")
}

func main() {
    f := func(i int) {
        print(i)
        if i > 0 {
            f(i - 1)
        }
    }
    f(10)
}
```

**答：10x**

**解析：**

这道题一眼看上去会输出 109876543210，其实这是错误的答案，这里不是递归。假设 main() 函数里为 f2()，外面的为 f1()，当声明 f2() 时，调用的是已经完成声明的 f1()。

看下面这段代码你应该会更容易理解一点：

```go
var x = 23

func main() {
    x := 2*x - 4
    println(x)    // 输出:42
}
```




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


