# defer,panic,recover
## defer
当defer语句被执行时，跟在defer后面的函数会被延迟执行。直到包含该defer语句的函数执行完毕时，defer后的函数才会被执行，不论包含defer语句的函数是通过return正常结束，还是由于panic导致的异常结束。可以在一个函数中执行多条defer语句，它们的执行顺序与声明顺序相反。

defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。通过defer机制，不论函数逻辑多复杂，都能保证在任何执行路径下，资源被释放。释放资源的defer应该直接跟在请求资源的语句后。

`defer语句中的函数会在return语句更新返回值变量后再执行，又因为在函数中定义的匿名函数可以访问该函数包括返回值变量在内的所有变量，所以，对匿名函数采用defer机制，可以使其观察函数的返回值。`
```go
func main() {
    _ = double(4) // "double(4) = 9"
}

func double(x int) (result int) {
    defer func() {
        result += 1
        fmt.Printf("double(%d) = %d\n", x, result) 
    }()
    return x + x
}
```

Defer栈，defer的特点就是LIFO，即后进先出，所以如果在同一个函数下多个defer的话，会逆序执行。
```go
func main() {
    f(3)
}
func f(x int) {
    fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
    defer fmt.Printf("defer %d\n", x)
    f(x - 1)
}
/* output:
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
panic: runtime error: integer divide by zero
*/
```

defer携带的表达式语句代表的是对某个函数或方法的调用。这个调用可能会有参数传入，比如：fmt.Print(i + 1)。如果传入参数的是一个表达式，那么在defer语句被执行的时候该表达式就会被求值了。注意，这与被携带的表达式语句的执行时机是不同的。请揣测下面这段代码的执行：
```go
func deferIt3() {
    f := func(i int) int {
        fmt.Printf("%d ",i)
        return i * 10
    }
    for i := 1; i < 5; i++ {
        defer fmt.Printf("%d ", f(i))
    }
}
// 它在被执行之后，标准输出上打印出1 2 3 4 40 30 20 10 。
```
   
如果defer携带的表达式语句代表的是对匿名函数的调用，那么我们就一定要非常警惕。请看下面的示例：
```go
func deferIt4() {
    for i := 1; i < 5; i++ {
        defer func() {
            fmt.Print(i)
        }()
    }
}
// 这里不对，实验也是4321？？？？？？
``` 
deferIt4函数在被执行之后标出输出上会出现5555，而不是4321。原因是defer语句携带的表达式语句中的那个匿名函数包含了对外部（确切地说，是该defer语句之外）的变量的使用。注意，等到这个匿名函数要被执行（且会被执行4次）的时候，包含该defer语句的那条for语句已经执行完毕了。此时的变量i的值已经变为了5。因此该匿名函数中的打印函数只会打印出5。正确的用法是：把要使用的外部变量作为参数传入到匿名函数中。修正后的deferIt4函数如下：
```go
func deferIt4() {
    for i := 1; i < 5; i++ {
        defer func(n int) {
            fmt.Print(n)
        }(i)
    }
}
```

### Panic
有些错误只能在运行时检查，如数组访问越界、空指针引用等。这些运行时错误会引起painc异常。当panic异常发生时，程序会中断运行，并立即执行在该goroutine中被延迟的函数(defer机制)。

`panic会停掉当前正在执行的程序（注意，不只是协程），但是与os.Exit(-1)这种直愣愣的退出不同，panic的撤退比较有秩序，他会先处理完当前goroutine已经defer挂上去的任务，执行完毕后再退出整个程序。panic仅保证当前goroutine下的defer都会被调到，但不保证其他协程的defer也会调到。`

直接调用内置的`panic函数`也会引发panic异常，panic函数接受任何值作为参数。参数通常是将出错的信息以字符串的形式来表示。panic会打印这个字符串，以及触发panic的调用栈。当某些不应该发生的场景发生时，我们就应该调用panic。

```go
package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    defer fmt.Println("defer main") // will this be printed when panic?
    var user = os.Getenv("USER_")
    go func() {
        defer fmt.Println("defer caller")
        func() {
            defer func() {
                fmt.Println("defer here")
            }()

            if user == "" {
                panic("should set user env.")
            }
        }()
    }()
    time.Sleep(1 * time.Second)
    panic("main") 
}
//defer here  
//defer caller  
//main中增加一个defer，但会睡1s，在这个过程中panic了，还没等到main睡醒，进程已经退出了，因此main的defer不会被调到；而跟panic同一个goroutine的”defer caller”还是会打印的，并且其打印在”defer here”之后。
```

### Recover
有时可以从异常中恢复，至少可以在程序崩溃前，做一些操作。例如当web服务器遇到问题时，在崩溃前应该将所有的连接关闭，否则会使得客户端一直于等待状态。

如果在deferred函数中调用了内置函数recover，并且定义该defer语句的函数发生了panic异常， recover会使程序从panic中恢复，并返回panic value。导致panic异常的函数不会继续运行，但能正常返回。注意 `recover只在defer的函数中有效，如果不是在refer上下文中调用，或在未发生panic时调用recover，recover会返回nil（选择性recover）`。
```go
package main

import (
    "os"
    "fmt"
    "time"
)

func main() {
    defer fmt.Println("defer main") // will this be called when panic?
    var user = os.Getenv("USER_")
    go func() {
        defer func() {
            fmt.Println("defer caller")
            if err := recover(); err != nil {
                fmt.Println("recover success.")
            }
        }()
        func() {
            defer func() {
                fmt.Println("defer here")
            }()

            if user == "" {
                panic("should set user env.")
            }
            fmt.Println("after panic")
        }()
    }()

    time.Sleep(1 * time.Second)
}
//defer here
//defer caller
//recover success.
//defer main
//recover力挽狂澜，避免了因为panic导致的节节败退，最终main拿到了结果，有秩序的退场了。注意，panic后面的”after panic”字符串并没有打印，这正是我们想要的。遇到解决不了的难题，只管甩panic就行，外面总有人能接的住。
```

panic可被意译为运行时恐慌。因为它只有在程序运行的时候才会被“抛出来”。并且，恐慌是会被扩散的。当有运行时恐慌发生时，它会被迅速地向调用栈的上层传递。如果我们不显式地处理它的话，程序的运行瞬间就会被终止 -- 程序崩溃。内建函数panic可以让我们人为地产生一个运行时恐慌。不过，这种致命错误是可以被恢复的。在Go语言中，内建函数recover就可以做到这一点。
  
实际上，内建函数panic和recover是天生的一对。前者用于产生运行时恐慌，而后者用于“恢复”它。不过要注意，recover函数必须要在defer语句中调用才有效。因为一旦有运行时恐慌发生，当前函数以及在调用栈上的所有代码都是失去对流程的控制权。只有defer语句携带的函数中的代码才可能在运行时恐慌迅速向调用栈上层蔓延时“拦截到”它。这里有一个可以起到此作用的defer语句的示例：

```go
defer func() {
    if p := recover(); p != nil {
        fmt.Printf("Fatal error: %s\n", p)
    }
}()
```

在这条defer语句中，我们调用了recover函数。该函数会返回一个interface{}类型的值。interface{}代表空接口。Go语言中的任何类型都是它的实现类型。我们把这个值赋给了变量p。如果p不为nil，那么就说明当前确有运行时恐慌发生。这时我们需根据情况做相应处理。注意，一旦defer语句中的recover函数调用被执行了，运行时恐慌就会被恢复，不论我们是否进行了后续处理。所以，我们一定不要只“拦截”不处理。
  
我们下面来反观panic函数。该函数可接受一个interface{}类型的值作为其参数。也就是说，我们可以在调用panic函数的时候可以传入任何类型的值。不过，我建议大家在这里只传入error类型的值。这样它表达的语义才是精确的。更重要的是，当我们调用recover函数来“恢复”由于调用panic函数而引发的运行时恐慌的时候，得到的值正是调用后者时传给它的那个参数。因此，有这样一个约定是很有必要的。
  
总之，运行时恐慌代表程序运行过程中的致命错误。我们只应该在必要的时候引发它。人为引发运行时恐慌的方式是调用panic函数。recover函数是我们常会用到的。因为在通常情况下，我们肯定不想因为运行时恐慌的意外发生而使程序崩溃。最后，在“恢复”运行时恐慌的时候，大家一定要注意处理措施的得当。

## 练习：

下面这段代码输出的内容：
```go
package main

import "fmt"

func main() {
    defer_call()
}

func defer_call()  {
    defer func() {fmt.Println("打印前")}()
    defer func() {fmt.Println("打印中")}()
    defer func() {fmt.Println("打印后")}()

    panic("触发异常")
}
// 打印后
// 打印中
// 打印前
// panic: 触发异常

// defer的执行顺序是先进后出。出现panic语句的时候，会先按照 `defer` 的后进先出顺序执行，最后才会执行panic。
```

下面这段代码输出什么?
```go
func hello(i int) {  
    fmt.Println(i)
}

func main() {  
    i := 5
    defer hello(i)
    i = i + 10
}
// 5
// hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用，所以是 5。
```

下面这段代码输出什么？
```go
func f(n int) (r int) {
    defer func() {
        r += n
        recover()
    }()

    var f func()

    defer f()
    f = func() {
        r += 2
    }
    return n + 1
}

func main() {
    fmt.Println(f(3))
}
// 7
// 第一步执行`r = n +1`，接着执行第二个 defer，由于此时 f() 未定义，引发异常，随即执行第一个 defer，异常被 recover()，程序正常执行，最后 return。
```

f1()、f2()、f3() 函数分别返回什么？
```go
func f1() (r int) {
    defer func() {
        r++
    }()
    return 0
}

func f2() (r int) {
    t := 5
    defer func() {
        t = t + 5
    }()
    return t
}

func f3() (r int) {
    defer func(r int) {
        r = r + 5
    }(r)
    return 1
}
// 1 5 1
// f1 返回值有名字，先执行 r= 0，然后执行 defer 函数，r++，最后返回 1。
// f2 返回值有名字，先执行 r= 5，然后执行 defer 函数，t+= 5，但t不影响r, 最后返回 5。
// f3 返回值有名字，先执行 r= 1, defer 函数中的r和返回值的r不一样，defer中的r += 5 不影响返回值，返回 1。如果defer 中没有传入参数，
defer func() { 
    r = r + 5 
} ()
// 则返回6.
```

下面代码输出什么？
```go
func increaseA() int {
    var i int
    defer func() {
        i++
    }()
    return i
}

func increaseB() (r int) {
    defer func() {
        r++
    }()
    return r
}

func main() {
    fmt.Println(increaseA())
    fmt.Println(increaseB())
}
// 0 1
// 结合上题，注意一下，increaseA() 的返回参数是匿名，返回i= 0, defer中i++不影响返回值。
// increaseB() 是具名，先执行 r = 0, defer中r++，最后返回 1。
```

下面代码段输出什么？
```go
type Person struct {
    age int
}

func main() {
    person := &Person{28}
    // 1. 
    defer fmt.Println(person.age)
    // 2.
    defer func(p *Person) {
        fmt.Println(p.age)
    }(person)  
    // 3.
    defer func() {
        fmt.Println(person.age)
    }()
    person.age = 29
}
// 29 29 28
// 变量 person 是一个指针变量 。
// 1.person.age 此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28；
// 2.defer 缓存的是结构体 Person{28} 的地址，最终 Person{28} 的 age 被重新赋值为 29，所以 defer 语句最后执行的时候，依靠缓存的地址取出的 age 便是 29，即输出 29；
// 3.闭包引用，没有参数传进去，所以最后执行是更新过的person. 输出 29；
// 又由于 defer 的执行顺序为先进后出，即 3 2 1，所以输出 29 29 28。
```

下面代码输出什么？

```go
type Person struct {
    age int
}

func main() {
    person := &Person{28}
    // 1.
    defer fmt.Println(person.age)
    // 2.
    defer func(p *Person) {
        fmt.Println(p.age)
    }(person)
    // 3.
    defer func() {
        fmt.Println(person.age)
    }()
    person = &Person{29}
}
// 29 28 28
// 前一题最后一行代码 `person.age = 29` 是修改引用对象的成员 age，这题最后一行代码 `person = &Person{29}` 是修改引用对象本身，来看看有什么区别。
// 1处.person.age 这一行代码跟之前含义是一样的，此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28；
// 2处.defer 缓存的是结构体 Person{28} 的地址，这个地址指向的结构体没有被改变，最后 defer 语句后面的函数执行的时候取出仍是 28；
// 3处.闭包引用，person 的值已经被改变，指向结构体 `Person{29}`，所以输出 29.
// 由于 defer 的执行顺序为先进后出，即 3 2 1，所以输出 29 28 28。
```

下面这段代码正确的输出是什么？
```go
func f() {
    defer fmt.Println("D")
    fmt.Println("F")
}

func main() {
    f()
    fmt.Println("M")
}
// F D M
// 被调用函数里的 defer 语句在返回之前就会被执行，所以输出顺序是 F D M。
```

return 之后的 defer 语句会执行吗，下面这段代码输出什么？
```go
var a bool = true
func main() {
    defer func(){
        fmt.Println("1")
    }()
    if a == true {
        fmt.Println("2")
        return
    }
    defer func(){
        fmt.Println("3")
    }()
}
// 2 1
// defer 关键字后面的函数或者方法想要执行必须先注册，return 之后的 defer 是不能注册的， 也就不能执行后面的函数或方法。
```

下面的代码输出什么？
```go
func F(n int) func() int {
    return func() int {
        n++
        return n
    }
}

func main() {
    f := F(5)
    defer func() {
        fmt.Println(f()) // 3rd call
    }()
    defer fmt.Println(f())   //1st call
    i := f()  // 2nd call
    fmt.Println(i)
}
// 7 6 8
// defer() 后面的函数如果带参数，会优先计算参数，并将结果存储在栈中，到真正执行 defer() 的时候取出。
// 另外 f 一直操作的是5所在的指针，所以 f() 的值会累加。

//帮助理解
func F(n int) func() int {
	return func() int {
		n++
		fmt.Println(n, &n)
		return n
	}
}

func main() {
	f := F(5)
	defer func() {
		fmt.Println(f())
	}()
	defer fmt.Println(f())
	i := f()
	fmt.Println(i)
}
/*
6 0xc000012028
7 0xc000012028
7
6
8 0xc000012028
8
*/
```

下面列举的是 recover() 的几种调用方式，哪些是正确的？
```go
// A
func main() {
    recover()
    panic(1)
}

// B
func main() {
    defer recover()
    panic(1)
}

// C
func main() {
    defer func() {
        recover()
    }()
    panic(1)
}

// D
func main() {
    defer func() {
        defer func() {
            recover()
        }()
    }()
    panic(1)
}
// 答：C
// recover() 必须在 defer() 函数中直接调用才有效。上面其他几种情况调用都是无效的：直接调用 recover()、在 defer() 中直接调用 recover() 和 defer() 调用时多层嵌套。
```

下面代码输出什么，请说明？
```go
func main() {
    defer func() {
        fmt.Print(recover())
    }()
    defer func() {
        defer fmt.Print(recover())
        panic(1)
    }()
    defer recover() 
    panic(2)
}
// 2 1
// recover() 必须在 defer() 函数中调用才有效，所以第 9 行代码捕获是无效的。在调用 defer() 时，便会计算函数的参数并压入栈中，所以执行第 6 行代码时，此时便会捕获 panic(2)；此后的 panic(1)，会被上一层的 recover() 捕获。所以输出 21。

func main() {
	defer func() {
		fmt.Print("1-", recover())
	}()
	defer func() {
		defer fmt.Print("2-", recover())
		panic(1)
	}()
	panic(2)
}
// 帮助理解 2-21-1

// 但如果是这样：
func main() {
    defer func() {
        fmt.Print("2 - ", recover())
    }()
    defer func() {
        defer func() {
            fmt.Print("1 - ", recover())
        }()
        panic(1)
    }()
    defer recover()
    panic(2)
}
// 输出就是 1- 12- 2
// panic 2 -- panic 1 -- defer嵌套中的 recover捕获panic 1, 然后最上面的defer捕获panic 2。
```

下面代码输出什么，请说明原因。
```go
type Slice []int

func NewSlice() Slice {
    return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
    *s = append(*s, elem)
    fmt.Print(elem)
    return s
}
func main() {
    s := NewSlice()
    defer s.Add(1).Add(2)
    s.Add(3)
}
// 1 3 2
// 1. Add() 方法的返回值依然是指针类型 *Slice，所以可以循环调用方法 Add()；
// 2. defer 函数的参数（包括接收者）是在 defer 语句出现的位置做计算的，而不是在函数执行的时候计算的，所以 s.Add(1) 会先于 s.Add(3) 执行。

// 相当于这个代码：
func main() {
	s := NewSlice()
	defer func(s *Slice) *Slice {
		return s.Add(2)
	}(s.Add(1))
	s.Add(3)
}

// 如果是：
func main() {
	s := NewSlice()
	defer func(s *Slice) *Slice {
		return s
	}(s.Add(1).Add(2))
	s.Add(3)
} // 123.

// 如果是：
func main() {
	s := NewSlice()
	defer func() *Slice {
		return s.Add(1).Add(2)
	}()
	s.Add(3)
} // 312.
```

函数执行时，如果由于 panic 导致了异常，则延迟函数不会执行。这一说法是否正确？ 答：false
由 panic 引发异常以后，程序停止执行，然后调用延迟函数（defer），就像程序正常退出一样。

下面代码输出什么，为什么？
```go
func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("recover:%#v", r)
        }
    }()
    panic(1)
    panic(2)
}

func main() {
    f()
}
// recover:1. IDE会提示：`unreachable code`
// 当程序 panic 时就不会往下执行，可以使用 recover() 捕获 panic 的内容。
```

下面的代码有什么问题，请说明？
```go
func main() {
    f, err := os.Open("file")
    defer f.Close()
    if err != nil {
        return
    }

    b, err := ioutil.ReadAll(f)
    println(string(b))
}
// defer 语句应该放在 if() 语句后面，先判断 err，再 defer 关闭文件句柄。
// 修复代码：
func main() {
    f, err := os.Open("file")
    if err != nil {
        return
    }
    defer f.Close()

    b, err := ioutil.ReadAll(f)
    println(string(b))
}
```

下面代码能编译通过吗？可以的话，输出什么？
```go
func main() {
    println(DeferTest1(1))
    println(DeferTest2(1))
}

func DeferTest1(i int) (r int) {
    r = i
    defer func() {
        r += 3
    }()
    return r
}

func DeferTest2(i int) (r int) {
    defer func() {
        r += i
    }()
    return 2
}
// 4 3
// 都是具名返回值。
```

下面这段代码输出什么？
```go
func main() {
    a := 1
    b := 2
    defer calc("1", a, calc("10", a, b))
    a = 0
    defer calc("2", a, calc("20", a, b))
    b = 1
}

func calc(index string, a, b int) int {
    ret := a + b
    fmt.Println(index, a, b, ret)
    return ret
}
/*
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4

程序执行到 main() 函数三行代码的时候，会先执行 calc() 函数的 b 参数，即：`calc("10",a,b)`，输出：10 1 2 3，得到值 3，因为defer 定义的函数是延迟函数，故 calc("1",1,3) 会被延迟执行；
程序执行到第五行的时候，同样先执行 calc("20",a,b) 输出：20 0 2 2 得到值 2，同样将 calc("2",0,2) 延迟执行；
程序执行到末尾的时候，按照栈先进后出的方式依次执行：calc("2",0,2)，calc("1",1,3)，则就依次输出：2 0 2 2，1 1 3 4。*/
```

关于异常的触发，下面说法正确的是？ 答：A B C D

- A. 空指针解析；
- B. 下标越界；
- C. 除数为0；
- D. 调用panic函数；


下面代码输出什么，请说明。 -- TBD
```go
func main() {
    x := []int{0, 1, 2}
    y := [3]*int{}
    for i, v := range x {
        defer func() {
            print(v)
        }()
        y[i] = &v
    }
    print(*y[0], *y[1], *y[2])
}
```

**答：012210     22222**

**解析：**

知识点：defer()、for-range。

~~for-range 虽然使用的是 :=，但是 v 不会重新声明，可以打印 v 的地址验证下。~~


下面两组代码输出：
```go
func b() (i int) { 	
    defer func() { 		
        i++ 		
        fmt.Println("defer2:", i) 	
    }() 	
    defer func() { 		
        i++ 		
        fmt.Println("defer1:", i) 	
    }() 	
    return 
} 
func main() { 	
    fmt.Println("return:", b()) 
} 
// defer 1
// defer 2
// return 2
```

```go
func c() *int { 	
    var i int 	
    defer func() { 		
        i++ 		
        fmt.Println("defer2:", i) 	
    }() 	
    defer func() { 		
        i++ 		
        fmt.Println("defer1:", i) 	
    }() 	
    return &i 
} 
func main() { 	
    fmt.Println("return:", *(c())) 
}
// defer 1
// defer 2
// return 2
```

输出：

```go
func main() {
    {
        defer fmt.Println("defer runs")
        fmt.Println("block ends")
    }

    fmt.Println("main ends")
}
// block ends
// main ends
// defer runs
// 不是在当前代码块的作用域时执行的, `defer`只会在当前函数和方法返回之前被调用
```

输出：

```go
type Test struct {
    value int
}

func (t Test) print() {
    println(t.value)
}

func main() {
    test := Test{}
    defer test.print()
    test.value += 1
}
// 0

type Test struct {
    value int
}

func (t *Test) print() {
    println(t.value)
}

func main() {
    test := Test{}
    defer test.print()
    test.value += 1
}
// 1
// 还是进行的值传递,不过发生复制的是指向`test`的指针,上面那个复制的是结构体,这段是复制的指针,修改`test.value`时,`defer`捕获的指针其实就能够访问到修改后的变量了
```