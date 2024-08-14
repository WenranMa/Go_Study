# Go_闭包

闭包（Closure）是指一个函数包含了它外部作用域中的变量，即使在外部作用域结束后，这些变量依然可以被内部函数访问和修改。闭包使得函数可以“记住”外部作用域的状态，这种状态在函数调用之间是保持的。

闭包的核心概念是函数内部可以引用外部作用域的变量，即使在函数内部外部作用域已经结束。

特点：
- 函数可以在定义的作用域之外被调用，仍然可以访问外部作用域的变量。
- 外部作用域中的变量不会被销毁，直到闭包不再引用它们。
- 多个闭包可以共享同一个外部作用域的变量。

## 原理
Go语言中的闭包是通过**函数值（Function Value）**实现的。在Go语言中，函数不仅是代码，还是数据，可以像其他类型的值一样被传递、赋值和操作。当一个函数内部引用了外部作用域的变量时，Go编译器会生成一个闭包实例，将外部变量的引用与函数代码绑定在一起。

示例：基本闭包
```go
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main(){
    counter := makeCounter()
    fmt.Println(counter()) // 输出 1
    fmt.Println(counter()) // 输出 2
    
    c2 := makeCounter()
    fmt.Println(c2()) // 输出 1
    fmt.Println(c2()) // 输出 2
}
```
在上述示例中，makeCounter 函数返回一个匿名函数，这个匿名函数持有了外部变量 count 的引用。每次调用 counter() 时，都会访问和修改外部作用域的 count 变量。
并且每个闭包函数的引用都会在自己的函数内部保存一份闭包变量 varInner，这样在调用过程中就不会互相影响。

## 应用场景

**状态保持和共享** 闭包常用于实现状态保持和共享。通过闭包，我们可以在函数调用之间保持状态，而无需使用全局变量。

```go
func makeAccumulator() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	acc := makeAccumulator()
	fmt.Println(acc(5)) // 输出 5
	fmt.Println(acc(3)) // 输出 8
}
```

**函数式编程** 在函数式编程中，函数可以作为参数传递、返回值和变量赋值。闭包使得函数可以更加灵活地用于函数式编程，实现函数的组合和转换。

```go
func mapInts(nums []int, f func(int) int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = f(num)
	}
	return result
}

func main() {
	double := func(x int) int {
		return x * 2
	}
	nums := []int{1, 2, 3, 4}
	doubledNums := mapInts(nums, double)
	fmt.Println(doubledNums)
}
// [2,4,6,8]
```

**并发编程** 在并发编程中，闭包将状态隔离在每个goroutine中，避免竞态条件和数据不一致问题。

```go
func startWorker(id int) {
    go func() {
        for {
            fmt.Printf("Worker %d is working\n", id)
            time.Sleep(time.Second)
        }
    }()
}

func main() {
    for i := 0; i < 3; i++ {
        startWorker(i)
    }
    // 主goroutine 继续执行其他操作
}
```

## 注意事项

**内存泄漏** 由于闭包持有外部作用域的变量引用，如果闭包一直被引用，外部作用域的变量不会被销毁，可能会导致内存泄漏。在使用闭包时，需要注意外部作用域变量的生命周期。

**竞态条件** 在并发编程中，由于多个goroutine可以共享闭包中的变量，可能会引发竞态条件和数据不一致问题。在并发场景下使用闭包时，需要保证变量的访问是安全的。
