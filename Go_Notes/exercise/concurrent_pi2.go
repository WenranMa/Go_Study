// concurrent_pi2.go

/*
使用以下序列在协程中计算 pi：开启一个协程来计算公式中的每一项并将结果放入通道，`main()` 函数收集并累加结果，
打印出 pi 的近似值。

计算执行时间（

再次声明这只是为了一边练习协程的概念一边找点乐子。

如果你需要的话可使用 `math.pi` 中的 `Pi`；而且不使用协程会运算的更快。一个急速版本：使用 `GOMAXPROCS`，
开启和 `GOMAXPROCS` 同样多个协程。
*/
package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

const NCPU = 2

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(2)
	fmt.Println(CalculatePi(5000))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func CalculatePi(end int) float64 {
	ch := make(chan float64)
	for i := 0; i < NCPU; i++ {
		go term(ch, i*end/NCPU, (i+1)*end/NCPU)
	}
	result := 0.0
	for i := 0; i < NCPU; i++ {
		result += <-ch
	}
	return result
}

func term(ch chan float64, start, end int) {
	result := 0.0
	for i := start; i < end; i++ {
		x := float64(i)
		result += 4 * (math.Pow(-1, x) / (2.0*x + 1.0))
	}
	ch <- result
}

/* Output:
3.1413926535917938
The calculation took this amount of time: 0.002000
*/
