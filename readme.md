# Go 笔记

- 类型检查：编译时
- 运行环境：编译成机器码直接运行
- 编程范式：面向借口，函数式编程，并发编程

- 采用CSP (Communication Sequential Process)模型
- 不需要锁，不需要callback.

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



## Project List: 

### IO
简单的IO练习

### Pipeline: 外部排序Pipeline
选自慕课网：搭建并行处理管道

- 原始数据过大，无法一次读入内存，所以分块读入内存。每个块数据进行内部排序（直接调用API排序），最后讲各个节点归并，归并选择两两归并。

