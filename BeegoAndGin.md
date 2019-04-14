# Beego & Gin

### MVC
- M: Model, 数据库中的记录转换成编程语言中的对象。
- V: View, 模板，展现结果的方式。
- C: Controller, M和V之间的纽带。

# Go 并发

- 基于线程和进程（Apache）,每来一个请求就分配一个进程或线程，每个线程服务一个用户。C10K问题：10000个线程。
- 异步非阻塞（NginX, NodeJS）。
- 协程（Golang, Lua), 底层还是线程。可以理解成轻量的线程。

##### 并发与并行
- 并发（Concurrency）：系统通过调度，来回切换交替执行多个任务，看起来像同时进行。
- 并行（Parallelism）：真正的同时进行。

##### Golang面向对象

```go
//“面向对象”
type Foo struct {
    name string
}

func (f *Foo) echo() {
    fmt.Println(f.name)
}

//“继承”
type Bar struct {
    Foo
}

func main() {
    b := Bar{Foo{name: "A"}}
    b.echo()
}

//“多态”
type Foo interface {
    qux()
}

type A struct{}
type B struct{}

func (a *A) qux() 
func (b *B) qux()

func main() {
    var f Foo
    f = A{}
    f = B{}
}

```

