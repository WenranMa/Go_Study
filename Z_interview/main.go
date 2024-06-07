package main

import (
	"fmt"
	"time"
)

func main() {
	f := NewFizzFuzz(20)
	go f.Fizz()
	go f.Fuzz()
	go f.FizzFuzz()
	go f.Num()

	go func() {
		for i := 1; i <= f.n; i++ {
			<-f.start
			if i%3 == 0 && i%5 == 0 {
				f.fizzfuzz <- i
			} else if i%3 == 0 {
				f.fizz <- i
			} else if i%5 == 0 {
				f.fuzz <- i
			} else {
				f.num <- i
			}
		}
	}()
	f.start <- 1
	time.Sleep(10 * time.Second)
}

type FizzFuzz struct {
	n        int
	fizz     chan int
	fuzz     chan int
	fizzfuzz chan int
	num      chan int
	start    chan int
}

func NewFizzFuzz(n int) *FizzFuzz {
	zeo := &FizzFuzz{
		n:        n,
		fizz:     make(chan int),
		fuzz:     make(chan int),
		fizzfuzz: make(chan int),
		num:      make(chan int),
		start:    make(chan int),
	}
	return zeo
}

func (f *FizzFuzz) Fizz() {
	for {
		<-f.fizz
		fmt.Println("fizz")
		f.start <- 1
	}
}

func (f *FizzFuzz) Fuzz() {
	for {
		<-f.fuzz
		fmt.Println("fuzz")
		f.start <- 1
	}
}

func (f *FizzFuzz) FizzFuzz() {
	for {
		<-f.fizzfuzz
		fmt.Println("fizzfuzz")
		f.start <- 1
	}
}

func (f *FizzFuzz) Num() {
	for {
		i := <-f.num
		fmt.Println(i)
		f.start <- 1
	}
}
