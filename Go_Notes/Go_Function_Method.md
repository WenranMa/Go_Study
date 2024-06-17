# Function, Method
## 函数F
### 函数声明(Declaration)
- parameter: 函数参数，局部变量，named result也是局部变量。
- argument: parameter values, supplied by the caller.
- 相同类型的参数可以只写一次类型。`func add(i, j int) int`.

Arguments are passed by value, so the function receives a copy of each argument; modifications to the copy do not affect the caller. However, if the argument contains some kind of reference, like a __pointer, slice, map, function, or channel__, then the caller may be affected.

### 递归(Recursion)
固定大小函数栈：Fixed-size function call stack; sizes from 64KB to 2MB are typical. Fixed-size stacks impose a limit on the depth of recursion.

Go implementations use __variable-size stacks__ that start small and grow as needed up to a limit on the order of a gigabyte. This lets us use recursion safely and without worrying about overflow.

### 多值返回
一个函数可以返回多个值。例如一个是期望得到的返回值，另一个是函数出错时的错误信息。一个函数内部可以将另一个有多返回值的函数作为返回值。

如果一个函数将所有的返回值都显示的变量名，那么该函数的return语句可以省略操作数。这称之 为bare return。
```go
func CountWordsAndImages(url string) (words, images int, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        err = fmt.Errorf("parsing HTML: %s", err)
        return
    }
    words, images = countWordsAndImages(doc)
    return
}
```
前两处return等价于 `return 0,0,err` (Go会将返回值words和images在函数体的开始处，根据它们的类型，将其初始化为0)，最后一处return等价于 `return words, image, nil`。

### Error
对于将运行失败看作是预期结果的函数，会返回一个额外的返回值，来传递错误信息。如果导致失败的原因只有一个，额外的返回值可以是一个布尔值，通常被命名为ok：
`value, ok := cache.Lookup(key)` 。

导致失败的原因不止一种，尤其是对I/O操作，用户需要了解更多的错误信息。所以额外的返回值不再是简单的布尔类型，而是error类型。内置的error是接口类型，可能是nil或者non-nil，nil意味着函数运行成功，non-nil表示失败。

`fmt.Errorf`函数使用fmt.Sprintf格式化错误信息并返回。
```go
func Errorf(format string, a ...interface{}) error {
    return errors.New(Sprintf(format, a...))
}
```

处理错误的策略：1.传播错误。2.重试并限制重试的时间间隔或重试的次数。3.输出错误信息并结束程序`os.Exit(1)`或者`log.Fatalf`，这种策略只应在main中执行。4.只输出错误，不中断程序。

io包保证任何由文件结束引起的读取失败都返回同一个错误 `io.EOF`: `var EOF = errors.New("EOF")`。

调用者只需通过简单的比较，就可以检测出这个错误。下面的例子展示了如何从标准输入中读取字符，以及判断文件结束。
```go
in := bufio.NewReader(os.Stdin)
for {
    r, _, err := in.ReadRune()
    if err == io.EOF {
        break // finished reading
    }
    if err != nil {
        return fmt.Errorf("read failed:%v", err)
    }
}
```

### 函数值
函数被看作第一类值(first­ class values):函数像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。对函数值(function value)的调用类似函数调用。
```go
func square(n int) int {
    return n * n
}
func main() {
    f := square
    fmt.Println(f(3)) // "9"

    var v func(int) int //函数类型
    if v != nil {
        v(3)
    }
}
```
函数类型的零值是nil。调用值为nil的函数值会引起panic错误。但是函数值之间是不可比较的，也不能用函数值作为map的key。


下面这段代码输出什么以及原因？
```go
func hello() []string {
    return nil
}

func main() {
    h := hello
    if h == nil {
        fmt.Println("nil")
    } else {
        fmt.Println("not nil")
    }
}
// not nil
// 这道题里面，是将 `hello()` 赋值给变量h，而不是函数的返回值，所以输出 `not nil` 
```


### 匿名函数
Named functions can be declared only at the package level, but we can use a function literal to denote a function value within any expression. A function literal is written like a function declaration, but without a name follow ing the `func` keyword. It is an expression, and its value is called an anonymous function.
```go
func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}
func main() {
    f := squares()
    fmt.Println(f()) // "1"
    fmt.Println(f()) // "4"
    fmt.Println(f()) // "9"
    fmt.Println(f()) // "16"

    v := squares()
    fmt.Println(v()) // 1
    fmt.Println(v()) // 4
    fmt.Println(v()) // 9
    fmt.Println(v()) // 16
}
```
在squares中定义的匿名内部函数可以访问和更新squares中的局部变量，这意味着匿名函数和squares中，存在变量引用。

__The squares example demonstrates that function values are not just code but can have state. 不只是代码，还有状态__ The anonymous inner function can access and update the local variables of the enclosing function squares. These hidden variable references are why we classify functions as __reference types__ and why function values are not comparable. Function values like these are implemented using a technique called closures(__闭包__), and Go programmers often use this term for function values.

Here we see an example where the lifetime of a variable is not determined by its scope: the variable x exists after squares has returned within main, even though x is hidden inside f.

在Go语言中，函数是一等（first-class）类型。这意味着，我们可以把函数作为值来传递和使用。函数代表着这样一个过程：它接受若干输入（参数），并经过一些步骤（语句）的执行之后再返回输出（结果）。特别的是，Go语言中的函数可以返回多个结果。

__函数类型__ 的字面量由关键字func、由圆括号包裹参数声明列表、空格以及可以由圆括号包裹的结果声明列表组成。其中，参数声明列表中的单个参数声明之间是由英文逗号分隔的。每个参数声明由参数名称、空格和参数类型组成。参数声明列表中的参数名称是可以被统一省略的。结果声明列表的编写方式与此相同。结果声明列表中的结果名称也是可以被统一省略的。并且，在只有一个无名称的结果声明时还可以省略括号。示例如下：

`func(input1 string ,input2 string) string`

这一类型字面量表示了一个接受两个字符串类型的参数且会返回一个字符串类型的结果的函数。如果我们在它的左边加入type关键字和一个标识符作为名称的话，那就变成了一个函数类型声明，就像这样：

`type MyFunc func(input1 string ,input2 string) string`

__函数值（或简称函数）__ 的写法与此不完全相同。编写函数的时候需要先写关键字func和函数名称，后跟参数声明列表和结果声明列表，最后是由花括号包裹的语句列表。例如：
```go
func myFunc(part1 string, part2 string) (result string) {
    result = part1 + part2
    return
}
```
我们在这里用到了一个小技巧：如果结果声明是带名称的，那么它就相当于一个已被声明但未被显式赋值的变量。我们可以为它赋值且在return语句中省略掉需要返回的结果值。至于什么是return语句，我就不用多说了吧。显然，该函数还有一种更常规的写法：
```go
func myFunc(part1 string, part2 string) string {
    return part1 + part2
}
```
注意，函数myFunc是函数类型MyFunc的一个实现。实际上，只要一个函数的参数声明列表和结果声明列表中的数据类型的顺序和名称与某一个函数类型完全一致，前者就是后者的一个实现。请大家回顾上面的示例并深刻理解这句话。
  
我们可以声明一个函数类型的变量，如：

`var splice func(string, string) string`
等价于 `var splice MyFunc` 然后把函数myFunc赋给它：

`splice = myFunc` 如此一来，我们就可以在这个变量之上实施调用动作了：`splice("1", "2")`

实际上，这是一个调用表达式。它由代表函数的标识符（这里是splice）以及代表调用动作的、由圆括号包裹的参数值列表组成。
  
如果你觉得上面对splice变量声明和赋值有些啰嗦，那么可以这样来简化它：
```go
var splice = func(part1 string, part2 string) string {
    return part1 + part2
}
```  
    
在这个示例中，我们直接使用了一个匿名函数来初始化splice变量。顾名思义，匿名函数就是不带名称的函数值。匿名函数直接由函数类型字面量和由花括号包裹的语句列表组成。注意，这里的函数类型字面量中的参数名称是不能被忽略的。
  
其实，我们还可以进一步简化——索性省去splice变量。既然我们可以在代表函数的变量上实施调用表达式，那么在匿名函数上肯定也是可行的。因为它们的本质是相同的。后者的示例如下：
```go
var result = func(part1 string, part2 string) string {
    return part1 + part2
}("1", "2")
```
可以看到，在这个匿名函数之后的即是代表调用动作的参数值列表。注意，这里的result变量的类型不是函数类型，而与后面的匿名函数的结果类型是相同的。
  
最后，函数类型的零值是nil。这意味着，一个未被显式赋值的、函数类型的变量的值必为nil。


### 可变参数
variadic function: that can be calle d with varying numbers of arguments. 典型的例子就是fmt.Printf和其变体。Printf首先接收一个的必备参数，之后接收任意个数的后续参数。在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上省略符号`...`，这表示该函数会接收任意数量的该类型参数。
```go
func sum(vals ...int) int {
    total := 0
    for _, val := range vals {
        total += val
    }
    return total
}
```


对 add() 函数调用正确的是（）
```go
func add(args ...int) int {

    sum := 0
    for _, arg := range args {
        sum += arg
    }
    return sum
}
// - A. add(1, 2)
// - B. add(1, 3, 7)
// - C. add([]int{1, 2})
// - D. add([]int{1, 3, 7}…)

// 答：ABD
```


下面这段代码有什么缺陷？
```go
func funcMui(x, y int) (sum int, error) {
    return x + y, nil
}
// `syntax error: mixed named and unnamed parameters`

// 在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。如果有多个返回值必须加上括号();如果只有一个返回值且命名也需要加上括号()。这里的第一个返回值有命名sum，第二个没有命名，编译错误。
```

关于函数声明，下面语法正确的是？ 答：A B D
```go
// A. func f(a, b int) (value int, err error)
// B. func f(a int, b int) (value int, err error)
// C. func f(a, b int) (value int, error)
// D. func f(a int, b int) (int, int, error)
```

下面的代码能否正确输出？
```go
func main() {
    var fn1 = func() {}
    var fn2 = func() {}

    if fn1 != fn2 {
        println("fn1 not equal fn2")
    }
}
// 编译错误
// invalid operation: fn1 != fn2 (func can only be compared to nil)
// 函数只能与 nil 比较。
```




## 方法
### 方法声明
在函数声明时，在其名字之前放上一个变量，即是一个方法。这个附加的参数会将该函数附加到这种类型上，相当于为这种类型定义了一个独占的方法。这个变量叫做方法的接收器(receiver)。

对于一个给定的类型，其内部的方法都必须有唯一的方法名，但是不同的类型却可以有同样的方法名。
```go
package main

import (
    "fmt"
    "math"
)

type Point struct {
    X, Y float64
}
//a method of the Point type, p is receiver.
func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point
// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
    sum := 0.0
    for i := range path {
        if i > 0 {
            sum += path[i-1].Distance(path[i])
        }
    }
    return sum
}

func main() {
    perim := Path{
        {1, 1},
        {5, 1},
        {5, 4},
        {1, 1},
    }
    fmt.Println(perim.Distance()) //12
}
```
`p.Distance` is called a selector, because it __selects__ the appropriate Distance method for the receiver p of type Point. Selectors are also used to select fields of struct types, as in p.X.

### 指针对象方法
当调用一个函数时，会对其每一个参数值进行拷贝：
```go
type Point struct {
    X, Y int
}
func (p Point) set(x int) Point {
    p.X = x
    return p
}
func main() {
    p := Point{3, 3}
    fmt.Println(p)         // {3, 3}
    fmt.Println(p.set(12)) // {12, 3}
    fmt.Println(p)         // {3, 3}
}
```
如果一个函数需要更新一个变量，或参数太大，需要用到指针：
```go
type Point struct {
    X, Y int
}
func (p *Point) set(x int) *Point {
    p.X = x  // 编译器在这里也会给我们隐式地插入*，与(*p).X = x 等价，这种简写只适用于变量。
    return p
}
func main() {
    p := &Point{3, 3}
    fmt.Println(*p)           // {3, 3}
    fmt.Println(*(p.set(12))) // {12, 3}
    fmt.Println(*p)           // {12, 3}
}
```
就像一些函数允许nil指针作为参数一样，方法理论上也可以用nil指针作为其接收器，尤其当nil对于对象来说是合法的零值时，比如map或者slice。

### 嵌入结构体扩展类型
```go
package main

import (
    "fmt"
    "image/color"
    "math"
)

type Point struct {
    X, Y float64
}

type ColoredPoint struct {
    Point
    Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
    var cp ColoredPoint
    cp.X = 1
    fmt.Println(cp.Point.X) // "1"
    cp.Point.Y = 2
    fmt.Println(cp.Y) // "2"  内嵌可以使我们在定义ColoredPoint时得到一种句法上的简写形式，并使其包含Point类型所具有的一切字段。

    red := color.RGBA{255, 0, 0, 255}
    blue := color.RGBA{0, 0, 255, 255}
    var p = ColoredPoint{Point{1, 1}, red}
    var q = ColoredPoint{Point{5, 4}, blue}
    fmt.Println(p.Distance(q.Point))  //方法也类似，可以省略p.Point.Distance()..
}
```

### 方法值 方法表达式
```go
p := Point{1, 2} 
q := Point{4, 6}
distanceFromP := p.Distance  //method value
fmt.Println(distanceFromP(q)) 
distance := Point.Distance  // method expression
fmt.Println(distance(p, q))
```

我们在前面多次提到过指针及指针类型。例如，*Person是Person的指针类型。又例如，表达式&p的求值结果是p的指针。方法的接收者类型的不同会给方法的功能带来什么影响？该方法所属的类型又会因此发生哪些潜移默化的改变？现在，我们就来解答第一个问题。

指针操作涉及到两个操作符——&和*。这两个操作符均有多个用途。但是当它们作为地址操作符出现时，前者的作用是取址，而后者的作用是取值。更通俗地讲，当地址操作符&被应用到一个值上时会取出指向该值的指针值，而当地址操作符*被应用到一个指针值上时会取出该指针指向的那个值。它们可以被视为相反的操作。
  
除此之外，当*出现在一个类型之前（如*Person和*[3]string）时就不能被看做是操作符了，而应该被视为一个符号。如此组合而成的标识符所表达的含义是作为第二部分的那个类型的指针类型。我们也可以把其中的第二部分所代表的类型称为基底类型。例如，*[3]string是数组类型[3]string的指针类型，而[3]string是*[3]string的基底类型。
  
好了，我们现在回过头去再看结构体类型Person。它及其两个方法的完整声明如下：
```go
type Person struct {
    Name    string
    Gender  string
    Age     uint8
    Address string
}

func (person *Person) Grow() {
    person.Age++
}

func (person *Person) Move(newAddress string) string {
    old := person.Address
    person.Address = newAddress
    return old
}
```
注意，Person的两个方法Grow和Move的接收者类型都是*Person，而不是Person。只要一个方法的接收者类型是其所属类型的指针类型而不是该类型本身，那么我就可以称该方法为一个指针方法。上面的Grow方法和Move方法都是Person类型的指针方法。
  
相对的，如果一个方法的接收者类型就是其所属的类型本身，那么我们就可以把它叫做值方法。我们只要微调一下Grow方法的接收者类型就可以把它从指针方法变为值方法：
```go
func (person Person) Grow() {
    person.Age++
}
```
那指针方法和值方法到底有什么区别呢？我们在保留上述修改的前提下编写如下代码：
```go
p := Person{"Robert", "Male", 33, "Beijing"}
p.Grow()
fmt.Printf("%v\n", p)  
// 33 
```
这段代码被执行后，标准输出会打印出什么内容呢？直觉上，34会被打印出来，但是被打印出来的却是33。
  
解答这个问题需要引出一条定论：__方法的接收者标识符所代表的是该方法当前所属的那个值的一个副本，而不是该值本身。__ 例如，在上述代码中，Person类型的Grow方法的接收者标识符person代表的是p的值的一个拷贝，而不是p的值。我们在调用Grow方法的时候，Go语言会将p的值复制一份并将其作为此次调用的当前值。正因为如此，Grow方法中的person.Age++语句的执行会使这个副本的Age字段的值变为34，而p的Age字段的值却依然是33。这就是问题所在。
  
只要我们把Grow变回指针方法就可以解决这个问题。原因是，这时的person代表的是p的值的指针的副本。指针的副本仍会指向p的值。另外，之所以选择表达式person.Age成立，是因为如果Go语言发现person是指针并且指向的那个值有Age字段，那么就会把该表达式视为(*person).Age。其实，这时的person.Age正是(*person).Age的速记法。


## 练习
下面这段代码输出什么？为什么？
```go
func (i int) PrintInt ()  {
    fmt.Println(i)
}

func main() {
    var i int = 1
    i.PrintInt()
}
// `cannot define new methods on non-local type int`
// 基于类型创建的方法必须定义在同一个包内，上面的代码基于 int 类型创建了 PrintInt() 方法，由于 int 类型和方法 PrintInt() 定义在不同的包内，所以编译出错。

// 解决的办法可以定义一种新的类型：
type Myint int

func (i Myint) PrintInt ()  {
    fmt.Println(i)
}

func main() {
    var i Myint = 1
    i.PrintInt()
}
```

下面这段代码输出什么？为什么？
```go
type People interface {
    Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
    if think == "speak" {
        talk = "speak"
    } else {
        talk = "hi"
    }
    return
}

func main() {
    var peo People = Student{}
    think := "speak"
    fmt.Println(peo.Speak(think))
}
// compilation error
// 编译错误 `Student does not implement People (Speak method has pointer receiver)`，值类型 `Student` 没有实现接口的 `Speak()` 方法，而是指针类型 `*Student` 实现该方法。要通过编译应该写成 ` var peo People = &Student{}`。
```

下面的代码有什么问题？---- TBD
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
// X{} 是不可寻址的，不能直接调用方法**
// `cannot call pointer method test on X`
// 在方法中，指针类型的接收者必须是合法指针（包括 nil）,或能获取实例地址。

//修复代码：
func main() {

    var a *X
    a.test()    // 相当于 test(nil)

    var x = X{}
    x.test()
}
```

下面的代码有几处问题？请详细说明。
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
// 有两处问题
// 1.直接返回的 T{} 不可寻址；
// 2.不可寻址的结构体不能调用带结构体指针接收者的方法；

// 修复代码：
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

下面的代码有什么问题？
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
// cannot assign to getT().n
// 直接返回的 T{} 无法寻址，不可直接赋值。

// 修复代码：
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

下面的代码有什么问题？
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
// calling method value with receiver p1 (type **N) requires explicit dereference
// calling method pointer with receiver p1 (type **N) requires explicit dereference
// 不能使用多级指针调用方法。

// 正确做法：
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

下面的代码输出什么？
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
// 10 11 12
// 方法表达式。
// 通过类型引用的方法表达式会被还原成普通函数样式，接收者是第一个参数，调用时显示传参。类型可以是 T 或 *T，只要目标方法存在于该类型的方法集中就可以。

// 还可以直接使用方法表达式调用：
func main()  {
    var n N = 10
    fmt.Println(n)
    n++
    N.test(n)
    n++
    (*N).test(&n)
}
```

下面代码输出什么？
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
// 13 11 12**
// 知识点：方法值。
// 当指针值赋值给变量或者作为函数参数传递时，会立即计算并复制该方法执行所需的接收者对象，与其绑定，以便在稍后执行时，能隐式第传入接收者参数。
```

下面的代码输出什么？
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
// 13 13 13
// 知识点：方法值
// 当目标方法的接收者是指针类型时，那么被复制的就是指针。
```

下面哪一行代码会 panic，请说明原因？
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
// 因为 s.bar 将被展开为 (*s.T).bar，而 s.T 是个空指针，解引用会 panic。

// 可以使用下面代码输出 s：
func main() {
    s := S{}
    fmt.Printf("%#v",s)   // 输出：main.S{T:(*main.T)(nil)}
}
```


下面代码输出什么？
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
// 100 110
// 闭包引用相同变量。
```

已知 Add() 函数的调用代码，则Add函数定义正确的是()  A,C
```go
func main() {
    var a Integer = 1
    var b Integer = 2
    var i interface{} = &a
    sum := i.(*Integer).Add(b)
    fmt.Println(sum)
}

// A
type Integer int
func (a Integer) Add(b Integer) Integer {
        return a + b
}

// B
type Integer int
func (a Integer) Add(b *Integer) Integer {
        return a + *b
}

// C
type Integer int
func (a *Integer) Add(b Integer) Integer {
        return *a + b
}

// D
type Integer int
func (a *Integer) Add(b *Integer) Integer {
        return *a + *b
}
```