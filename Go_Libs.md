# Go Libs

## fmt
fmt库 fmt.Stringer接口？？？

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
这是定时器相关类型。本章最后会讨论定时器。

1.5. Weekday 和 Month
这两个类型的原始类型都是 int，定义它们，语义更明确，同时，实现 fmt.Stringer 接口，方便输出。

## os
os.Stdin
os.Stderr??

## flag

## net

## text?
text/tabwriter

## io

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

缓冲区io，标准库中bufio包支持缓冲区io操作，可以轻松处理文本内容。例如：bufio.Scanner。

ioutil，io包下面的一个子包ioutil封装了一些非常方便的功能，例如，使用函数ReadFile将文件内容加载到[]byte中。ioutil.ReadFile和ioutil.WriteFile都使用*os.File的Read和Write方法。


### sort





