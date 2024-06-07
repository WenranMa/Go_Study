


# interface， reflect

## 18. 下面这段代码能否编译通过？如果可以，输出什么？

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
```

**答：编译失败**

**解析：**

`i (variable of type int) is not an interface`

只有接口类型才能使用类型选择
类型选择的语法形如：i.(type)，其中i是接口，type是固定关键字。



## 26. 下面这段代码输出什么？

```go
func main() {  
    var i interface{}
    if i == nil {
        fmt.Println("nil")
        return
    }
    fmt.Println("not nil")
}
```

- A. nil
- B. not nil
- C. compilation error  

**答：A**

**解析：**

当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil

## 39. 下面这段代码输出什么？

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
```

**答：13 23**

**解析：**

知识点：接口。

一种类型实现多个接口，结构体 Work 分别实现了接口 A、B，所以接口变量 a、b 调用各自的方法 ShowA() 和 ShowB()，输出 13、23。

## 45. 下面代码输出什么？

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
```

- A. 13 23
- B. compilation error

**答：A**

**解析：**

知识点：类型断言。

## 42. 下面代码输出什么？

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
```

- A. 23 13
- B. compilation error

**答：B**

**解析：**

知识点：接口的静态类型。

结合39题一起。a、b 具有相同的动态类型和动态值，分别是结构体 work 和 {3}；a 的静态类型是 A，b 的静态类型是 B，接口 A 不包括方法 ShowB()，接口 B 也不包括方法 ShowA()，编译报错。看下编译错误：

```shell
a.ShowB undefined (type A has no field or method ShowB)
b.ShowA undefined (type B has no field or method ShowA)
```

## 62. 下面这段代码输出什么？为什么？

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
```

**答：`s is nil` 和 `p is not nil`**

**解析：**

结合26题。

这道题会不会有点诧异，我们分配给变量 p 的值明明是 nil，然而 p 却不是 nil。记住一点，**当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil**。上面的代码，给变量 p 赋值之后，p 的动态值是 nil，但是动态类型却是 *Student，是一个 nil 指针，所以相等条件不成立。


## 106. 下面代码输出什么？--- TBD

```go
func main() {
    x := interface{}(nil)
    y := (*int)(nil)
    a := y == x
    b := y == nil
    _, c := x.(interface{})
    println(a, b, c)
}
```

- A. true true true
- B. false true true
- C. true true true
- D. false true false

**答：D**

**解析：**

知识点：类型断言。

类型断言语法：i.(Type)，其中 i 是接口，Type 是类型或接口。编译时会自动检测 i 的动态类型与 Type 是否一致。但是，如果动态类型不存在，则断言总是失败。


## 190. 下面的代码有什么问题，请说明。

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
```

**答：编译报错**

```shell
cannot use data literal (type data) as type printer in assignment:
data does not implement printer (print method has pointer receiver)
```

**解析：**

结构体类型 data 没有实现接口 printer。而是*data实现了接口。

## 181. interface{} 是可以指向任意对象的 Any 类型，是否正确？

- A. false
- B. true

**答：B**


## 94. 下面代码是否能编译通过？如果通过，输出什么？

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
```

**答：non-empty interface**

**解析：**

结合26，62题。

考点：interface 的内部结构，我们知道接口除了有静态类型，还有动态类型和动态值，当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。这里的 x 的动态类型是 `*int`，所以 x 不为 nil。

## 51. A、B、C、D 哪些选项有语法错误？

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
```

**答：BD**

**解析：**

```
cannot use s (variable of type S) as *interface{} value in argument to g: S does not implement *interface{} (type *interface{} is pointer to interface, not interface)
cannot use p (variable of type *S) as *interface{} value in argument to g: *S does not implement *interface{} (type *interface{} is pointer to interface, not interface)
```

函数参数为 interface{} 时可以接收任何类型的参数，包括用户自定义类型等，即使是接收指针类型也用 interface{}，而不是使用 *interface{}。


## 137. 下面哪一行代码会 panic，请说明原因？

```go
package main

func main() {
  var x interface{}
  var y interface{} = []int{3, 5}
  _ = x == x
  _ = x == y
  _ = y == y
}
```

**答：第 8 行 _ = y == y**

**解析：**

`comparing uncomparable type []int`

因为两个比较值的动态类型为同一个不可比较类型。



## interface --- TBD

interface 是 Go 里所提供的非常重要的特性。一个 interface 里可以定义一个或者多个函数，例如系统自带的io.ReadWriter的定义如下所示：
```go
type ReadWriter interface {
    Read(b []byte) (n int, err error)
    Write(b []byte) (n int, err error)
}
```
任何类型只要它提供了 Read 和 Write 的实现，那么这个类型便实现了这个 interface（duck-type）。

### 一个nil与interface{}比较的问题。

首先看两个例子：
```go
func IsNil(i interface{}) {
	if i == nil {
		fmt.Println("i is nil")
		return
	}
	fmt.Println("i isn't nil")
}

func main() {
	var sl []string
	if sl == nil {
		fmt.Println("sl is nil")
	}
	IsNil(sl)
}
// 输出：
// sl is nil
// i isn't nil


func main() {
	var x interface{} = nil
	var y *int = nil
	interfaceIsNil(x)
	interfaceIsNil(y)
}

func interfaceIsNil(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
// 输出：
// empty interface
// non-empty interface
```

为啥一个nil切片，nil int指针，经过空接口interface{}一中转，就变成了非 nil。

想要理解这个问题，首先需要理解 interface{} 变量的本质。

Go 语言中有两种略微不同的接口，**一种是带有一组方法的接口，另一种是不带任何方法的空接口 interface{}。**

Go 语言使用runtime.iface表示带方法的接口，使用runtime.eface表示不带任何方法的空接口interface{}。

一个 interface{} 类型的变量包含了 2 个指针，一个指针指向值的类型，另外一个指针指向实际的值。在 Go 源码中 runtime 包下，我们可以找到runtime.eface的定义。
```go
type eface struct { // 16 字节
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

再细究的话，你可能会问，既然 sl 是一个有类型有值的切片，为什么又是个 nil。

针对具体类型的变量，判断是否是 nil 要根据其值是否为零值。因为 sl 一个切片类型，而切片类型的定义在源码包src/runtime/slice.go我们可以找到。

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
```
不出所料，果然是零值。

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

切记，Go 中变量是否为 nil 要看变量的值是否是零值。

切记，不要将值为 nil 的变量赋给空接口。


https://juejin.cn/post/7100078516334493733






### interface怎么判断nil

https://www.cnblogs.com/wpgraceii/p/10528183.html