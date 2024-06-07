
# Defer 相关 

## 下面这段代码输出的内容：

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
```

**答：**

```shell
打印后
打印中
打印前
panic: 触发异常
```

**解析：**
`defer` 的执行顺序是先进后出。出现panic语句的时候，会先按照 `defer` 的后进先出顺序执行，最后才会执行panic。


## 32. 下面这段代码输出什么?

```go
func hello(i int) {  
    fmt.Println(i)
}

func main() {  
    i := 5
    defer hello(i)
    i = i + 10
}
```

**答：5**

**解析：**

这个例子中，hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用，所以是 5。


## 69. 下面这段代码输出什么？ -- TBD
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
```

**答：7**

**解析：**

第一步执行`r = n +1`，接着执行第二个 defer，由于此时 f() 未定义，引发异常，随即执行第一个 defer，异常被 recover()，程序正常执行，最后 return。


## 46. f1()、f2()、f3() 函数分别返回什么？ -- TBD

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
```

**答：1 5 1**

**解析：**

知识点：`defer`函数的执行顺序，结合44题。

## 44. 下面代码输出什么？ -- TBD

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
```

- A. 1 1
- B. 0 1
- C. 1 0
- D. 0 0

**答：B**

**解析：**

知识点：defer、返回值。

注意一下，increaseA() 的返回参数是匿名，increaseB() 是具名。


## 47. 下面代码段输出什么？

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
```

**答：29 29 28**

**解析：**

知识点：`defer`函数的执行顺序，结合1，46题。

变量 person 是一个指针变量 。

1.person.age 此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28；

2.defer 缓存的是结构体 Person{28} 的地址，最终 Person{28} 的 age 被重新赋值为 29，所以 defer 语句最后执行的时候，依靠缓存的地址取出的 age 便是 29，即输出 29；

3.闭包引用，输出 29；

又由于 defer 的执行顺序为先进后出，即 3 2 1，所以输出 29 29 28。

## 48. 下面这段代码正确的输出是什么？

```go
func f() {
    defer fmt.Println("D")
    fmt.Println("F")
}

func main() {
    f()
    fmt.Println("M")
}
```

- A. F M D
- B. D F M
- C. F D M

**答：C**

**解析：**

被调用函数里的 defer 语句在返回之前就会被执行，所以输出顺序是 F D M，结合1，47题。

## 49. 下面代码输出什么？

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
```

**答：29 28 28**

**解析：**

知识点：`defer`函数的执行顺序。

这道题在第 `47` 题目的基础上做了一点点小改动，前一题最后一行代码 `person.age = 29` 是修改引用对象的成员 age，这题最后一行代码 `person = &Person{29}` 是修改引用对象本身，来看看有什么区别。

1处.person.age 这一行代码跟之前含义是一样的，此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28；

2处.defer 缓存的是结构体 Person{28} 的地址，这个地址指向的结构体没有被改变，最后 defer 语句后面的函数执行的时候取出仍是 28；

3处.闭包引用，person 的值已经被改变，指向结构体 `Person{29}`，所以输出 29.

由于 defer 的执行顺序为先进后出，即 3 2 1，所以输出 29 28 28。


## 54. return 之后的 defer 语句会执行吗，下面这段代码输出什么？

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
```

**答：2 1**

**解析：**

defer 关键字后面的函数或者方法想要执行必须先注册，return 之后的 defer 是不能注册的， 也就不能执行后面的函数或方法。

## 150. 下面的代码输出什么？

```go
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
```

**答：768**

**解析：**

知识点：`匿名函数`、`defer()`。

defer() 后面的函数如果带参数，会优先计算参数，并将结果存储在栈中，到真正执行 defer() 的时候取出。

```go
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

## 151. 下面列举的是 recover() 的几种调用方式，哪些是正确的？

* A

  ```go
  func main() {
      recover()
      panic(1)
  }
  ```

* B

  ```go
  func main() {
      defer recover()
      panic(1)
  }
  ```

* C

  ```go
  func main() {
      defer func() {
          recover()
      }()
      panic(1)
  }
  ```

* D

  ```go
  func main() {
      defer func() {
          defer func() {
              recover()
          }()
      }()
      panic(1)
  }
  ```

**答：C**

**解析：**

recover() 必须在 defer() 函数中直接调用才有效。上面其他几种情况调用都是无效的：直接调用 recover()、在 defer() 中直接调用 recover() 和 defer() 调用时多层嵌套。

# 152. 下面代码输出什么，请说明？ -- TBD

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
```

**答：21**

**解析：**

recover() 必须在 defer() 函数中调用才有效，所以第 9 行代码捕获是无效的。在调用 defer() 时，便会计算函数的参数并压入栈中，所以执行第 6 行代码时，此时便会捕获 panic(2)；此后的 panic(1)，会被上一层的 recover() 捕获。所以输出 21。

```go
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
```


## 154. 下面的代码输出什么，请说明？-- TBD

```go
func main() {
    defer func() {
        fmt.Print(recover())
    }()
    defer func() {
        defer func() {
            fmt.Print(recover())
        }()
        panic(1)
    }()
    defer recover()
    panic(2)
}
```

**答：12**

**解析：**

152题与之类似

## 160. 下面代码输出什么，请说明。

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

## 166. 下面代码输出什么，请说明原因。

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
```

**答：132**

**解析：**

这一题有两点需要注意：

1. Add() 方法的返回值依然是指针类型 *Slice，所以可以循环调用方法 Add()；
2. defer 函数的参数（包括接收者）是在 defer 语句出现的位置做计算的，而不是在函数执行的时候计算的，所以 s.Add(1) 会先于 s.Add(3) 执行。

```go
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

## 167. 下面的代码输出什么，请说明。

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
    defer func() {
        s.Add(1).Add(2)
    }()
    s.Add(3)
}
```

**答：312**

**解析：**

对比昨天的第`166`题，本题的 s.Add(1).Add(2) 作为一个整体包在一个匿名函数中，会延迟执行。

## 191. 函数执行时，如果由于 panic 导致了异常，则延迟函数不会执行。这一说法是否正确？

- A. true
- B. false

**答：B**

**解析：**

由 panic 引发异常以后，程序停止执行，然后调用延迟函数（defer），就像程序正常退出一样。



## 174. 下面代码输出什么，为什么？

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
```

**答：recover:1**

编译错误：`unreachable code`

**解析：**

知识点：`panic`、`recover()`。

当程序 panic 时就不会往下执行，可以使用 recover() 捕获 panic 的内容。

## 173. 下面的代码有什么问题，请说明？

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
```

**答：defer 语句应该放在 if() 语句后面，先判断 err，再 defer 关闭文件句柄。**

**解析：**

修复代码：

```go
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


## 164. 下面代码能编译通过吗？可以的话，输出什么？

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
```

**答：43**

**解析：**

都是具名返回值。


## 58. 下面这段代码输出什么？

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
```

**答：**

```shell
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
```

**解析：**

程序执行到 main() 函数三行代码的时候，会先执行 calc() 函数的 b 参数，即：`calc("10",a,b)`，输出：10 1 2 3，得到值 3，因为
defer 定义的函数是延迟函数，故 calc("1",1,3) 会被延迟执行；

程序执行到第五行的时候，同样先执行 calc("20",a,b) 输出：20 0 2 2 得到值 2，同样将 calc("2",0,2) 延迟执行；

程序执行到末尾的时候，按照栈先进后出的方式依次执行：calc("2",0,2)，calc("1",1,3)，则就依次输出：2 0 2 2，1 1 3 4。
