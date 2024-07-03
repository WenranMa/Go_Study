/*
练习14.12：[general_lazy_evaluation2.go](exercises/chapter_14/general_lazy_evalution2.go)
通过使用 14.12 中工厂函数生成前 10 个斐波那契数

提示：因为斐波那契数增长很迅速，使用 `uint64` 类型。
注：这种计算通常被定义为递归函数，但是在没有尾递归的语言中，例如 go 语言，这可能会导致栈溢出，但随着 go 语言中堆栈可扩展的优化，这个问题就不那么严重。这里的诀窍是使用了惰性求值。gccgo 编译器在某些情况下会实现尾递归。

*/

package main

import (
	"fmt"
)

type Any interface{}
type EvalFunc func(Any) (Any, Any)

func main() {
	fibFunc := func(state Any) (Any, Any) {
		os := state.([]uint64)
		v1 := os[0]
		v2 := os[1]
		ns := []uint64{v2, v1 + v2}
		return v1, ns
	}
	fib := BuildLazyUInt64Evaluator(fibFunc, []uint64{0, 1})

	for i := 0; i < 10; i++ {
		fmt.Printf("Fib nr %v: %v\n", i, fib())
	}
}

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		var actState Any = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}
	retFunc := func() Any {
		return <-retValChan
	}
	go loopFunc()
	return retFunc
}

func BuildLazyUInt64Evaluator(evalFunc EvalFunc, initState Any) func() uint64 {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() uint64 {
		return ef().(uint64)
	}
}

/* Output:
Fib nr 0: 0
Fib nr 1: 1
Fib nr 2: 1
Fib nr 3: 2
Fib nr 4: 3
Fib nr 5: 5
Fib nr 6: 8
Fib nr 7: 13
Fib nr 8: 21
Fib nr 9: 34
*/
