# Go Libs

## fmt
### 常用函数：
```go
func Errorf(format string, a ...interface{}) error
//Errorf returns the string as a value that satisfies error.
```
```go
package main
import (
    "fmt"
)
func main() {
    const name, age = "wrma", 32
    err := fmt.Errorf("user %q (age %d) not found", name, age)
    fmt.Printf("Error: %v", err)
}
//Error: user "wrma" (age 32) not found
```
```go
func Print(a ...interface{}) (n int, err error)
//直接打印

func Printf(format string, a ...interface{}) (n int, err error)
//格式打印

func Println(a ...interface{}) (n int, err error)
//直接打印并换行
```
以上三个函数分别用下面三个函数实现：
```go
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
```
```go
func Sprintf(format string, a ...interface{}) string
```
格式化并返回一个string。
```go
package main
import (
    "fmt"
    "io"
    "os"
)
func main() {
    const name, age = "wrma", 32
    s := fmt.Sprintf("%s is %d years old.\n", name, age)
    io.WriteString(os.Stdout, s) // Ignoring error for simplicity.
}
// wrma is 32 years old.
```
以上函数接受的参数都是接口类型！

### fmt.Stringer接口：
该接口包含String()函数。任何类型只要定义了String()函数，进行Print输出时，就可以得到定制输出。fmt.Println会判断这个变量是否实现了Stringer接口，如果实现了，则调用这个变量的String()方法。
```go
type Stringer interface {
    String() string
}

package main
import (
    "fmt"
)

type Animal struct {
    Name string
    Age  uint
}
// String makes Animal satisfy the Stringer interface.
func (a Animal) String() string {
    return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}
func main() {
    a := Animal{
        Name: "Gopher",
        Age:  2,
    }
    fmt.Println(a)
}
// Gopher (2)
```

---

## io
Go的io包提供了io.Reader和io.Writer接口，分别用于数据的输入和输出。

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
#### os.File
类型os.File表示本地系统上的文件。它实现了io.Reader和io.Writer，因此可以在任何io上下文中使用。
```go
func main() {
    test := []string{"Hello ", "World\n",}
    file, err := os.Create("./test.txt") //返回*os.File类型
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    for _, p := range test {
        // file 类型实现了 io.Writer
        n, err := file.Write([]byte(p))
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        if n != len(p) {
            fmt.Println("failed to write data")
            os.Exit(1)
        }
    }
    fmt.Println("file write done")
}

func main() {
    file, err := os.Open("./test.txt") //返回*os.File类型
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    p := make([]byte, 4)
    for {
        n, err := file.Read(p)
        if err == io.EOF {
            break
        }
        fmt.Print(string(p[:n]))
    }
}
```

#### os.Stdout/os.Stdin/os.Stderr
os包中的三个变量，它们的类型为*os.File，分别代表系统标准输入，系统标准输出和系统标准错误的文件句柄。因为是*os.File类型，所以也实现了io.Read和io.Write方法。
```go
func main() {
    test := []string{"Hello ", "World\n"}
    for _, p := range test {
        os.Stdout.Write([]byte(p))
    }
}
```

#### io.Copy()
io.Copy()可以轻松地将数据从一个Reader拷贝到另一个Writer。接收的参数为writer和reader。
```go
func main() {
    file, _ := os.Open("./test.txt")
    defer file.Close()
    io.Copy(os.Stdout, file) //os.Stdout, file都是*os.File类型，所以是writer和reader.
}
```

#### io.WriteString()
将字符串类型写入一个Writer：`io.WriteString(file, "Go is fun!")`

#### bufio
缓冲区io，标准库中bufio包支持缓冲区io操作，可以轻松处理文本内容。
例如`bufio.NewReader`和`bufio.NewWriter`本别可以接收一个`io.Reader`和`io.Writer`参数，然后返回`*bufio.Reader`和`*bufio.Writer`。可以实现更高效的读写操作。

#### ioutil
io包下面的一个子包ioutil封装了一些非常方便的功能，例如，使用函数ReadFile将文件内容加载到[]byte中。`ioutil.ReadFile`和`ioutil.WriteFile`都使用`*os.File`的Read和Write方法。

#### strings.NewReader()
strings包下也有strings.Reader类型，可以将string打包，返回`*strings.Reader`返回。

---

## stirngs
Package strings implements simple functions to manipulate UTF-8 encoded strings.
##### 常用函数：
```go
Contains(s, substr string) bool, //用Index()函数实现。
Index(s, substr string) int //返回字符index。
Split(s sep string) []string //切割字符串，返回字符串数组。
Join(a []string, sep string) string //合并字符串。
HasPrefix(s, prefix string) bool
HasSuffix(s, suffix string) bool //判断是否有前缀或后缀，返回bool。
```

## strconv
Package strconv implements conversions to and from string representations of basic data types.
##### 常用函数：
```go
Aoti(s string) (int, error) //String to Integer.
Itoa(i int) string // Integer to string. FormatInt()函数实现。
ParseBool(s string) (bool, error) //字符串1,t,T,TRUE,true,True,0,f,F,FALSE,false,False都可以返回。
//还有其他类型的parse函数。将字符串parse成对应类型。
FormatBool(b bool) string //将boolean转成string.
//还有其他类型的format函数，将该类型转成string.
```

https://blog.golang.org/strings

---

### encoding/xml encoding/json
Marshal() 返回[]byte数组
print(string([]byte))

MarshalIndent()

Unmarshl

可以用于sturct to xml
结构体加tag `xml:"age,attr"` 可以把结构体的field变成 xml标签中的属性。

xml.NewDecoder
for循环遍历decoder.Tocken()



### os.Args

### flag
flag.String()
flag.Int()
flag.Parse()

flag.StringVar()
flag.IntVar()

NArg()
Usage()



---


## text
text/tabwriter



---

## time
time包提供了时间的显示和计量用的功能。日历的计算采用的是公历。提供的主要类型有Time, Duration, Location, Timer, Ticker等。

#### Time
代表一个纳秒精度的时间点，是公历时间。

#### Duration
代表两个时间点之间经过的时间，以纳秒为单位。可表示的最长时间段大约290年，也就是说如果两个时间点相差超过290年，会返回290年，也就是minDuration(-1 << 63)或maxDuration(1 << 63 - 1)。

类型定义：type Duration int64。

将Duration类型直接输出时，因为实现了fmt.Stringer 接口，会输出人类友好的可读形式，如：72h3m0.5s。

1.Location
代表一个地区，并表示该地区所在的时区（可能多个）。Location 通常代表地理位置的偏移，比如 CEST 和 CET 表示中欧。下一节将详细讲解 Location。

1.4. Timer 和 Ticker

1.5. Weekday 和 Month
这两个类型的原始类型都是 int，定义它们，语义更明确，同时，实现 fmt.Stringer 接口，方便输出。

---

## os/signal
信号是事件发生时对进程的通知机制。有时也称之为软件中断。因为一个具有合适权限的进程可以向另一个进程发送信号，这可以称为进程间的一种同步技术。当然，进程也可以向自身发送信号。然而，发往进程的诸多信号，通常都是源于内核。引发内核为进程产生信号的各类事件有：

- 硬件发生异常，即硬件检测到一个错误条件并通知内核，随即再由内核发送相应信号给相关进程。比如执行一条异常的机器语言指令（除0，引用无法访问的内存区域）。
- 用户键入了能够产生信号的终端特殊字符。如中断字符（通常是 Control-C）、暂停字符（通常是 Control-Z）。
- 发生了软件事件。如调整了终端窗口大小，定时器到期等。

针对每个信号，都定义了一个唯一的（小）整数，从1开始顺序展开。系统会用相应常量表示。Linux 中，1-31 为标准信号；32-64 为实时信号（通过 kill -l 可以查看）。

信号到达后，进程视具体信号执行如下默认操作之一：

- 忽略信号，也就是内核将信号丢弃，信号对进程不产生任何影响。
- 终止（杀死）进程。
- 产生 coredump 文件，同时进程终止。
- 暂停（Stop）进程的执行。
- 恢复进程执行。

当然，对于有些信号，程序是可以改变默认行为的，这也就是 os/signal 包的用途。

兼容性问题：信号的概念来自于 Unix-like 系统。Windows 下只支持 os.SIGINT 信号。

#### Go对信号的处理
程序无法捕获信号SIGKILL和SIGSTOP（终止和暂停进程），因此os/signal包对这两个信号无效。

1.2.1. Go 程序对信号的默认行为
Go 语言实现了自己的运行时，因此，对信号的默认处理方式和普通的 C 程序不太一样。

SIGBUS（总线错误）, SIGFPE（算术错误）和 SIGSEGV（段错误）称为同步信号，它们在程序执行错误时触发，而不是通过 os.Process.Kill 之类的触发。通常，Go 程序会将这类信号转为 run-time panic。
SIGHUP（挂起）, SIGINT（中断）或 SIGTERM（终止）默认会使得程序退出。
SIGQUIT, SIGILL, SIGTRAP, SIGABRT, SIGSTKFLT, SIGEMT 或 SIGSYS 默认会使得程序退出，同时生成 stack dump。
SIGTSTP, SIGTTIN 或 SIGTTOU，这是 shell 使用的，作业控制的信号，执行系统默认的行为。
SIGPROF（性能分析定时器，记录 CPU 时间，包括用户态和内核态）， Go 运行时使用该信号实现 runtime.CPUProfile。
其他信号，Go 捕获了，但没有做任何处理。
信号可以被忽略或通过掩码阻塞（屏蔽字 mask）。忽略信号通过 signal.Ignore，没有导出 API 可以直接修改阻塞掩码，虽然 Go 内部有实现 sigprocmask 等。Go 中的信号被 runtime 控制，在使用时和 C 是不太一样的。

1.2.2. 改变信号的默认行为
这就是 os/signal 包的功能。

Notify 改变信号处理，可以改变信号的默认行为；Ignore 可以忽略信号；Reset 重置信号为默认行为；Stop 则停止接收信号，但并没有重置为默认行为。

1.2.3. SIGPIPE
文档中对这个信号单独进行了说明。如果 Go 程序往一个 broken pipe 写数据，内核会产生一个 SIGPIPE 信号。

如果 Go 程序没有为 SIGPIPE 信号调用 Notify，对于标准输出或标准错误（文件描述符1或2），该信号会使得程序退出；但其他文件描述符对该信号是啥也不做，当然 write 会返回错误 EPIPE。

如果 Go 程序为 SIGPIPE 调用了 Notify，不论什么文件描述符，SIGPIPE 信号都会传递给 Notify channel，当然 write 依然会返回 EPIPE。

也就是说，默认情况下，Go 的命令行程序跟传统的 Unix 命令行程序行为一致；但当往一个关闭的网络连接写数据时，传统 Unix 程序会 crash，但 Go 程序不会。


1.3. signal 中 API 详解
1.3.1. Ignore 函数
func Ignore(sig ...os.Signal)

忽略一个、多个或全部（不提供任何信号）信号。如果程序接收到了被忽略的信号，则什么也不做。对一个信号，如果先调用 Notify，再调用 Ignore，Notify 的效果会被取消；如果先调用 Ignore，在调用 Notify，接着调用 Reset/Stop 的话，会回到 Ingore 的效果。注意，如果 Notify 作用于多个 chan，则 Stop 需要对每个 chan 都调用才能起到该作用。

1.3.2. Notify 函数
func Notify(c chan<- os.Signal, sig ...os.Signal)

类似于绑定信号处理程序。将输入信号转发到 chan c。如果没有列出要传递的信号，会将所有输入信号传递到c；否则只传递列出的输入信号。

channel c 缓存如何决定？因为 signal 包不会为了向c发送信息而阻塞（就是说如果发送时 c 阻塞了，signal包会直接放弃）：调用者应该保证 c 有足够的缓存空间可以跟上期望的信号频率。对使用单一信号用于通知的channel，缓存为1就足够了。

相关源码：

// src/os/signal/signal.go process 函数
for c, h := range handlers.m {
    if h.want(n) {
        // send but do not block for it
        select {
        case c <- sig:
        default:    // 保证不会阻塞，直接丢弃
        }
    }
}
可以使用同一 channel 多次调用 Notify：每一次都会扩展该 channel 接收的信号集。唯一从信号集去除信号的方法是调用 Stop。可以使用同一信号和不同 channel 多次调用 Notify：每一个 channel 都会独立接收到该信号的一个拷贝。

1.3.3. Stop 函数
func Stop(c chan<- os.Signal)

让 signal 包停止向 c 转发信号。它会取消之前使用 c 调用的所有Notify 的效果。当 Stop 返回后，会保证 c 不再接收到任何信号。

1.3.4. Reset 函数
func Reset(sig ...os.Signal)

取消之前使用 Notify 对信号产生的效果；如果没有参数，则所有信号处理都被重置。

1.3.5. 使用示例
注：syscall 包中定义了所有的信号常量

package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

var firstSigusr1 = true

func main() {
    // 忽略 Control-C (SIGINT)
    // os.Interrupt 和 syscall.SIGINT 是同义词
    signal.Ignore(os.Interrupt)

    c1 := make(chan os.Signal, 2)
    // Notify SIGHUP
    signal.Notify(c1, syscall.SIGHUP)
    // Notify SIGUSR1
    signal.Notify(c1, syscall.SIGUSR1)
    go func() {
        for {
            switch <-c1 {
            case syscall.SIGHUP:
                fmt.Println("sighup, reset sighup")
                signal.Reset(syscall.SIGHUP)
            case syscall.SIGUSR1:
                if firstSigusr1 {
                    fmt.Println("first usr1, notify interrupt which had ignore!")
                    c2 := make(chan os.Signal, 1)
                    // Notify Interrupt
                    signal.Notify(c2, os.Interrupt)
                    go handlerInterrupt(c2)
                }
            }
        }
    }()

    select {}
}

func handlerInterrupt(c <-chan os.Signal) {
    for {
        switch <-c {
        case os.Interrupt:
            fmt.Println("signal interrupt")
        }
    }
}
编译后运行，先后给该进程发送如下信号：SIGINT、SIGUSR1、SIGINT、SIGHUP、SIGHUP，看输出是不是和你预期的一样。

1.3.6. 关于信号的额外说明
查看 Go 中 Linux/amd64 信号的实现，发现大量使用的是 rt 相关系统调用，这是支持实时信号处理的 API。
C 语言中信号处理涉及到可重入函数和异步信号安全函数问题；Go 中不存在此问题。
Unix 和信号处理相关的很多系统调用，Go 都隐藏起来了，Go 中对信号的处理，signal 包中的函数基本就能搞定。



---


golang中对信号的处理主要使用os/signal包中的两个方法：一个是Notify方法用来监听收到的信号；一个是Stop方法用来取消监听。

`func Notify(c chan<- os.Signal, sig ...os.Signal)`

第一个参数表示接收信号的channel, 第二个及后面的参数表示设置要监听的信号，如果不设置表示监听所有的信号。
```go
package main
import (
    "fmt"
    "os"
    "os/signal"
)
func main() {
    c := make(chan os.Signal, 0)
    signal.Notify(c)
    s := <-c // 阻塞
    fmt.Println("Signal: ", s) 
}
```
结果分析：运行该程序，然后在终端中通过kill命令杀死对应的进程，便会得到结果

`func Stop(c chan<- os.Signal)`
```go
func main() {
    c := make(chan os.Signal, 0)
    signal.Notify(c)
 
    signal.Stop(c) //不允许继续往c中存入内容
    s := <-c       //c无内容，此处阻塞，所以不会执行下面的语句，也就没有输出
    fmt.Println("Got signal:", s)
}
```
由于signal存入channel中，所以可以利用channel特性，通过select针对不同的signal使得系统或者进程执行不同的操作。



---


Linux Signal及Golang中的信号处理
信号(Signal)是Linux, 类Unix和其它POSIX兼容的操作系统中用来进程间通讯的一种方式。一个信号就是一个异步的通知，发送给某个进程，或者同进程的某个线程，告诉它们某个事件发生了。
当信号发送到某个进程中时，操作系统会中断该进程的正常流程，并进入相应的信号处理函数执行操作，完成后再回到中断的地方继续执行。
如果目标进程先前注册了某个信号的处理程序(signal handler),则此处理程序会被调用，否则缺省的处理程序被调用。

发送信号
kill 系统调用(system call)可以用来发送一个特定的信号给进程。
kill 命令允许用户发送一个特定的信号给进程。
raise 库函数可以发送特定的信号给当前进程。

在Linux下运行man kill可以查看此命令的介绍和用法。

The command kill sends the specified signal to the specified process or process group. If no signal is specified, the TERM signal is sent. The TERM signal will kill processes which do not catch this signal. For other processes, it may be necessary to use the KILL (9) signal, since this signal cannot be caught.

Most modern shells have a builtin kill function, with a usage rather similar to that of the command described here. The '-a' and '-p' options, and the possibility to specify pids by command name is a local extension.

If sig is 0, then no signal is sent, but error checking is still performed.

一些异常比如除以0或者 segmentation violation 相应的会产生SIGFPE和SIGSEGV信号，缺省情况下导致core dump和程序退出。
内核在某些情况下发送信号，比如在进程往一个已经关闭的管道写数据时会产生SIGPIPE信号。
在进程的终端敲入特定的组合键也会导致系统发送某个特定的信号给此进程：

Ctrl-C 发送 INT signal (SIGINT)，通常导致进程结束
Ctrl-Z 发送 TSTP signal (SIGTSTP); 通常导致进程挂起(suspend)
Ctrl-\ 发送 QUIT signal (SIGQUIT); 通常导致进程结束 和 dump core.
Ctrl-T (不是所有的UNIX都支持) 发送INFO signal (SIGINFO); 导致操作系统显示此运行命令的信息
kill -9 pid 会发送 SIGKILL信号给进程。

处理信号
Signal handler可以通过signal()系统调用进行设置。如果没有设置，缺省的handler会被调用，当然进程也可以设置忽略此信号。
有两种信号不能被拦截和处理: SIGKILL和SIGSTOP。

当接收到信号时，进程会根据信号的响应动作执行相应的操作，信号的响应动作有以下几种：

中止进程(Term)
忽略信号(Ign)
中止进程并保存内存信息(Core)
停止进程(Stop)
继续运行进程(Cont)
用户可以通过signal或sigaction函数修改信号的响应动作（也就是常说的“注册信号”）。另外，在多线程中，各线程的信号响应动作都是相同的，不能对某个线程设置独立的响应动作。

信号类型
个平台的信号定义或许有些不同。下面列出了POSIX中定义的信号。
Linux 使用34-64信号用作实时系统中。
命令man 7 signal提供了官方的信号介绍。

在POSIX.1-1990标准中定义的信号列表

信号  值   动作  说明
SIGHUP  1   Term    终端控制进程结束(终端连接断开)
SIGINT  2   Term    用户发送INTR字符(Ctrl+C)触发
SIGQUIT 3   Core    用户发送QUIT字符(Ctrl+/)触发
SIGILL  4   Core    非法指令(程序错误、试图执行数据段、栈溢出等)
SIGABRT 6   Core    调用abort函数触发
SIGFPE  8   Core    算术运行错误(浮点运算错误、除数为零等)
SIGKILL 9   Term    无条件结束程序(不能被捕获、阻塞或忽略)
SIGSEGV 11  Core    无效内存引用(试图访问不属于自己的内存空间、对只读内存空间进行写操作)
SIGPIPE 13  Term    消息管道损坏(FIFO/Socket通信时，管道未打开而进行写操作)
SIGALRM 14  Term    时钟定时信号
SIGTERM 15  Term    结束程序(可以被捕获、阻塞或忽略)
SIGUSR1 30,10,16    Term    用户保留
SIGUSR2 31,12,17    Term    用户保留
SIGCHLD 20,17,18    Ign 子进程结束(由父进程接收)
SIGCONT 19,18,25    Cont    继续执行已经停止的进程(不能被阻塞)
SIGSTOP 17,19,23    Stop    停止进程(不能被捕获、阻塞或忽略)
SIGTSTP 18,20,24    Stop    停止进程(可以被捕获、阻塞或忽略)
SIGTTIN 21,21,26    Stop    后台程序从终端中读取数据时触发
SIGTTOU 22,22,27    Stop    后台程序向终端中写数据时触发
在SUSv2和POSIX.1-2001标准中的信号列表:

信号  值   动作  说明
SIGTRAP 5   Core    Trap指令触发(如断点，在调试器中使用)
SIGBUS  0,7,10  Core    非法地址(内存地址对齐错误)
SIGPOLL     Term    Pollable event (Sys V). Synonym for SIGIO
SIGPROF 27,27,29    Term    性能时钟信号(包含系统调用时间和进程占用CPU的时间)
SIGSYS  12,31,12    Core    无效的系统调用(SVr4)
SIGURG  16,23,21    Ign 有紧急数据到达Socket(4.2BSD)
SIGVTALRM   26,26,28    Term    虚拟时钟信号(进程占用CPU的时间)(4.2BSD)
SIGXCPU 24,24,30    Core    超过CPU时间资源限制(4.2BSD)
SIGXFSZ 25,25,31    Core    超过文件大小资源限制(4.2BSD)
Windows中没有SIGUSR1,可以用SIGBREAK或者SIGINT代替。

Go中的Signal发送和处理
有时候我们想在Go程序中处理Signal信号，比如收到SIGTERM信号后优雅的关闭程序(参看下一节的应用)。
Go信号通知机制可以通过往一个channel中发送os.Signal实现。
首先我们创建一个os.Signal channel，然后使用signal.Notify注册要接收的信号。

package main
import "fmt"
import "os"
import "os/signal"
import "syscall"
func main() {
    // Go signal notification works by sending `os.Signal`
    // values on a channel. We'll create a channel to
    // receive these notifications (we'll also make one to
    // notify us when the program can exit).
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)
    // `signal.Notify` registers the given channel to
    // receive notifications of the specified signals.
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    // This goroutine executes a blocking receive for
    // signals. When it gets one it'll print it out
    // and then notify the program that it can finish.
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()
    // The program will wait here until it gets the
    // expected signal (as indicated by the goroutine
    // above sending a value on `done`) and then exit.
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
go run main.go执行这个程序，敲入ctrl-C会发送SIGINT信号。 此程序接收到这个信号后会打印退出。

Go网络服务器如果无缝重启
Go很适合编写服务器端的网络程序。DevOps经常会遇到的一个情况是升级系统或者重新加载配置文件，在这种情况下我们需要重启此网络程序，如果网络程序暂停的时间较长，则给客户的感觉很不好。
如何实现优雅地重启一个Go网络程序呢。主要要解决两个问题：

进程重启不需要关闭监听的端口
既有请求应当完全处理或者超时
@humblehack 在他的文章Graceful Restart in Golang中提供了一种方式，而Florian von Bock根据此思路实现了一个框架endless。
此框架使用起来超级简单:

1
err := endless.ListenAndServe("localhost:4242", mux)
只需替换 http.ListenAndServe 和 http.ListenAndServeTLS。

它会监听这些信号： syscall.SIGHUP, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGINT, syscall.SIGTERM, 和 syscall.SIGTSTP。

此文章提到的思路是：

通过exec.Command fork一个新的进程，同时继承当前进程的打开的文件(输入输出，socket等)

file := netListener.File() // this returns a Dup()
path := "/path/to/executable"
args := []string{
    "-graceful"}
cmd := exec.Command(path, args...)
cmd.Stdout = os.Stdout
cmd.Stderr = os.Stderr
cmd.ExtraFiles = []*os.File{file}
err := cmd.Start()
if err != nil {
    log.Fatalf("gracefulRestart: Failed to launch, error: %v", err)
}
子进程初始化
网络程序的启动代码

server := &http.Server{Addr: "0.0.0.0:8888"}
 var gracefulChild bool
 var l net.Listever
 var err error
 flag.BoolVar(&gracefulChild, "graceful", false, "listen on fd open 3 (internal use only)")
 if gracefulChild {
     log.Print("main: Listening to existing file descriptor 3.")
     f := os.NewFile(3, "")
     l, err = net.FileListener(f)
 } else {
     log.Print("main: Listening on a new file descriptor.")
     l, err = net.Listen("tcp", server.Addr)
 }
父进程停止

if gracefulChild {
    parent := syscall.Getppid()
    log.Printf("main: Killing parent pid: %v", parent)
    syscall.Kill(parent, syscall.SIGTERM)
}
server.Serve(l)
同时他还提供的如何处理已经正在处理的请求。可以查看它的文章了解详细情况。

因此，处理特定的信号可以实现程序无缝的重启。



## net


## sort

## bytes

## encoding

## template

## reflect

---

LittelEndian


---

## go mod 包管理

#### 环境变量 GO111MODULE

GO111MODULE 有三个值：off, on和auto（默认值）。

- GO111MODULE=off，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。
- GO111MODULE=on，go命令行会使用modules，而一点也不会去GOPATH目录下查找。
- GO111MODULE=auto，默认值，go命令行将会根据当前目录来决定是否启用module功能。

当modules 功能启用时，依赖包的存放位置变更为$GOPATH/pkg，允许同一个package多个版本并存，且多个项目可以共享缓存的 module。

#### go mod
golang 提供了 go mod命令来管理包。

go mod 有以下命令：

|命令   |说明   |
| ------ | ------ |
|download | download modules to local cache(下载依赖包) |
|edit     | edit go.mod from tools or scripts（编辑go.mod |
|graph    | print module requirement graph (打印模块依赖图) |
|init     | initialize new module in current directory（在当前目录初始化mod）|
|tidy     | add missing and remove unused modules(拉取缺少的模块，移除不用的模块) |
|vendor   | make vendored copy of dependencies(将依赖复制到vendor下) |
|verify   | verify dependencies have expected content (验证依赖是否正确）|
|why      |explain why packages or modules are needed(解释为什么需要依赖)|



---

## gRPC
RPC is a acronym for Remote Procedure Call

### Types of gRPC applications

gRPC applications can be written using 3 types of processing, as follows:

Unary RPCs where the client sends a single request to the server and gets a single response back, just like a normal function call.

`rpc SayHello(HelloRequest) returns (HelloResponse);`

Server streaming RPCs where the client sends a request to the server and gets a stream to read a sequence of messages back. The client reads from the returned stream until there are no more messages. gRPC guarantees message ordering within an individual RPC call.

`rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);`

Client streaming RPCs where the client writes a sequence of messages and sends them to the server, again using a provided stream. Once the client has finished writing the messages, it waits for the server to read them and return its response. Again gRPC guarantees message ordering within an individual RPC call.

`rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);`

Bidirectional streaming RPCs where both sides send a sequence of messages using a read-write stream. The two streams operate independently, so clients and servers can read and write in whatever order they like: for example, the server could wait to receive all the client messages before writing its responses, or it could alternately read a message then write a message, or some other combination of reads and writes. The order of messages in each stream is preserved.

`rpc BidiHello(stream HelloRequest) returns (stream HelloResponse);`

Synch vs. Asynch
As the name implies, synchronous processing occurs when we have a communication where the client thread is blocked when a message is sent and is been processed.

Asynchronous processing occurs when we have this communication with the processing been done by other threads, making the whole process been non-blocking.

On gRPC we have both styles of processing supported, so it is up to the developer to use the better approach for his solutions.

Deadlines & timeouts
A deadline stipulates how much time a gRPC client will wait on a RPC call to return before assuming a problem has happened. On the server’s side, the gRPC services can query this time, verifying how much time it has left.

If a response couldn’t be received before the deadline is met, a DEADLINE_EXCEEDED error is thrown and the RPC call is terminated.

RPC termination
On gRPC, both clients and servers decide if a RPC call is finished or not locally and independently. This means that a server can decide to end a call before a client has transmitted all their messages and a client can decide to end a call before a server has transmitted one or all of their responses.

This point is important to remember specially when working with streams, in a sense that logic must pay attention to possible RPC terminations when treating sequences of messages.

Channels
Channels are the way a client stub can connect with gRPC services on a given host and port. Channels can be configured specific by client, such as turning message compression on and off for a specific client.

Metadata
Metadata is information about a particular RPC call (such as authentication details) in the form of a list of key-value pairs, where the keys are strings and the values are typically strings, but can be binary data. Metadata is opaque to gRPC itself - it lets the client provide information associated with the call to the server and vice versa.

Access to metadata is language dependent.

### Protocol Buffer

The first step when working with protocol buffers is to define the structure for the data you want to serialize in a proto file: this is an ordinary text file with a .proto extension. Protocol buffer data is structured as messages, where each message is a small logical record of information containing a series of name-value pairs called fields.

Then, once you’ve specified your data structures, you use the protocol buffer compiler protoc to generate data access classes in your preferred language(s) from your proto definition.

