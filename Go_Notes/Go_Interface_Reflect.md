# 接口

接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。

在Go语言中接口（interface）是一种类型，一种抽象的类型。相较于之前章节中讲到的那些具体类型（字符串、切片、结构体等）更注重“我是谁”，接口类型更注重“我能做什么”的问题。接口类型就像是一种约定——概括了一种类型应该具备哪些方法，在Go语言中提倡使用面向接口的编程方式实现解耦。

接口类型是一种抽象的类型。它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合。它们只会展示出它们自己的方法。也就是说当你有看到一个接口类型的值时，不知道它是什么，但知道可以通过它的方法来做什么。

## 接口类型

接口是一种由程序员来定义的类型，一个接口类型就是一组方法的集合，它规定了需要实现的所有方法。相较于使用结构体类型，当我们使用接口类型说明相比于它是什么更关心它能做什么。

### 接口的定义

每个接口类型由任意个方法签名组成，接口的定义格式如下：

```go
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
```

其中：

- 接口类型名：Go语言的接口在命名时，一般会在单词后面添加`er`，如有写操作的接口叫`Writer`，有关闭操作的接口叫`closer`等。接口名最好要能突出该接口的类型含义。
- 方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
- 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

举个例子，定义一个包含`Write`方法的`Writer`接口。

```go
type Writer interface{
    Write([]byte) error
}
```

当你看到一个`Writer`接口类型的值时，你不知道它是什么，唯一知道的就是可以通过调用它的`Write`方法来做一些事情。

### 实现接口的条件

接口就是规定了一个**需要实现的方法列表**，在 Go 语言中一个类型只要实现了接口中规定的所有方法，那么我们就称它实现了这个接口。

我们定义的`Singer`接口类型，它包含一个`Sing`方法。

```go
// Singer 接口
type Singer interface {
	Sing()
}
```

我们有一个`Bird`结构体类型如下。

```go
type Bird struct {}
```

因为`Singer`接口只包含一个`Sing`方法，所以只需要给`Bird`结构体添加一个`Sing`方法就可以满足`Singer`接口的要求。

```go
// Sing Bird类型的Sing方法
func (b Bird) Sing() {
	fmt.Println("汪汪汪")
}
```

这样就称为`Bird`实现了`Singer`接口。

### 为什么要使用接口？

现在假设我们的代码世界里有很多小动物，下面的代码片段定义了猫和狗，它们饿了都会叫。

```go
package main

import "fmt"

type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}

type Dog struct{}

func (d Dog) Say() {
	fmt.Println("汪汪汪~")
}

func main() {
	c := Cat{}
	c.Say()
	d := Dog{}
	d.Say()
}
```

这个时候又跑来了一只羊，羊饿了也会发出叫声。

```go
type Sheep struct{}

func (s Sheep) Say() {
	fmt.Println("咩咩咩~")
}
```

我们接下来定义一个饿肚子的场景。

```go
// MakeCatHungry 猫饿了会喵喵喵~
func MakeCatHungry(c Cat) {
	c.Say()
}

// MakeSheepHungry 羊饿了会咩咩咩~
func MakeSheepHungry(s Sheep) {
	s.Say()
}
```

接下来会有越来越多的小动物跑过来，我们的代码世界该怎么拓展呢？

在饿肚子这个场景下，我们可不可以把所有动物都当成一个“会叫的类型”来处理呢？当然可以！使用接口类型就可以实现这个目标。 我们的代码其实并不关心究竟是什么动物在叫，我们只是在代码中调用它的`Say()`方法，这就足够了。

我们可以约定一个`Sayer`类型，它必须实现一个`Say()`方法，只要饿肚子了，我们就调用`Say()`方法。

```go
type Sayer interface {
    Say()
}
```

然后我们定义一个通用的`MakeHungry`函数，接收`Sayer`类型的参数。

```go
// MakeHungry 饿肚子了...
func MakeHungry(s Sayer) {
	s.Say()
}
```

我们通过使用接口类型，把所有会叫的动物当成`Sayer`类型来处理，只要实现了`Say()`方法都能当成`Sayer`类型的变量来处理。

```go
var c cat
MakeHungry(c)
var d dog
MakeHungry(d)
```

在电商系统中我们允许用户使用多种支付方式（支付宝支付、微信支付、银联支付等），我们的交易流程中可能不太在乎用户究竟使用什么支付方式，只要它能提供一个实现支付功能的`Pay`方法让调用方调用就可以了。

再比如我们需要在某个程序中添加一个将某些指标数据向外输出的功能，根据不同的需求可能要将数据输出到终端、写入到文件或者通过网络连接发送出去。在这个场景下我们可以不关注最终输出的目的地是什么，只需要它能提供一个`Write`方法让我们把内容写入就可以了。

Go语言中为了解决类似上面的问题引入了接口的概念，接口类型区别于我们之前章节中介绍的那些具体类型，让我们专注于该类型提供的方法，而不是类型本身。使用接口类型通常能够让我们写出更加通用和灵活的代码。

### 面向接口编程

PHP、Java等语言中也有接口的概念，不过在PHP和Java语言中需要显式声明一个类实现了哪些接口，在Go语言中使用隐式声明的方式实现接口。只要一个类型实现了接口中规定的所有方法，那么它就实现了这个接口。

Go语言中的这种设计符合程序开发中抽象的一般规律，例如在下面的代码示例中，我们的电商系统最开始只设计了支付宝一种支付方式：

```go
type ZhiFuBao struct {
	// 支付宝
}

// Pay 支付宝的支付方法
func (z *ZhiFuBao) Pay(amount int64) {
  fmt.Printf("使用支付宝付款：%.2f元。\n", float64(amount/100))
}

// Checkout 结账
func Checkout(obj *ZhiFuBao) {
	// 支付100元
	obj.Pay(100)
}

func main() {
	Checkout(&ZhiFuBao{})
}
```

随着业务的发展，根据用户需求添加支持微信支付。

```go
type WeChat struct {
	// 微信
}

// Pay 微信的支付方法
func (w *WeChat) Pay(amount int64) {
	fmt.Printf("使用微信付款：%.2f元。\n", float64(amount/100))
}
```

在实际的交易流程中，我们可以根据用户选择的支付方式来决定最终调用支付宝的Pay方法还是微信支付的Pay方法。

```go
// Checkout 支付宝结账
func CheckoutWithZFB(obj *ZhiFuBao) {
	// 支付100元
	obj.Pay(100)
}

// Checkout 微信支付结账
func CheckoutWithWX(obj *WeChat) {
	// 支付100元
	obj.Pay(100)
}
```

实际上，从上面的代码示例中我们可以看出，我们其实并不怎么关心用户选择的是什么支付方式，我们只关心调用Pay方法时能否正常运行。这就是典型的“不关心它是什么，只关心它能做什么”的场景。

在这种场景下我们可以将具体的支付方式抽象为一个名为`Payer`的接口类型，即任何实现了`Pay`方法的都可以称为`Payer`类型。

```go
// Payer 包含支付方法的接口类型
type Payer interface {
	Pay(int64)
}
```

此时只需要修改下原始的`Checkout`函数，它接收一个`Payer`类型的参数。这样就能够在不修改既有函数调用的基础上，支持新的支付方式。

```go
// Checkout 结账
func Checkout(obj Payer) {
	// 支付100元
	obj.Pay(100)
}

func main() {
	Checkout(&ZhiFuBao{}) // 之前调用支付宝支付

	Checkout(&WeChat{}) // 现在支持使用微信支付
}
```

像类似的例子在我们编程过程中会经常遇到：

- 比如一个网上商城可能使用支付宝、微信、银联等方式去在线支付，我们能不能把它们当成“支付方式”来处理呢？
- 比如三角形，四边形，圆形都能计算周长和面积，我们能不能把它们当成“图形”来处理呢？
- 比如满减券、立减券、打折券都属于电商场景下常见的优惠方式，我们能不能把它们当成“优惠券”来处理呢？

接口类型是Go语言提供的一种工具，在实际的编码过程中是否使用它由你自己决定，但是通常使用接口类型可以使代码更清晰易读。

### 接口类型变量

那实现了接口又有什么用呢？一个接口类型的变量能够存储所有实现了该接口的类型变量。

例如在上面的示例中，`Dog`和`Cat`类型均实现了`Sayer`接口，此时一个`Sayer`类型的变量就能够接收`Cat`和`Dog`类型的变量。

```go
var x Sayer // 声明一个Sayer类型的变量x
a := Cat{}  // 声明一个Cat类型变量a
b := Dog{}  // 声明一个Dog类型变量b
x = a       // 可以把Cat类型变量直接赋值给x
x.Say()     // 喵喵喵
x = b       // 可以把Dog类型变量直接赋值给x
x.Say()     // 汪汪汪
```

## 值接收者和指针接收者

在结构体那一章节中，我们介绍了在定义结构体方法时既可以使用值接收者也可以使用指针接收者。那么对于实现接口来说使用值接收者和使用指针接收者有什么区别呢？接下来我们通过一个例子看一下其中的区别。

我们定义一个`Mover`接口，它包含一个`Move`方法。

```go
// Mover 定义一个接口类型
type Mover interface {
	Move()
}
```

### 值接收者实现接口

我们定义一个`Dog`结构体类型，并使用值接收者为其定义一个`Move`方法。

```go
// Dog 狗结构体类型
type Dog struct{}

// Move 使用值接收者定义Move方法实现Mover接口
func (d Dog) Move() {
	fmt.Println("狗会动")
}
// 此时实现`Mover`接口的是`Dog`类型。

var x Mover    // 声明一个Mover类型的变量x

var d1 = Dog{} // d1是Dog类型
x = d1         // 可以将d1赋值给变量x
x.Move()

var d2 = &Dog{} // d2是Dog指针类型
x = d2          // 也可以将d2赋值给变量x
x.Move()
```
从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是结构体类型还是对应的结构体指针类型的变量都可以赋值给该接口变量。

### 指针接收者实现接口

我们再来测试一下使用指针接收者实现接口有什么区别。

```go
// Cat 猫结构体类型
type Cat struct{}

// Move 使用指针接收者定义Move方法实现Mover接口
func (c *Cat) Move() {
	fmt.Println("猫会动")
}
// 此时实现`Mover`接口的是`*Cat`类型，我们可以将`*Cat`类型的变量直接赋值给`Mover`接口类型的变量`x`。

var c1 = &Cat{} // c1是*Cat类型
x = c1          // 可以将c1当成Mover类型
x.Move()

// 但是不能给将`Cat`类型的变量赋值给`Mover`接口类型的变量`x`。

// 下面的代码无法通过编译
var c2 = Cat{} // c2是Cat类型
x = c2         // 不能将c2当成Mover类型
```

**注意** 对于每一个命名过的具体类型T;它一些方法的接收者是类型T本身而另一些则是一个T指针。`T类型变量上调用\*T的方法是合法的，但必须是变量。编译器隐式的获取了它的地址。但这仅仅是一个语法糖:T类型的值不拥有*T指针的方法。` 由于Go语言中有对指针求值的语法糖，对于值接收者实现的接口，无论使用值类型还是指针类型都没有问题。但是我们并不总是能对一个值求址，所以对于指针接收者实现的接口要额外注意。

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
但由于只有\*IntSet类型有String方法，也只有\*IntSet类型实现了fmt.Stringer接口
```go
var _ fmt.Stringer = &s // OK
var _ fmt.Stringer = s // compile error: IntSet lacks String method
```

## 类型与接口的关系

### 一个类型实现多个接口

一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。例如狗不仅可以叫，还可以动。我们完全可以分别定义`Sayer`接口和`Mover`接口，具体代码示例如下。

```go
// Sayer 接口
type Sayer interface {
	Say()
}

// Mover 接口
type Mover interface {
	Move()
}
// Dog既可以实现`Sayer`接口，也可以实现`Mover`接口。

type Dog struct {
	Name string
}

// 实现Sayer接口
func (d Dog) Say() {
	fmt.Printf("%s会叫汪汪汪\n", d.Name)
}

// 实现Mover接口
func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.Name)
}
// 同一个类型实现不同的接口互相不影响使用。

var d = Dog{Name: "旺财"}

var s Sayer = d
var m Mover = d

s.Say()  // 对Sayer类型调用Say方法
m.Move() // 对Mover类型调用Move方法
```

### 多种类型实现同一接口

Go语言中不同的类型还可以实现同一接口。例如在我们的代码世界中不仅狗可以动，汽车也可以动。我们可以使用如下代码体现这个关系。

```go
// 实现Mover接口
func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.Name)
}

// Car 汽车结构体类型
type Car struct {
	Brand string
}

// Move Car类型实现Mover接口
func (c Car) Move() {
	fmt.Printf("%s速度70迈\n", c.Brand)
}
```

这样我们在代码中就可以把狗和汽车当成一个会动的类型来处理，不必关注它们具体是什么，只需要调用它们的`Move`方法就可以了。

```go
var obj Mover

obj = Dog{Name: "旺财"}
obj.Move()

obj = Car{Brand: "宝马"}
obj.Move()

//旺财会动
//宝马速度70迈
```

一个接口的所有方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。

```go
// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}
```

## 接口组合

接口与接口之间可以通过互相嵌套形成新的接口类型，例如Go标准库`io`源码中就有很多接口之间互相组合的示例。

```go
// src/io/io.go

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// ReadWriter 是组合Reader接口和Writer接口形成的新接口类型
type ReadWriter interface {
	Reader
	Writer
}

// ReadCloser 是组合Reader接口和Closer接口形成的新接口类型
type ReadCloser interface {
	Reader
	Closer
}

// WriteCloser 是组合Writer接口和Closer接口形成的新接口类型
type WriteCloser interface {
	Writer
	Closer
}
```

对于这种由多个接口类型组合形成的新接口类型，同样只需要实现新接口类型中规定的所有方法就算实现了该接口类型。

接口也可以作为结构体的一个字段，我们来看一段Go标准库`sort`源码中的示例。

```go
// src/sort/sort.go

// Interface 定义通过索引对元素排序的接口类型
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}


// reverse 结构体中嵌入了Interface接口
type reverse struct {
    Interface
}
```

通过在结构体中嵌入一个接口类型，从而让该结构体类型实现了该接口类型，并且还可以改写该接口的方法。

```go
// Less 为reverse类型添加Less方法，重写原Interface接口类型的Less方法
func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}
```

`Interface`类型原本的`Less`方法签名为`Less(i, j int) bool`，此处重写为`r.Interface.Less(j, i)`，即通过将索引参数交换位置实现反转。

在这个示例中还有一个需要注意的地方是`reverse`结构体本身是不可导出的（结构体类型名称首字母小写），`sort.go`中通过定义一个可导出的`Reverse`函数来让使用者创建`reverse`结构体实例。

```go
func Reverse(data Interface) Interface {
	return &reverse{data}
}
```

这样做的目的是保证得到的`reverse`结构体中的`Interface`属性一定不为`nil`，否者`r.Interface.Less(j, i)`就会出现空指针panic。

此外在Go内置标准库`database/sql`中也有很多类似的结构体内嵌接口类型的使用示例，各位读者可自行查阅。

## 空接口

### 空接口的定义

空接口是指没有定义任何方法的接口类型。因此任何类型都可以视为实现了空接口。也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值。

```go
package main

import "fmt"

// 空接口

// Any 不包含任何方法的空接口类型
type Any interface{}

// Dog 狗结构体
type Dog struct{}

func main() {
	var x Any

	x = "你好" // 字符串型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = 100 // int型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = true // 布尔型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = Dog{} // 结构体类型
	fmt.Printf("type:%T value:%v\n", x, x)
}
```

通常我们在使用空接口类型时不必使用`type`关键字声明，可以像下面的代码一样直接使用`interface{}`。

```go
var x interface{}  // 声明一个空接口类型变量x
```

### 空接口的应用

#### 空接口作为函数的参数

使用空接口实现可以接收任意类型的函数参数。

```go
// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
```

#### 空接口作为map的值

使用空接口实现可以保存任意值的字典。

```go
// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
```

## 接口值

由于接口类型的值可以是任意一个实现了该接口的类型值，所以接口值除了需要记录具体**值**之外，还需要记录这个值属于的**类型**。也就是说一个`接口的值`由两个部分组成，一个具体的`类型`和那个`类型的值`。鉴于这两部分会根据存入值的不同而发生变化，它们被称为接口的`动态类型`和`动态值`。

我们接下来通过一个示例来加深对接口值的理解。

下面的示例代码中，定义了一个`Mover`接口类型和两个实现了该接口的`Dog`和`Car`结构体类型。

```go
type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

func (d *Dog) Move() {
	fmt.Println("狗在跑~")
}

type Car struct {
	Brand string
}

func (c *Car) Move() {
	fmt.Println("汽车在跑~")
}
```

首先，我们创建一个`Mover`接口类型的变量`m`。此时，接口变量`m`是接口类型的零值，也就是它的类型和值部分都是`nil`，我们可以使用`m == nil`来判断此时的接口值是否为空。**注意：**我们不能对一个空接口值调用任何方法，否则会产生panic。
```go
var m Mover
fmt.Println(m == nil)  // true
m.Move() // panic: runtime error: invalid memory address or nil pointer dereference
```

接下来，我们将一个`*Dog`结构体指针赋值给变量`m`。此时，接口值`m`的动态类型会被设置为`*Dog`，动态值为结构体变量的拷贝。然后，我们给接口变量`m`赋值为一个`*Car`类型的值。这一次，接口值`m`的动态类型为`*Car`，动态值为`nil`。**注意：**此时接口变量`m`与`nil`并不相等，因为它只是动态值的部分为`nil`，而动态类型部分保存着对应值的类型。一个不包含任何值的nil接口值和一个刚好包含nil指针的接口值是不同的。

```go
m = &Dog{Name: "旺财"}

var c *Car
m = c
fmt.Println(m == nil) // false
```

接口值是支持相互比较的，当且仅当接口值的动态类型和动态值都相等时才相等。
两个接口值相等仅当它们`都是nil值`或者`它们的动态类型相同并且动态值也根据这个动态类型的==操作相等`。

```go
var (
	x Mover = new(Dog)
	y Mover = new(Car)
)
fmt.Println(x == y) // false
```

但是有一种特殊情况需要特别注意，如果接口值保存的动态类型相同，但是这个动态类型不支持互相比较（比如切片），那么对它们相互比较时就会引发panic。

```go
var z interface{} = []int{1, 2, 3}
fmt.Printf("%T\n", z) // []int
fmt.Println(z == z) // panic: runtime error: comparing uncomparable type []int
```

对于一个`接口的零值就是它的类型和值的部分都是nil`。

## 类型断言

接口值可能赋值为任意类型的值，那我们如何从接口值获取其存储的具体数据呢？我们可以借助标准库`fmt`包的格式化打印获取到接口值的动态类型。
```go
var m Mover

m = &Dog{Name: "旺财"}
fmt.Printf("%T\n", m) // *main.Dog

m = new(Car)
fmt.Printf("%T\n", m) // *main.Car
```

而`fmt`包内部其实是使用反射的机制在程序运行时获取到动态类型的名称。

而想要从接口值中获取到对应的实际值需要使用类型断言，其语法格式如下。

```go
x.(T)
```

其中：

- x：表示接口类型的变量
- T：表示断言`x`可能是的类型。

该语法返回两个参数，第一个参数是`x`转化为`T`类型后的变量，第二个值是一个布尔值，若为`true`则表示断言成功，为`false`则表示断言失败。

举个例子：

```go
var n Mover = &Dog{Name: "旺财"}
v, ok := n.(*Dog)
if ok {
	fmt.Println("类型断言成功")
	v.Name = "富贵" // 变量v是*Dog类型
} else {
	fmt.Println("类型断言失败")
}
```

如果对一个接口值有多个实际类型需要判断，推荐使用`switch`语句来实现。

```go
// justifyType 对传入的空接口类型变量x进行类型断言
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
```

类型断言是一个`使用在接口值上的操作`。`x.(T)`，一个类型断言检查x对象的类型是否和断言的类型T匹配。`不能在一个非接口类型的值上应用类型断言来判定它是否属于某一个接口类型的。我们必须先把前者转换成空接口类型的值`。这又涉及到了Go语言的类型转换。

Go语言的类型转换规则定义了是否能够以及怎样可以把一个类型的值转换另一个类型的值。另一方面，所谓`空接口类型即是不包含任何方法声明的接口类型，用interface{}表示，常简称为空接口`。`interface{}`被称为`空接口类型`，因为空接口类型对实现它的类型没有要求，所以我们可以将任意一个值赋给空接口类型。`var any interface{}`。Go语言中的包含预定义的任何数据类型都可以被看做是空接口的实现。我们可以直接使用类型转换表达式把一个*Person类型转换成空接口类型的值，就像这样：
```go
p := Person{"Robert"}
v := interface{}(&p)

h, ok := v.(Animal)
```
请注意第二行。在类型字面量后跟由圆括号包裹的值（或能够代表它的变量、常量或表达式）就构成了一个类型转换表达式，意为将后者转换为前者类型的值。在这里，我们把表达式&p的求值结果转换成了一个空接口类型的值，并由变量v代表。注意，表达式&p（&是取址操作符）的求值结果是一个*Person类型的值，即p的指针。在这之后，我们就可以在v上应用类型断言。

`如果T是一个具体类型，类型断言检查x的动态类型是否和T相同`。如果这个检查成功，类型断言的结果是x的动态值。如果检查失败，接下来这个操作会抛出panic。`T也可以是一个接口类型，类型断言检查x的动态类型是否满足T（实现了T的方法）。`

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

我们在讲接口的时候说过，如果一个数据类型所拥有的方法集合中包含了某一个接口类型中的所有方法声明的实现，那么就可以说这个数据类型实现了那个接口类型。要获知一个数据类型都包含哪些方法并不难。但是要注意指针方法与值方法的区别。

拥有指针方法Move的指针类型*Person是接口类型Animal的实现类型，但是它的基底类型Person却不是。这样的表象隐藏着另一条规则：`一个指针类型拥有以它以及以它的基底类型为接收者类型的所有方法，而它的基底类型却只拥有以它本身为接收者类型的方法。`

以Dog类型为例。即使我们把Name, Age改为值方法，\*Dog类型也仍会是Pet接口的实现类型。另一方面，如果Name, Age是指针方法，Dog类型就不可能是Pet接口的实现类型，只有\*Dog类型才是。

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

`断言操作的对象是一个nil接口值，断言都会失败。`

只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。切记不要为了使用接口类型而增加不必要的抽象，导致不必要的运行时损耗。

在 Go 语言中接口是一个非常重要的概念和特性，使用接口类型能够实现代码的抽象和解耦，也可以隐藏某个功能的内部实现，但是缺点就是在查看源码的时候，不太方便查找到具体实现接口的类型。

相信很多读者在刚接触到接口类型时都会有很多疑惑，请牢记接口是一种类型，一种抽象的类型。区别于我们在之前章节提到的那些具体类型（整型、数组、结构体类型等），它是一个只要求实现特定方法的抽象类型。

**小技巧：** 下面的代码可以在程序编译阶段验证某一结构体是否满足特定的接口类型。
```go
// 摘自gin框架routergroup.go
type IRouter interface{ ... }

type RouterGroup struct { ... }

var _ IRouter = &RouterGroup{}  // 确保RouterGroup实现了接口IRouter
```
上面的代码中也可以使用`var _ IRouter = (*RouterGroup)(nil)`进行验证。

## 接口实例
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

下面代码输出？
```go
type Mover2 interface {
	move2()
}

type Pig struct {
	Name string
}

func (p Pig) move2() {
	//TODO implement me
	panic("implement me")
}

func main() {
	var n Mover2 = &Pig{Name: "pig"}  //var n Mover2 = Pig{Name: "pig"}  
	v, ok := n.(*Pig) // v, ok := n.(Pig) 改成这两行也work
	if ok {
		fmt.Println("类型断言成功")
		v.Name = "pig1" // 变量v是*Pig类型
		fmt.Println(v.Name)
	} else {
		fmt.Println("类型断言失败")
	}
}
// 类型断言成功
// pig1
// 指针类型可以获得值方法
```

下面代码输出？
```go
package main

import "fmt"

type WashingMachine interface {
	wash()
	dry()
}

type dryer struct{}

func (d dryer) dry() {
	fmt.Println("甩一甩")
}

type haier struct {
	dryer
}

func (h haier) wash() {
	fmt.Println("洗刷刷")
}

var h haier
var wa WashingMachine

func main() {
	wa = h
	wa.wash() // 洗刷刷
	wa.dry() // 甩一甩

	var haier = haier{}
	var s = WashingMachine(haier)
	s.wash() // 洗刷刷
	s.dry() // 甩一甩
	haier.dry() // 甩一甩
	haier.dryer.dry() // 甩一甩
	haier.wash() // 洗刷刷
}
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

关于GetPodAction定义，下面赋值正确的是。 答：A C D
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
// A. var fragment Fragment = new(GetPodAction)
// B. var fragment Fragment = GetPodAction
// C. var fragment Fragment = &GetPodAction{}
// D. var fragment Fragment = GetPodAction{}
```

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


## 其他
//请您简短介绍下反射具体可以实现哪些功能，如果您在工作中曾经使用过反射，请说说具体的应用场景？
//1. golang中反射最常见的使用场景是做对象的序列化，go语言标准库的encoding/json、xml、gob、binary等包就大量依赖于反射功能来实现
//2. 反射能够在运行时检查类型，它还允许在运行时检查，修改和创建变量，函数和结构体。


```go
// tag获取
package main

import (
	"reflect"
	"fmt"
)

type Server struct {
	ServerName string `key1:"value1" key11:"value11"`
	ServerIP   string `key2:"value2"`
}

func main() {
	s := Server{}
	st := reflect.TypeOf(s)

	field1 := st.Field(0)
	fmt.Printf("key1:%v\n", field1.Tag.Get("key1"))
	fmt.Printf("key11:%v\n", field1.Tag.Get("key11"))

	filed2 := st.Field(1)
	fmt.Printf("key2:%v\n", filed2.Tag.Get("key2"))

}
```