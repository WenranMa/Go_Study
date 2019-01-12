# Go 笔记

## 介绍

- 类型检查：编译时。
- 运行环境：编译成机器码直接运行 (静态编译)。
- 编程范式：面向借口，函数式编程，并发编程。

Go语言原生支持Unicode，它可以处理全世界任何语言的文本。

main包比较特殊。它定义了一个独立可执行的程序，而不是一个库。main函数也很特殊，它是整个程序执行时的入口(C系语言差不多都这样)。

Go语言不需要在语句或者声明的末尾添加分号，除非一行上有多条语句。实际上，编译器会主动 把特定符号后的换行符转换为分号，因此换行符添加的位置会影响Go代码的正确解析。举个例子, 函数的左括号{必须和函数声明在同一行上，且位于末尾，不能独占一行。而在表达式x+y中，可在+后换行，不能在+前换行（译注:以+结尾的话不会被插入分号分隔符，但是以x结尾的话则会被分号分隔符，从而导致编译错误）。

os包以跨平台的方式，提供了一些与操作系统交互的函数和变量。程序的命令行参数可从os包的Args变量获取。os.Args变量是一个字符串(string)的切片(slice是一个简版的动态数组)。和大多数编程语言类似，区间索引时，Go言里也采用左闭右开形式，区间包括第一个索引元素，不包括最后一个(比如`a = [1, 2, 3, 4, 5]`, `a[0:3] = [1, 2, 3]`)。os.Args的第一个元素，os.Args[0], 是命令本身的名字;其它的元素则是程序启动时传给它的参数，因此可以简写成os.Args[1:]。

Go语言只有for循环这一种循环语句（没有while）。for循环的这三个部分每个都可以省略。
```go
for initializaion, condition, post {
    //...
}
```
bufio包使处理输入和输出方便高效。Scanner类型读取输入并将其拆成行或单词，通常是处理行形式的输入最简单的方法。
`input:=bufio.NewScanner(os.Stdin)`从程序的标准输入中读取内容。每次调用，即读入下一行，并移除行末的换行符，读取的内容可以调用`input.Text()`得到。

fmt.Printf函数对表达式产生格式化输出。
%d:十进制整数。%x,%o,%b:十六进制，八进制，二进制整数。%f,%g,%e:浮点数:3.141593 3.141592653589793 3.141593e+00。%t:布尔:true或false。%c:字符(rune) (Unicode码点)。%s:字符串。%q:带双引号的字符串"abc"或带单引号的字符'c'。%v:变量的自然形式(natural format)。%T:变量的类型。%%:字面上的百分号标志(无操作数)。按照惯例，以字母f结尾的格式化函数，如Printf和Errorf。而以ln结尾的在最后添加一个换行符。
 
os.Stdin

Fprintf ????

os.Stderr??

map函数参数传递？？？

实现上，bufio.Scanner、ioutil.ReadFile和ioutil.WriteFile都使用*os.File的Read和Write方法。

---
- 采用CSP (Communication Sequential Process)模型
- 不需要锁，不需要callback.

---
## 程序结构

#### 命名
关键字有25个：

- break case chan const continue default func defer go else goto fallthrough if
for import interface map package range return select struct switch type var

此外，还有大约30多个预定义的名字，主要对应内建的常量、类型和函数。

- 内建常量: true false iota nil
- 内建类型: int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64 complex128 complex64 bool byte rune string error
- 内建函数: make len cap new append copy close delete complex real imag
panic recover

这些内部预先定义的名字并不是关键字，可以在定义中重新使用它们。

一个名字是在函数内部定义，就只在函数内部有效。如果是在函数外部定义，将 在当前包的所有文件中都可以访问。名字的开头字母的大小写决定了名字在包外的可见性。如果一个名字是大写字母开头的(译注:必须是在函数外部定义的包级名字;包级函数名本身也是包级名字)，那么它将是导出的，也就是说可以被外部的包访问，例如fmt包的Printf函数就是导出的，可以在fmt包外部访问。包本身的名字一般总是用小写字母。

Go语言主要有四种类型的声明语句:var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明。

#### 变量
`var 变量名字 类型 = 表达式` 其中“类型”或“=表达式”两个部分可以省略其中的一个。如果省略的是类型信息，那么将根据初始化表达式来推导变量的类型信息。如果初始化表达式被省略，那么将用零值初始化该变量。数值类型变量对应的零值是0，布尔类型变量对应的零值是false，字符串类型对应的零值是空字符串，接口或引用类型(包括slice、指针、map、chan和函数)变量对应的零值是nil。数组或结构体等聚合类型对应的零值是每个元素或字段都是对应该类型的零值。Go语言中不存在未初始化的变量。

可以在一个声明语句中同时声明一组变量，或用一组初始化表达式声明并初始化一组变量。
```go
var i, j, k int // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string
```
在函数内部，有一种称为__短声明__的形式可用于声明和初始化局部变量。它以`名字 := 表达式`形式声明变量，变量的类型根据表达式来自动推导。
```go
anim := gif.GIF{LoopCount: nframes} 
freq := rand.Float64() * 3.0
t := 0.0
```
简短变量声明语句也可以用来声明和初始化一组变量: `i, j := 0, 1`。

多值简短变量声明左边的变量可能并不是全部都是刚刚声明的。如果有一些已经在相同的词法域声明过了，那么简短变量声明语句对这些已经声明过的变量就只有赋值行为了。
在下面第一句声明了in和err两个变量。在第二句只声明了out一个变量，然后对已经声明的err进行了赋值操作：
```go
in, err := os.Open(infile)
out, err := os.Create(outfile)
```
简短变量声明语句中必须至少要声明一个新的变量，下面的代码将不能编译通过:
```go
f, err := os.Open(infile)
f, err := os.Create(outfile) // compile error: no new variables
```
简短变量声明语句只有对已经在同级词法域声明过的变量才和赋值操作语句等价，如果变量是在外 部词法域声明的，那么简短变量声明语句将会在当前词法域重新声明一个新的变量。

任何类型的指针的零值都是nil。如果p指向某个有效变量，那么p != nil测试为真。指针之间也是 可以进行相等测试的，只有当它们指向同一个变量或全部是nil时才相等。
```go
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
```
在Go语言中，返回函数中局部变量的地址也是安全的。下面的代码，调用f函数时创建局部变量v，在局部变量地址被返回之后依然有效，因为指针p依然引用这个变量。
```go
var p = f()
func f() *int { 
    v := 1
    return &v 
}
```
每次调用f函数都将返回不同的结果:
`fmt.Println(f() == f()) // "false"`

另一个创建变量的方法是调用用内建的new函数。表达式new(T)将创建一个T类型的匿名变量，初
始化为T类型的零值，然后返回变量地址，返回的指针类型为*T。
```go
p := new(int) // p, *int 类型, 指向匿名的 int 变量 
fmt.Println(*p) // "0"
```
new只是一个预定义的函数，它并不是一个关键字，因此我们可以将new名字重新定义为别的类型。例如下面的例子: `var new int = 1` 由于new被定义为int类型的变量名，因此在函数内部是无法使用内置的new函数的。

#### 赋值
自增语句`i++`给i加1;这和`i += 1`是等价的。这是语句，而不像C系的其它语言那样是表达式。所以`j = i++`非法，而且++和­­都只能放在变量名后面，因此`--i`也非法。

Go允许同时更新多个变量的值。在赋值之前，赋值语句右边的所有表达式将会先进行求值，然后再统一更新左边对应变量的值。这样交换两个变量的值:
```go
x, y = y, x
a[i], a[j] = a[j], a[i]
```
有些表达式会产生多个值，比如map查找、类型断言或通道接收，它们都可能会产生两个结果，有一个额外的布尔结果表示操作是否成功:
```go
v, ok = m[key] // map lookup
v, ok = x.(T)  // type assertion
v, ok = <-ch   // channel receive
```
Go语言不允许使用无用的局部变量(local variables)，这会导致编译错误。解决方法是用空标识符(blank identifier)，即_(也就是下划线)。

#### 类型
一个类型声明语句创建了一个新的类型名称。新命名的类型用来分隔不同概念的类型，这样即使它们底层类型相同也是不兼容的。

`type 类型名字 底层类型` 

类型声明语句一般出现在包一级，因此如果新创建的类型名字的首字符大写，则在外部包也可以使
用。对于中文汉字，Unicode标志都作为小写字母处理，因此中文的命名默认不能导出，在Go2中有可能会将中日韩等字符当作大写字母处理。

对于每一个类型T，都有一个对应的类型转换操作T(x)，用于将x转为T类型(译注:如果T是指针类 型，可能会需要用小括弧包装T，比如 (*int)(0))。只有当两个类型的底层基础类型相同时，才允 许这种转型操作，或者是两者都是指向相同底层结构的指针类型，这些转换只改变类型而不会影响值本身。数值类型之间的转型也是允许的，并且在字符串和一些特定类型的slice之间也是可以转换的。

底层数据类型决定了内部结构和表达方式，也决定是否可以像底层类型一样对内置运算符的支持。这意味着，`type Celsius float64`和 `type Fahrenheit float64`类型的算术运算行为和底层的float64类型是一样的。
```go
var c Celsius
var f Fahrenheit 
fmt.Println(c == 0) // "true"
fmt.Println(f >= 0) // "true"
fmt.Println(c == f) // compile error: type mismatch
fmt.Println(c == Celsius(f))  // "true"!
fmt.Println(c - f)  // compile error: type mismatch
```

#### 包和文件
Go语言的代码通过包(package)组织。一个包由位于单个目录下的一个或多个.go源代码文件组成。每个源文件都以一条package声明语句开始（例如package main）表示该文件属于哪个包，紧跟着一系列导入(import)的包。import声明必须跟在文件的package声明之后。

必须恰当导入需要的包，缺少了必要的包或者导入了不需要的包，程序都无法编译通过。这项严格 要求避免了程序开发过程中引入未使用的包(Go语言编译过程没有警告信息，争议特性之一)。在Go语言程序中，每个包都是有一个全局唯一的导入路径。例如：`"github.com/influxdata/influxdb/client/v2"`。

当我们import了一个包路径包含有多个单词的package时，比如image/color，通常我们只需要用最后那个单词表示这个包就可以。所以当我们写color.White时，这个变量指向的是image/color包里的变量，同理gif.GIF是属于image/gif包里的变量。

每个包都对应一个独立的名字空间。例如，在image包中的Decode函数和在unicode/utf16包中的 Decode函数是不同的。要在外部引用该函数，必须显式使用image.Decode或utf16.Decode形式访问。

包的初始化首先是解决包级变量的依赖顺序，然后按照包级变量声明出现的顺序依次初始化:
```go
var a = b + c//a 第三个初始化, 为 3
var b = f() //b 第二个初始化, 为 2, 通过调用 f(依赖c) 
var c = 1 //c 第一个初始化, 为 1
func f() int { 
    return c + 1 
}
```
一个特殊的init初始化函数来简化初始化工作。每个文件都可以包含多个init初始化函数
`func init() { /* ... */ }` 这样的init初始化函数除了不能被调用或引用外，其他行为和普通函数类似。

#### 作用域
不要将作用域和生命周期混为一谈。声明语句的作用域对应的是一个源代码的文本区域;它是一个 编译时的属性。一个变量的生命周期是指程序运行时变量存在的有效时间段，是一个运行时的概念。

语法块定了内部声明的名字的作用域范围。有一个语法块为整个源代码，称为全局语法块;然后是每个包的包语法块;每个for、if和switch语句的语法块;每个switch或select的分支也有独立的语法块;当然也包括显式书写的语法块(花括弧 包含的语句)。

当编译器遇到一个名字引用时，它首先从最内层的词法域向全局作用域查找。如果查找失败，则报告“未声明的名字”这样的错误。如果该名字在内部和外部的块分别声明过，则内部块的声明首先被找到。在这种情况下，内部声明屏蔽了外部同名的声明，让外部的声明的名字无法被访问。
```go
var cwd string
func init() {
    cwd, err := os.Getwd()  // compile error: unused: cwd 
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err) 
    }
}
```
虽然cwd在外部已经声明过，但是 := 语句还是将cwd和err重新声明为新的局部变量。因为内部声明 的cwd将屏蔽外部的声明，因此上面的代码并不会正确更新包级声明的cwd变量。最直接的方法是通过单独声明err变量，来避免使用:=
的简短声明方式:
```go
var cwd string
func init() {
    var err error
    cwd, err = os.Getwd() 
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
}
```
---

## 基础数据类型
Go语言将数据类型分为四类:基础类型、复合类型、引用类型和接口类型。

基础类型包括:数字、字符串和布尔型。复合数据类型——数组和结构体，是通过组合简单类型，来表达更加复杂的数据结构。引用类型包括指针、切片、字典、函数、通道，虽然数据种类很多，但它们都是对程序中一个变量或状态的间接引用。这意味着对任一引用类型数据的修改都会影响所有该引用的拷贝。

#### 整型
Go语言同时提供了有符号和无符号类型的整数运算。这里有int8、int16、int32和int64四种不同大小的有符号整数类型，分别对应8、16、32、64bit大小的有符号整数，与此对应的是uint8、uint16、uint32和uint64四种无符号整数类型。另外两种特定CPU平台机器字大小的有符号和无符号整数int和uint;这两种类型都有同样的大小，32或64bit，不同的编译器即使在相同的硬件平台上可能产生不同的大小。int和int32是不同的类型，即使int的大小也是32bit，在需要将int当作int32类型的地方需要一个显式的类型转换操作。

Unicode字符rune类型是和int32等价的类型。这两个名称可以互换使用。同样byte也是uint8类型的等价类型，byte类型一般用于强调数值是一个原始的数据而不是一个小的整数。

还有一种无符号的整数类型uintptr，没有指定具体的bit大小但是足以容纳指针。uintptr类型 只有在底层编程时才需要，特别是Go语言和C语言函数库或操作系统接口相交互的地方。

其中有符号整数采用2的__补码__形式表示，也就是最高bit位用来表示符号位，一个n­ bit的有符号数的 值域是从`−2^(n−1)`到`2^(n-1)−1`。无符号整数的所有bit位都用于表示非负数，值域是0到`2^n-1`。例如，int8类型整数的值域是从­-128到127，而uint8类型整数的值域是从0到255。

下面是Go语言中关于算术运算、逻辑运算和比较运算的二元运算符，它们按照先级递减的顺序的排列:
```go
*  /  %  <<  >>  &  &^ 
+  -  |  ^
== != <  <=  >  >=
&&
||
```
二元运算符有五种优先级。在同一个优先级，使用左优先结合规则，但是使用括号可以明确优先顺 序，使用括号也可以用于提升优先级，例如mask & (1 << 28)。

取模运算符%仅用于整数间的运算。在Go语言中，%取模运算符的符号和被取模数的符号总是一致的，因此 `-5%3` 和 `-5%-3`结果都是­2。除法运算符 / 的行为则依赖于操作数是否为全为整数，比如 5.0/4.0 的结果是1.25，但是5/4的结果是1，因为整数除法会向着0方向截断余数。

计算结果是溢出，超出的高位的bit位部分将被丢弃。如果原始的数值是有符号类型，而且最左边的bit为是1的话，那么最终结果可能是负的：
```go
var u uint8 = 255
fmt.Println(u, u+1, u*u) // "255 0 1"
var i int8 = 127
fmt.Println(i, i+1, i*i) // "127 -128 1"
```

位操作运算符 ^ 作为二元运算符时是按位异或(XOR)，当用作一元运算符时表示按位取反;也就 是说，它返回一个每个bit位都取反的数。位操作运算符&^用于按位置零(AND NOT):如果对应y中bit位为1的话, 表达式`z = x &^ y`结果z的对应的bit位为0，否则z对应的bit位等于x相应的bit位的值。
```go
var x uint8 = 1<<1 | 1<<5 
var y uint8 = 1<<1 | 1<<2
fmt.Printf("%08b\n", x) // "00100010"
fmt.Printf("%08b\n", y) // "00000110"
fmt.Printf("%08b\n", x&y) // "00000010"
fmt.Printf("%08b\n", x|y) // "00100110"
fmt.Printf("%08b\n", x^y) // "00100100" 
fmt.Printf("%08b\n", x&^y) // "00100000"
fmt.Printf("%08b\n", x<<1) // "01000100"
fmt.Printf("%08b\n", x>>1) // "00010001"
```
整数字面值都可以用以0开始的八进制格式书写，例如0666;或用以0x或0X开头的十六进制格式书写，例如0xdeadbeef。十六进制数字可以用大写或小写字母。

当使用fmt包打印一个数值时，我们可以用%d、%o或%x参数控制输出的进制格式:
```go
o := 0666
fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666" 
x := int64(0xdeadbeef)
fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
```
%之后的[1]副词告诉Printf函数再次使用第一个操作数。%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。

#### 浮点型
浮点数的范围极限值可以在math包找到。常量math.MaxFloat32、math.MaxFloat64。
math包提供了IEEE754浮点数标准中定义的特殊值:正无穷大和负无穷大，分别用于表示太大溢出的数字和除零的结果;还有NaN非数表示无效的除法操作结果0/0或Sqrt(­-1)。
```go
var z float64
fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
```

#### 复数
复数类型:complex64和complex128，分别对应float32和float64两种浮 点数精度。内置的complex函数用于构建复数，内建的real和imag函数分别返回复数的实部和虚部:
```go
var x complex128 = complex(1, 2)
var y complex128 = complex(3, 4)
fmt.Println(x*y)
fmt.Println(real(x*y))
fmt.Println(imag(x*y))
```

#### 字符串

#### 常量
常量表达式的值在编译期计算，而不是在运行期。常量的值不可修改，这样可以防止在运行期被意外或恶意的修改。如：`const pi = 3.14159` 。可以批量声明多个常量：
```go
const(
    e = 2.718
    pi = 3.14
)
```
常量间的所有算术运算、逻辑运算和比较运算的结果也是常量。一个常量的声明也可以包含一个类型和一个值，但是如果没有显式指明类型，那么将从右边的表达式推断类型。

常量声明可以使用iota常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每 行都写一遍初始化表达式。在一个const声明语句中，在第一个声明的常量所在的行，iota将会被置 为0，然后在每一个有常量声明的行加一。
```go
type Weekday int
const (
    Sunday Weekday = iota 
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```
周日将对应0，周一为1，如此等等。在其它编程语言中，这种类型一般被称为枚举类型。

复杂的常量表达式中使用iota:
```go
const (
    i = 1 << iota   // i = 1 << 0, 1 
    j               // j = 1 << 1, 2 
    k               // j = 1 << 2, 4
)
```

有六种未明确类型的常量类型，分别是无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串。通过延迟明确常量的具体类型，无类型的常量不仅可以提供更高的运算精度，而且可以直接用于更多的表达式而不需要显式的类型转换。

例如:math.Pi无类型的浮点数常量: `var x float32 = math.Pi` `var y float64 = math.Pi` `var z complex128 = math.Pi`。

---

## 复合数据类型
数组和结构体是聚合类型;它们的值由许多元素或成员字段的值组成。数组元素都是完全相同的类型;结构体则是由异构的元素组成的。数组和结构体都是有固定内存大小的数据结构。slice和map则是动态的数据结构，它们将根据需要动态增长。

#### 数组
一个数组可以由零个或多个元素组成。数组的长度是固定的。定义方式：
```go
var a [3]int
var b [3]int = [3]int{1, 2, 3} 
var c [3]int = [3]int{1, 2}
d := [...]int{1, 2, 2, 4}
```
数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型。数组的长度必 须是常量表达式，因为数组的长度需要在编译阶段确定。

#### Slice
Slice(切片)代表变长的序列，序列中每个元素都有相同的类型。slice的语法和数组很像，只是没有固定长度而已。而且slice的底层引用一个数组对象。一个slice由三个部分构成:指针、长度和容量。指针指向第一个slice元素对应的底层数组元素的地址，要注意的是slice的第一个元素并不一定就是数组的第一个元素。长度对应slice中元素的数目;长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置。内置的len和cap函数分别返回slice的长度和容量。

#### Map
一个无序的key/value对的集合，其中所有的key都是不同的，然后通过给定的key可以在__常数时间复杂度__内检索、更新或删除对应的value。map类型可以写为map[K]V，其中K和V分别对应key和value。map中所有的key都有相同的类型，所有的value也有着相同的类型，但是key和value之间可以是不同的数据类型。其中K对应的key必须是支持`==`比较运算符的数据类型，所以map可以通过测试key是否相等来判断是否已经存在。



#### Struct

#### JSON

#### 模板Template

---

## 函数

#### 递归

#### 多值返回

#### Error

#### 函数值

#### 匿名函数

#### 可变参数

#### Defer

#### Panic

#### Recover

---

## 方法

#### 方法声明

#### 指针对象方法

#### 嵌入结构体扩展类型

#### 方法值 方法表达式

#### Bit数组

#### 封装

---

## 接口

#### 接口约定

#### 接口类型 

#### 实现接口的条件

#### flag.Value接口

#### 接口值

#### sort.Interface接口

#### http.Handler接口

#### error接口

#### 示例：表达式请求

#### 类型断言

#### 基于类型断言区别错误类型

#### 通过类型断言询问行为

#### 类型开关

#### 示例：基于标记的XML解码

---

## Goroutines和Channels

#### Goroutines

#### 示例：并发的Clock服务

#### 示例：并发的Echo服务

#### Channels

#### 并发的循环

#### 示例：并发Web爬虫

#### 基于Select的多路复用

#### 示例：并发的字典遍历

#### 并发的退出

#### 示例：聊天服务

---

## 基于共享变量的并发

#### 竞争条件

#### sync.Mutex出斥锁

#### sync.RWMutex读写锁

#### 内存同步

#### sync.Once初始化

#### 竞争条件检测

#### 示例：并发的非阻塞缓存

#### Goroutines和线程

--- 

## 反射

#### 为什么反射

#### refect.Type和reflect.Value

#### Display,递归打印器

#### 示例：编码为S表达式

#### 通过reflect.Value修改值

#### 示例：解码S表达式

#### 获取结构体字段标识

#### 显示一个类型的方法集

---

## IO
在Go中，输入和输出操作是使用原语实现的，这些原语将数据模拟成可读的或可写的字节流。Go的io包提供了io.Reader和io.Writer接口，分别用于数据的输入和输出，如图：
![io](./file/img/io.png)

### io.Reader
io.Reader表示一个读取器，它将数据从某个资源读取到传输缓冲区。在缓冲区中，数据可以被流式传输和使用。对于要用作读取器的类型，它必须实现io.Reader接口的唯一方法 `Read(p []byte)`。换句话说，只要实现了`Read(p []byte)`，那它就是一个读取器。
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```
Read()方法有两个返回值，一个是读取到的字节数，一个是发生错误时的错误。同时，如果资源内容已全部读取完毕，应该返回io.EOF错误。

### io.Writer
io.Writer表示一个编写器，它从缓冲区读取数据，并将数据写入目标资源。对于要用作编写器的类型，必须实现io.Writer接口的唯一方法`Write(p []byte)`，只要实现了`Write(p []byte)`，那它就是一个编写器。
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
Write()方法有两个返回值，一个是写入到目标资源的字节数，一个是发生错误时的错误。

### 其他用到io.Reader/io.Writer的类型，方法
类型os.File表示本地系统上的文件。它实现了io.Reader和io.Writer，因此可以在任何io上下文中使用。

缓冲区io，标准库中bufio包支持缓冲区io操作，可以轻松处理文本内容。

ioutil，io包下面的一个子包ioutil封装了一些非常方便的功能，例如，使用函数ReadFile将文件内容加载到[]byte中。

---

## Project List: 

### IO
简单的IO练习

### Pipeline: 外部排序Pipeline
选自慕课网：搭建并行处理管道

- 原始数据过大，无法一次读入内存，所以分块读入内存。每个块数据进行内部排序（直接调用API排序），最后讲各个节点归并，归并选择两两归并。

---


1.

slice1:= slice[0:2]
引用，非复制，所以任何对slice1或slice的修改都会影响对方

data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
data1 := data[0:2]
data1[0] = 99
fmt.Println(data1)
fmt.Println(data)
[99 2]

[99 2 3 4 5 6 7 8 9 0]



2.append


append 比较特殊

声明:

源slice= src

添加slice = app

结果slice=tar

1）如果len(src) + len(app) <= cap(src)  src和tar 是指向同一数据引用 ，即修改src或tar，会影响对方

2）否则 tar 是copy的方式 src + app ，即修改src或tar，不会影响对方

无论哪种情况不会影响app，因为app都会用copy的方式进入tar

func test2() {
    data := make([]int, 10, 20)
    data[0] = 1
    data[1] = 2
    dataappend := make([]int, 10, 20)//len <=10 则   result[0] = 99 会 影响源Slice
    dataappend[0] = 1
    dataappend[1] = 2
    result := append(data, dataappend...)
    result[0] = 99
    result[11] = 98
    fmt.Println("length:", len(data), ":", data)
    fmt.Println("length:", len(result), ":", result)
    fmt.Println("length:", len(dataappend), ":", dataappend)
}

---

命名返回参数可被同名局部变量遮蔽，此时需要显式返回。

func add(x, y int) (z int) {
    { // 不能在一个级别，引发 "z redeclared in this block" 错误。
        var z = x + y
        // return   // Error: z is shadowed during return
        return z // 必须显式返回。
    }
}



函数的有右小括弧也可以另起一行缩进，同时为了防止编译器在行尾自动插入分号而导致的 编译错误，可以在末尾的参数变量后面显式插入逗号。




em必须为interface类型才可以进行类型断言

比如下面这段代码，

s := "hello world"
if v, ok := s.(string); ok {
    fmt.Println(v)
}
运行报错， invalid type assertion: s.(string) (non-interface type string on left)

在这里只要是在声明时或函数传进来的参数不是interface类型那么做类型断言都是回报 non-interface的错误的 所以我们只能通过将s作为一个interface{}的方法来进行类型断言，如下代码所示：

x := "hello world"
if v, ok := interface{}(x).(string); ok { // interface{}(x):把 x 的类型转换成 interface{}
    fmt.Println(v)
}


断言返回值个数 不一定是两个
