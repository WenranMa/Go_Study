# interface，reflect 相关
## 接口
接口类型是一种抽象的类型。它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合。它们只会展示出它们自己的方法。也就是说当你有看到一个接口类型的值时，不知道它是什么，但知道可以通过它的方法来做什么。

`fmt.Printf`会把结果写到标准输出，这个函数都使用了另一个函数`fmt.Fprintf`。
```go
package fmt

func Fprintf(w io.Writer, format string, a ...any) (n int, err error) 

func Printf(format string, a ...any) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}
```
在Printf函数中的第一个参数os.Stdout是*os.File类型，Fprintf函数中的第一个参数也不是一个文件类型。它是io.Writer类型，这是一个接口类型
```go
package io
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
\*os.File和\*bytes.Buffer都实现了io.Writer接口。所以可以用作Fprintf输入。

接口类型具体描述了一系列方法的集合，一个实现了这些方法的具体类型是这个接口类型的实例。

### 实现接口的条件
一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口。接口指定的规则非常简单:表达一个类型属于某个接口只要这个类型实现这个接口。

例如，\*os.File类 型实现了io.Reader，Writer，Closer和ReadWriter接口。\*bytes.Buffer实现了Reader，Writer和ReadWriter这些接口，但是它没有实现Closer接口因为它不具有Close方法。
```go
    var w io.Writer
    w = os.Stdout         //OK: *os.File has Write method
    w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
    w = time.Second       // compile error: time.Duration lacks Write method

    var rwc io.ReadWriteCloser //
    rwc = os.Stdout            // OK: *os.File has Read, Write, Close methods
    rwc = new(bytes.Buffer)    //compile error: *bytes.Buffer lacks Close method

    w = rwc // OK: io.ReadWriteCloser has Write method
    rwc = w // compile error: io.Writer lacks Close method
```

对于每一个命名过的具体类型T;它一些方法的接收者是类型T本身而另一些则是一个T指针。`T类型变量上调用\*T的方法是合法的，但必须是变量。编译器隐式的获取了它的地址。但这仅仅是一个语法糖:T类型的值不拥有*T指针的方法。`

```go
type IntSet struct { /* ... */ }
func (*IntSet) String() string
var _ = IntSet{}.String() // compile error: String requires *IntSet receiver
```
但是我们可以在一个IntSet值上调用这个方法:
```go
var s IntSet
var _ = s.String() // OK: s is a variable and &s has a String method
```
但由于只有*IntSet类型有String方法，也只有*IntSet类型实现了fmt.Stringer接口
```go
var _ fmt.Stringer = &s // OK
var _ fmt.Stringer = s // compile error: IntSet lacks String method
```

### flag.Value接口，fmt.Stringer接口
```go
func main() {
    var period = flag.Duration("period", 1*time.Second, "sleep period")
    flag.Parse()
    fmt.Printf("Sleeping for %v...", *period)
    time.Sleep(*period)
    fmt.Println("Done!")
}
// Sleeping for 5s...Done!
```
可以通过`-period`这个命令行标记来控制：`go run main.go -period 5s`。

为自定义数据类型定义新的标记符号，需要定义一个实现`flag.Value`接口的类型。
```go
package flag
// Value is the interface to the value stored in a flag.
type Value interface {
    String() string
    Set(string) error
}
```
String方法格式化标记的值用在命令行帮组消息中，每一个flag.Value也是一个fmt.Stringer。Set方法解析它的字符串参数并且更新标记变量的值。实际上，Set方法和String是两个相反的操作。

fmt.Stringer接口定义了String()方法，实现了fmt.Stringer接口的类型，会在被格式化打印时，调用给类型的String()方法。

举例：下面代码输出什么？
```go
type ConfigOne struct {
    Daemon string
}

func (c *ConfigOne) String() string {
    return fmt.Sprintf("print: %v", c)
}

func main() {
    c := &ConfigOne{}
    c.String()
}

/*
如果类型实现 String() 方法，当格式化输出时会自动使用 String() 方法。上面这段代码是在该类型的 String() 方法内使用格式化输出，导致递归调用。

如果go build会报错：
fmt.Sprintf format %v with arg c causes recursive (*play.ConfigOne).String method call

如果go run main.go 会运行时错误：
runtime: goroutine stack exceeds 1000000000-byte limit
fatal error: stack overflow
*/
```

再比如：
```go
type Orange struct {
    Quantity int
}

func (o *Orange) Increase(n int) {
    o.Quantity += n
}

func (o *Orange) Decrease(n int) {
    o.Quantity -= n
}

func (o *Orange) String() string {
    return fmt.Sprintf("%#v", o.Quantity)
}

func main() {
    var orange Orange
    orange.Increase(10)
    orange.Decrease(5)
    fmt.Println(orange)
}
/*
输出：{5}

String() 是指针方法，而不是值方法，所以使用 Println() 输出时不会调用到 String() 方法。

可以这样修复：
func main() {
    orange := &Orange{}
    orange.Increase(10)
    orange.Decrease(5)
    fmt.Println(orange)
}
这样输出就是：5

如果String()方法为：
func (o *Orange) String() string {
    return fmt.Sprintf("%v", o)
}
则会循环输出，就像第一个例子。

但 fmt.Sprintf("%#v", o) 会输出 &main.Orange{Quantity:5}

格式化打印中：
p := point{1, 2}
fmt.Printf("%v\n", p)  //输出结果为 {1 2}
fmt.Printf("%+v\n", p) //输出结果为 {x:1 y:2}，%+v 变体将包括结构体的字段名。
fmt.Printf("%#v\n", p) //输出结果为 main.point{x:1, y:2}， %#v 变体打印值的 Go 语法表示，即将生成该值的源代码片段。
```

### 接口值
一个`接口的值`由两个部分组成，一个具体的`类型`和那个`类型的值`。它们被称为`接口的动态类型和动态值`。Go语言这种静态类型的语言，类型是编译期的概念；因此一个类型不是一个值。

`var w io.Writer` 在Go语言中，变量总是被一个定义明确的值初始化，即使接口类型也不例外。对于一个`接口的零值就是它的类型和值的部分都是nil`。

`w = os.Stdout` 这个赋值过程调用了一个`具体类型到接口类型的隐式转换`，这和显式的使用`io.Writer(os.Stdout)`是等价的。这个接口值的动态类型被设为\*os.File指针，它的动态值持有os.Stdout的拷贝。(os.Stdout是一个\*os.File全局变量，所以它的值是固定的，所以它的拷贝也是固定的。)

接口值可以使用==和!=来进行比较。两个接口值相等仅当它们`都是nil值`或者`它们的动态类型相同并且动态值也根据这个动态类型的==操作相等`。可以用在map的键或者作为switch语句的操作数。如果两个接口值的动态类型相同，但这个动态类型是不可比较(比如切片)，将它们进行比较就会失败并且panic:
```go
    var x interface{} = []int{1, 2, 3}
    fmt.Printf("%T\n", x) // []int
    fmt.Println(x == x) // panic: comparing uncomparable type []int
```
使用fmt包的`%T`动作可以获得接口值动态类型。

`一个不包含任何值的nil接口值和一个刚好包含nil指针的接口值是不同的。`
```go
func main() {
    var buf *bytes.Buffer
    f(buf) // NOTE: subtly incorrect! //panic.
}
func f(out io.Writer) {
    if out != nil {
        out.Write([]byte("done!\n")) 
    }
}
```
当main函数调用函数f时，它给f函数的out参数赋了一个\*bytes.Buffer的空指针，所以out的动态值是nil。然而，它的动态类型是*bytes.Buffer，意思就是out变量是一个包含空指针值的非空接口，所以防御性检查out!=nil的结果依然是true。解决方案就是将main函数中的变量buf的类型改为io.Writer。

在Go语言中，一个接口类型总是代表着某一种类型（即所有实现它的类型）的行为。一个接口类型的声明通常会包含关键字type、类型名称、关键字interface以及由花括号包裹的若干方法声明。示例如下：
```go
type Animal interface {
    Move(int) bool
}

type Person struct {
    Name string
}

func (p *Person) Move(dis int) bool {
	fmt.Println("Person Move: ", dis)
    return true
}
```
注意，接口类型中的方法声明是普通的方法声明的简化形式。它们只包括方法名称、参数声明列表和结果声明列表。其中的参数的名称和结果的名称都可以被省略。不过，出于文档化的目的，我还是建议大家在这里写上它们。因此，Move方法的声明至少应该是这样的：`Move(dis int) (check bool)`。

如果一个数据类型所拥有的方法集合中包含了某一个接口类型中的所有方法声明的实现，那么就可以说这个数据类型实现了那个接口类型。所谓实现一个接口中的方法是指，具有与该方法相同的声明并且添加了实现部分（由花括号包裹的若干条语句）。相同的方法声明意味着完全一致的名称、参数类型列表和结果类型列表。其中，参数类型列表即为参数声明列表中除去参数名称的部分。一致的参数类型列表意味着其长度以及顺序的完全相同。对于结果类型列表也是如此。

我们无需在一个数据类型中声明它实现了哪个接口。只要满足了“方法集合为其超集”的条件，就建立了“实现”关系。这是典型的无侵入式的接口实现方法。

### 断言和类型转换
如何判断一个数据类型是否实现了某个接口类型，这需要用到Go语言中的类型断言。类型断言是一个`使用在接口值上的操作`。`x.(T)`，一个类型断言检查x对象的类型是否和断言的类型T匹配。`不能在一个非接口类型的值上应用类型断言来判定它是否属于某一个接口类型的。我们必须先把前者转换成空接口类型的值`。这又涉及到了Go语言的类型转换。

Go语言的类型转换规则定义了是否能够以及怎样可以把一个类型的值转换另一个类型的值。另一方面，所谓`空接口类型即是不包含任何方法声明的接口类型，用interface{}表示，常简称为空接口`。`interface{}`被称为`空接口类型`，因为空接口类型对实现它的类型没有要求，所以我们可以将任意一个值赋给空接口类型。`var any interface{}`。Go语言中的包含预定义的任何数据类型都可以被看做是空接口的实现。我们可以直接使用类型转换表达式把一个*Person类型转换成空接口类型的值，就像这样：
```go
p := Person{"Robert"}
v := interface{}(&p)

h, ok := v.(Animal)
```
请注意第二行。在类型字面量后跟由圆括号包裹的值（或能够代表它的变量、常量或表达式）就构成了一个类型转换表达式，意为将后者转换为前者类型的值。在这里，我们把表达式&p的求值结果转换成了一个空接口类型的值，并由变量v代表。注意，表达式&p（&是取址操作符）的求值结果是一个*Person类型的值，即p的指针。在这之后，我们就可以在v上应用类型断言。

`如果T是一个具体类型，类型断言检查x的动态类型是否和T相同`。如果这个检查成功，类型断言的结果是x的动态值。如果检查失败，接下来这个操作会抛出panic。`T也可以是一个接口类型，类型断言检查x的动态类型是否满足T（实现了T的方法）。`

```go
    var w io.Writer
    w = os.Stdout
    f := w.(*os.File) // success: f == os.Stdout
    c := w.(*bytes.Buffer) //panic: interface holds *os.File, not *bytes.Buffer
    d := w.(sort.Interface) //panic: interface conversion: *os.File is not sort.Interface: missing method Len
```
不是interface类型做类型断言都是回报non-interface的错误的，
```go
    s := "hello world"
    if v, ok := s.(string); ok {
        fmt.Println(v)
    }
    //invalid type assertion: s.(string) (non-interface type string on left)
    //所以我们只能通过将s作为一个interface{}的方法来进行类型断言，如下代码所示：
    x := "hello world"
    if v, ok := interface{}(x).(string); ok { // interface{}(x):把 x 的类型转换成 interface{}
        fmt.Println(v)
    }
```
断言可以返回两个值，第二个结果常规地赋值给一个命名为ok的变量。如果这个操作失败了，那么ok就是false值，第一个结果等于被断言类型的零值。

`断言操作的对象是一个nil接口值，断言都会失败。`

类型断言表达式v.(Animal)的求值结果可以有两个。第一个结果是被转换后的那个目标类型（这里是Animal）的值，而第二个结果则是转换操作成功与否的标志。显然，ok代表了一个bool类型的值。它也是这里判定实现关系的重要依据。

我们在讲接口的时候说过，如果一个数据类型所拥有的方法集合中包含了某一个接口类型中的所有方法声明的实现，那么就可以说这个数据类型实现了那个接口类型。要获知一个数据类型都包含哪些方法并不难。但是要注意指针方法与值方法的区别。

拥有指针方法Move的指针类型*Person是接口类型Animal的实现类型，但是它的基底类型Person却不是。这样的表象隐藏着另一条规则：`一个指针类型拥有以它以及以它的基底类型为接收者类型的所有方法，而它的基底类型却只拥有以它本身为接收者类型的方法。`

以Person类型为例。即使我们把Move改为值方法，*Person类型也仍会是Animal接口的实现类型。另一方面，如果Move是指针方法，Person类型就不可能是Animal接口的实现类型。

另外，还有一点需要大家注意，我们在`基底类型的值上仍然可以调用它的指针方法`。例如，若我们有一个Person类型的变量bp，则调用表达式bp.Grow()是合法的。这是因为，如果Go语言发现我们调用的Grow方法是bp的指针方法，那么它会把该调用表达式视为(&bp).Grow()。实际上，这时的bp.Grow()是(&bp).Grow()的速记法。

```go
package main

import "fmt"

type Pet interface {
    Name() string
    Age() uint8
}

type Dog struct {
    name string
    age uint8
}

func (d Dog) Name() string {
    return d.name
}

func (d Dog) Age() uint8 {
    return d.age
}

func main() {
    myDog := Dog{"Little D", 3}
    _, ok1 := interface{}(&myDog).(Pet)
    _, ok2 := interface{}(myDog).(Pet)
    fmt.Printf("%v, %v\n", ok1, ok2)  //true, true
    //如果其中一个方法 或 两个方法都是指针方法，则ok2 == false
}
```

### sort.Interface接口
Go语言的sort.Sort函数不会对具体的序列和它的元素做任何假设。它使用了一个接口类型sort.Interface来指定通用的排序算法。序列的表示经常是一个切片。

一个内置的排序算法需要知道三个东西:序列的长度，表示两个元素比较的结果，一种交换两个元素的方式;这就是sort.Interface的三个方法:
```go
package sort
type Interface interface { 
    Len() int
    Less(i, j int) bool // i, j are indices of sequence elements
    Swap(i, j int) }
```
为了对序列进行排序，我们需要定义一个实现了这三个方法的类型，然后对这个类型的一个实例应用sort.Sort函数。

例：B_07_Sort

### http.Handler接口
```go
package http
type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
func ListenAndServe(address string, h Handler) error
```
ListenAndServe函数需要一个例如“localhost:8000”的服务器地址，和一个所有请求都可以分派的Handler接口实例。它会一直运行，直到这个服务因为一个错误而失败，返回值一定是一个非空的错误。

例：B_08_Http_01 可以在浏览器输入：`http://localhost:8000/price?item=socks` 。

net/http包提供了一个请求多路器ServeMux来简化URL和handlers的联系。一个ServeMux将一批http.Handler聚集到一个单一的http.Handler中。

例：B_09_Http_02

### error接口
error接口类型有一个返回错误信息的单一方法:
```go
package errors

// error interface不在erros package下，而在buildin package下。
type error interface {
    Error() string
}

func New(text string) error {
    return &errorString{text}
}

type errorString struct{ 
    text string 
}

func (e *errorString) Error() string {
    return e.text
}
```
创建一个error最简单的方法就是调用errors.New函数，它会根据传入的错误信息返回一个新的error。

errorString是一个结构体而非一个字符串。并且因为是指针类型*errorString而不是errorString类型满足error接口，所以每个New函数的调用都分配了一个单独地址，即使错误信息相同，实例也不相等：`fmt.Println(errors.New("EOF") == errors.New("EOF")) // "false"`。

有一个方便的封装函数fmt.Errorf，它还会处理字符串格式化：
```go
package fmt

import "errors"

func Errorf(format string, args ...interface{}) error {
    return errors.New(Sprintf(format, args...))
}
```

### 示例：表达式请求

### 基于类型断言区别错误类型

### 通过类型断言询问行为

### 类型开关

### 示例：基于标记的XML解码

---
## 反射
有时候我们需要编写一个函数能够处理一类并不满足普通公共接口的类型的值。一个大家熟悉的例子是fmt.Fprintf函数提供的字符串格式化处理逻辑，它可以用来对任意类型的值格式化并打印，甚至支持用户自定义的类型。

反射可以检查未知类型的表示方式。

### refect.Type和reflect.Value
反射是由reflect包提供的。它定义了两个重要的类型，Type和Value。Type是接口，表示一个Go类型。函数`reflect.TypeOf接受任意的interface{}类型，并以reflect.Type形式返回其动态类型。`
```go
    t := reflect.TypeOf(3)  // a reflect.Type //一个隐式的接口转换操作
    fmt.Println(t.String()) // "int"
    fmt.Println(t)          // "int"
```

`reflect.TypeOf总是返回具体的类型。`

```go
var w io.Writer = os.Stdout
fmt.Println(reflect.TypeOf(w)) // 打印 "*os.File" 而不是 "io.Writer"。
```
reflect.Type接口是满足fmt.Stringer接口的，fmt.Printf提供了一个缩写`%T`参数, 内部使用 reflect.TypeOf来输出:`fmt.Printf("%T\n", 3) // "int"`。

reflect包中另一个类型是Value，可以装载任意类型的值。`函数reflect.ValueOf接受任意的interface{}类型，并返回一个装载着其动态值的reflect.Value。`和reflect.TypeOf类似，reflect.ValueOf返回的结果也是具体的类型，但是reflect.Value也可以持有一个接口值。
reflect.Value也满足fmt.Stringer接口, fmt包的`%v`标志参数会对reflect.Values特殊处理。对 Value 调用 Type 方法将返回具体类型所对应的 reflect.Type。

```go
    v := reflect.ValueOf(3) // a reflect.Value
    fmt.Println(v)          // "3"
    fmt.Printf("%v\n", v)   // "3"
    fmt.Println(v.String()) // NOTE: "<int Value>"
    t := v.Type()           // a reflect.Type
    fmt.Println(t.String()) // "int"
```

reflect.Value的Kind方法可以用来枚举类型，kinds类型是有限的: Bool, String和所有数字类型的基础类型; Array和Struct对应的聚合类型; Chan, Func, Ptr, Slice和Map对应的引用类型; interface类型; 还有表示空值的Invalid类型(空的reflect.Value的kind即为Invalid)。

例：B_24_Format

### Display,递归打印器

### 示例：编码为S表达式

### 通过reflect.Value修改值

### 示例：解码S表达式

### 获取结构体字段标识

### 显示一个类型的方法集

--- 
## 练习：

下面这段代码能否编译通过？如果可以，输出什么？
```go
func GetValue() int {
    return 1
}

func main() {
    i := GetValue()
    switch i.(type) {
    case int:
        fmt.Println("int")
    case string:
        fmt.Println("string")
    case interface{}:
        fmt.Println("interface")
    default:
        fmt.Println("unknown")
    }
}
/*
编译失败
i (variable of type int) is not an interface

只有接口类型才能使用类型选择，类型选择的语法形如：i.(type)，其中i是接口，type是固定关键字。
*/
```

下面这段代码输出什么？
```go
func main() {  
    var i interface{}
    if i == nil {
        fmt.Println("nil")
        return
    }
    fmt.Println("not nil")
}
/*
输出：nil

当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil
*/
```

下面代码是否能编译通过？如果通过，输出什么？
```go
func Foo(x interface{}) {
    if x == nil {
        fmt.Println("empty interface")
        return
    }
    fmt.Println("non-empty interface")
}
func main() {
    var x *int = nil
    Foo(x)
}
/*
输出 non-empty interface**

接口除了有静态类型，还有动态类型和动态值，当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。这里的 x 的动态类型是 `*int`，所以 x 不为 nil。
*/
```

下面这段代码输出什么？
```go
type A interface {
    ShowA() int
}

type B interface {
    ShowB() int
}

type Work struct {
    i int
}

func (w Work) ShowA() int {
    return w.i + 10
}

func (w Work) ShowB() int {
    return w.i + 20
}

func main() {
    c := Work{3}
    var a A = c
    var b B = c
    fmt.Println(a.ShowA())
    fmt.Println(b.ShowB())
}
/*
输出 13 23

一种类型实现多个接口，结构体 Work 分别实现了接口 A、B，所以接口变量 a、b 调用各自的方法 ShowA() 和 ShowB()，输出 13、23。
*/
```

下面代码输出什么？
```go
type A interface {
    ShowA() int
}

type B interface {
    ShowB() int
}

type Work struct {
    i int
}

func (w Work) ShowA() int {
    return w.i + 10
}

func (w Work) ShowB() int {
    return w.i + 20
}

func main() {
    var a A = Work{3}
    s := a.(Work)
    fmt.Println(s.ShowA())
    fmt.Println(s.ShowB())
}
/* 
输出 13 23

类型断言，并且断言成功，返回Work的值。
```

下面代码输出什么？
```go
type A interface {
    ShowA() int
}

type B interface {
    ShowB() int
}

type Work struct {
    i int
}

func (w Work) ShowA() int {
    return w.i + 10
}

func (w Work) ShowB() int {
    return w.i + 20
}

func main() {
    c := Work{3}
    var a A = c
    var b B = c
    fmt.Println(a.ShowB())
    fmt.Println(b.ShowA())
}
/*
输出 compilation error
a.ShowB undefined (type A has no field or method ShowB)
b.ShowA undefined (type B has no field or method ShowA)

知识点：接口的静态类型。

a、b 具有相同的动态类型和动态值，分别是结构体 Work 和 {3}；a 的静态类型是 A，b 的静态类型是 B，接口 A 不包括方法 ShowB()，接口 B 也不包括方法 ShowA().
*/
```

下面这段代码输出什么？
```go
type People struct{}

func (p *People) ShowA() {
    fmt.Println("showA")
    p.ShowB()
}
func (p *People) ShowB() {
    fmt.Println("showB")
}

type Teacher struct {
    People
}

func (t *Teacher) ShowB() {
    fmt.Println("teacher showB")
}

func main() {
    t := Teacher{}
    t.ShowB()
}
// teacher showB
// 在嵌套结构体中，People 称为内部类型，Teacher 称为外部类型；通过嵌套，内部类型的属性、方法，可以为外部类型所有，就好像是外部类型自己的一样。此外，外部类型还可以定义自己的属性和方法，甚至可以定义与内部相同的方法，这样内部类型的方法就会被“屏蔽”。这个例子中的 ShowB() 就是同名方法。
```

下面这段代码输出什么？
```go
type People struct{}

func (p *People) ShowA() {
    fmt.Println("showA")
    p.ShowB()
}
func (p *People) ShowB() {
    fmt.Println("showB")
}

type Teacher struct {
    People
}

func (t *Teacher) ShowB() {
    fmt.Println("teacher showB")
}

func main() {
    t := Teacher{}
    t.ShowA()
}
// showA
// showB
// Teacher 没有自己 ShowA()，所以调用内部类型 People 的同名方法，需要注意的是第 5 行代码调用的是 People 自己的 ShowB 方法。
```

下面这段代码输出什么？
```go
type People interface {
    Show()
}

type Student struct{}

func (stu *Student) Show() {
}

func main() {
    var s *Student
    if s == nil {
        fmt.Println("s is nil")
    } else {
        fmt.Println("s is not nil")
    }
    var p People = s
    if p == nil {
        fmt.Println("p is nil")
    } else {
        fmt.Println("p is not nil")
    }
}
/*
输出：s is nil 和 p is not nil

当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。
第一个判断，s是一个nil指针，给变量 p 赋值之后，p 的动态值是 nil，但是动态类型却是 *Student，是一个 nil 指针，所以相等条件不成立。
*/
```

下面代码输出什么？
```go
func main() {
    x := interface{}(nil)
    y := (*int)(nil)
    a := y == x
    b := y == nil
    _, c := x.(interface{})
    println(a, b, c)
}
/*
输出： false true false

类型断言语法：i.(Type)，其中 i 是接口，Type 是类型或接口。编译时会自动检测 i 的动态类型与 Type 是否一致。但是，如果动态类型不存在，则断言总是失败。
*/
```

下面的代码有什么问题，请说明。
```go
type data struct {
    name string
}

func (p *data) print() {
    fmt.Println("name:", p.name)
}

type printer interface {
    print()
}

func main() {
    d1 := data{"one"}
    d1.print()

    var in printer = data{"two"}
    in.print()
}
/*
cannot use data literal (type data) as type printer in assignment:
data does not implement printer (print method has pointer receiver)

结构体类型 data 没有实现接口 printer。而是*data实现了接口。
```

interface{} 是可以指向任意对象的 Any 类型，是否正确？ -- 正确

下面的代码有什么问题？
```go
func main() {  
    var x = nil 
    _ = x
}
/*
use of untyped nil in variable declaration

nil 用于表示 interface、函数、maps、slices 和 channels 的“零值”。如果不指定变量的类型，编译器猜不出变量的具体类型，导致编译错误。

修复代码：
*/
func main() {
    var x interface{} = nil
    _ = x
}
```

A、B、C、D 哪些选项有语法错误？
```go
type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
    s := S{}
    p := &s
    f(s) //A
    g(s) //B
    f(p) //C
    g(p) //D
}
/*
B D

cannot use s (variable of type S) as *interface{} value in argument to g: S does not implement *interface{} (type *interface{} is pointer to interface, not interface)
cannot use p (variable of type *S) as *interface{} value in argument to g: *S does not implement *interface{} (type *interface{} is pointer to interface, not interface)

函数参数为 interface{} 时可以接收任何类型的参数，包括用户自定义类型等，即使是接收指针类型也用 interface{}，而不是使用 *interface{}。
*/
```

下面哪一行代码会 panic，请说明原因？
```go
package main

func main() {
  var x interface{}
  var y interface{} = []int{3, 5}
  _ = x == x
  _ = x == y
  _ = y == y
}
/*
第 8 行 _ = y == y
comparing uncomparable type []int 因为两个比较值的动态类型为同一个不可比较类型。
```

下面这段代码输出什么？
```go
type S1 struct{}

func (s1 S1) f() {
    fmt.Println("S1.f()")
}
func (s1 S1) g() {
    fmt.Println("S1.g()")
}

type S2 struct {
    S1
}

func (s2 S2) f() {
    fmt.Println("S2.f()")
}

type I interface {
    f()
}

func printType(i I) {
    fmt.Printf("%T\n", i)
    if s1, ok := i.(S1); ok {
        s1.f()
        s1.g()
    }
    if s2, ok := i.(S2); ok {
        s2.f()
        s2.g()
    }
}

func main() {
    printType(S1{})
    printType(S2{})
}
// main.S1
// S1.f()
// S1.g()
// main.S2
// S2.f()
// S1.g()

// 知识点：类型断言，结构体嵌套。
// 结构体 S2 嵌套了结构体 S1，S2 自己没有实现 g() ，调用的是 S1 的 g()。
```

下面这段代码能否通过编译？如果通过，输出什么？
```go
package main

import "fmt"

type MyInt1 int
type MyInt2 = int

func main() {
    var i int = 0
    var i1 MyInt1 = i
    var i2 MyInt2 = i
    fmt.Println(i1, i2)
}
// `cannot use i (variable of type int) as MyInt1 value in variable declaration`

// 这道题考的是 `类型别名` 与 `类型定义` 的区别
// 第5行代码是基于类型 `int` 创建了新类型 `MyInt1` ，第6行代码是创建了int的类型别名 `MyInt2` ，注意类型别名的定义是 `=` 。所以，第10行代码相当于是将int类型的变量赋值给MyInt1类型的变量，Go是强类型语言，编译当然不通过；而MyInt2只是int的别名，本质上还是int，可以赋值。
// 第10行代码的赋值可以使用强制类型转换 `var i1 MyInt1 = MyInt1(i)` 
```

关于类型转化，下面选项正确的是？-- TBD
```go
//A.
type MyInt int
var i int = 1
var j MyInt = i

//B.
type MyInt int
var i int = 1
var j MyInt = (MyInt)i

//C.
type MyInt int
var i int = 1
var j MyInt = MyInt(i)

//D.
type MyInt int
var i int = 1
var j MyInt = i.(MyInt)

//**答：C**
```

下面赋值正确的是（）答：B、D

- A. var x = nil
- B. var x interface{} = nil
- C. var x string = nil
- D. var x error = nil

A错在没有写类型，C错在字符串的空值是 `""` 而不是nil。
知识点：nil只能赋值给指针、chan、func、interface、map、或slice、类型的变量。

## 一个nil与interface{}比较的问题。
Go 语言中有两种略微不同的接口，`一种是带有一组方法的接口，另一种是不带任何方法的空接口 interface{}。` Go 语言使用runtime.iface表示带方法的接口，使用runtime.eface表示不带任何方法的空接口interface{}。

一个 interface{} 类型的变量包含了 2 个指针，一个指针指向值的类型，另外一个指针指向实际的值。在 Go 源码中 runtime 包下，我们可以找到runtime.eface的定义。
```go
type eface struct {
	_type *_type
	data  unsafe.Pointer
}
```
从空接口的定义可以看到，当一个空接口变量为 nil 时，需要其两个指针均为 0 才行。

回到最初的问题，我们打印下传入函数中的空接口变量值，来看看它两个指针值的情况。

```go
// InterfaceStruct 定义了一个 interface{} 的内部结构
type InterfaceStruct struct {
	pt uintptr // 到值类型的指针
	pv uintptr // 到值内容的指针
}

// ToInterfaceStruct 将一个 interface{} 转换为 InterfaceStruct
func ToInterfaceStruct(i interface{}) InterfaceStruct {
	return *(*InterfaceStruct)(unsafe.Pointer(&i))
}

func IsNil(i interface{}) {
	fmt.Printf("i value is %+v\n", ToInterfaceStruct(i))
}

func main() {
	var sl []string
	IsNil(sl)
	IsNil(nil)
}
// 运行输出：
// i value is {pt:6769760 pv:824635080536}
// i value is {pt:0 pv:0}
```
可见，虽然 sl 是 nil 切片，但是其本上是一个类型为 []string，值为空结构体 slice 的一个变量，所以 sl 传给空接口时是一个非 nil 变量。

再细究的话，你可能会问，既然 sl 是一个有类型有值的切片，为什么又是个 nil。针对具体类型的变量，判断是否是 nil 要根据其值是否为零值。因为 sl 一个切片类型，而切片类型的定义在源码包src/runtime/slice.go我们可以找到。

```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```
我们继续看一下值为 nil 的切片对应的 slice 是否为零值。
```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	var sl []string
	fmt.Printf("sl value is %+v\n", *(*slice)(unsafe.Pointer(&sl)))
}
// 运行输出：
// sl value is {array:<nil> len:0 cap:0}
// 是零值。
```

至此解释了开篇出乎意料的比较结果背后的原因：空切片为 nil 因为其值为零值，类型为 []string 的空切片传给空接口后，因为空接口的值并不是零值，所以接口变量不是 nil。

有两个办法：
1. 既然值为 nil 的具型变量赋值给空接口会出现如此莫名其妙的情况，我们不要这么做，再赋值前先做判空处理，不为 nil 才赋给空接口；
2. 使用reflect.ValueOf().IsNil()来判断。不推荐这种做法，因为当空接口对应的具型是值类型，会 panic。

```go
func IsNil(i interface{}) {
	if i != nil {
		if reflect.ValueOf(i).IsNil() {
			fmt.Println("i is nil")
			return
		}
		fmt.Println("i isn't nil")
	}
	fmt.Println("i is nil")
}

func main() {
	var sl []string
	IsNil(sl)	// i is nil
	IsNil(nil)  // i is nil
}
```

Go 中变量是否为 nil 要看变量的值是否是零值。

Ref:

https://juejin.cn/post/7100078516334493733
https://www.cnblogs.com/wpgraceii/p/10528183.html