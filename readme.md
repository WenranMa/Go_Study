# Go 笔记

## 介绍

- 类型检查：编译时。
- 运行环境：编译成机器码直接运行 (静态编译)。
- 编程范式：面向借口，函数式编程，并发编程。

Go语言原生支持Unicode，它可以处理全世界任何语言的文本。Go语言的代码通过包(package)组织。一个包由位于单个目录下的一个或多个.go源代码文件组成。每个源文件都以一条package声明语句开始（例如package main）表示该文件属于哪个包，紧跟着一系列导入(import)的包。import声明必须跟在文件的package声明之后。

main包比较特殊。它定义了一个独立可执行的程序，而不是一个库。main函数也很特殊，它是整个程序执行时的入口(C系语言差不多都这样)。

必须恰当导入需要的包，缺少了必要的包或者导入了不需要的包，程序都无法编译通过。这项严格 要求避免了程序开发过程中引入未使用的包(Go语言编译过程没有警告信息，争议特性之一)。

Go语言不需要在语句或者声明的末尾添加分号，除非一行上有多条语句。实际上，编译器会主动 把特定符号后的换行符转换为分号，因此换行符添加的位置会影响Go代码的正确解析。举个例子, 函数的左括号{必须和函数声明在同一行上，且位于末尾，不能独占一行。而在表达式x+y中，可在+后换行，不能在+前换行（译注:以+结尾的话不会被插入分号分隔符，但是以x结尾的话则会被分号分隔符，从而导致编译错误）。

os包以跨平台的方式，提供了一些与操作系统交互的函数和变量。程序的命令行参数可从os包的Args变量获取。os.Args变量是一个字符串(string)的切片(slice是一个简版的动态数组)。和大多数编程语言类似，区间索引时，Go言里也采用左闭右开形式，区间包括第一个索引元素，不包括最后一个(比如`a = [1, 2, 3, 4, 5]`, `a[0:3] = [1, 2, 3]`)。os.Args的第一个元素，os.Args[0], 是命令本身的名字;其它的元素则是程序启动时传给它的参数，因此可以简写成os.Args[1:]。

自增语句`i++`给i加1;这和`i += 1`是等价的。这是语句，而不像C系的其它语言那样是表达式。所以`j = i++`非法，而且++和­­都只能放在变量名后面，因此`‐‐i`也非法。

Go语言只有for循环这一种循环语句（没有while）。for循环的这三个部分每个都可以省略。
```go
for initializaion, condition, post {
    //...
}
```
Go语言不允许使用无用的局部变量(local variables)，这会导致编译错误。解决方法是用空标识符(blank identifier)，即_(也就是下划线)。

bufio包使处理输入和输出方便高效。Scanner类型读取输入并将其拆成行或单词，通常是处理行形式的输入最简单的方法。
`input:=bufio.NewScanner(os.Stdin)`从程序的标准输入中读取内容。每次调用，即读入下一行，并移除行末的换行符，读取的内容可以调用`input.Text()`得到。

fmt.Printf函数对表达式产生格式化输出。
%d:十进制整数。%x,%o,%b:十六进制，八进制，二进制整数。%f,%g,%e:浮点数:3.141593 3.141592653589793 3.141593e+00。%t:布尔:true或false。%c:字符(rune) (Unicode码点)。%s:字符串。%q:带双引号的字符串"abc"或带单引号的字符'c'。%v:变量的自然形式(natural format)。%T:变量的类型。%%:字面上的百分号标志(无操作数)。按照惯例，以字母f结尾的格式化函数，如Printf和Errorf。而以ln结尾的在最后添加一个换行符。
 
os.Stdin
Fprintf ????
os.Stderr??
map函数参数传递？？？

实现上，bufio.Scanner、ioutil.ReadFile和ioutil.WriteFile都使用*os.File的Read和Write方法。

当我们import了一个包路径包含有多个单词的package时，比如image/color，通常我们只需要用最后那个单词表示这个包就可以。所以当我们写color.White时，这个变量指向的是image/color包里的变量，同理gif.GIF是属于image/gif包里的变量。

---
- 采用CSP (Communication Sequential Process)模型
- 不需要锁，不需要callback.

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

===============================================================================

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

---

不同于常规变量声明，在至少有一个非空白变量时，短变量声明可在相同块中，对原先声明的变量以相同的类型重声明。 因此，重声明只能出现在多变量短声明中。 重声明不能生成新的变量；它只能赋予新的值给原来的变量。

field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)  // 重声明 offset  在同一个块中使用var 重复声明是非法的
a, a := 1, 2                              // 非法：重复声明了 a，或者若 a 在别处声明，但此处没有新的变量