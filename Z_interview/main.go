package main

import (
    "fmt"
    "time"
    "context"
    "math"
    "image/color"
)

func main()  {
    deferIt4()
    /*
    deferIt3()
    f(3)

    ctx,cancel := context.WithCancel(context.Background())
    go Speak(ctx)
    time.Sleep(10*time.Second)
    cancel()
    time.Sleep(1*time.Second)
    */

    p := Point{1, 2} 
    q := Point{4, 6}
    distanceFromP := p.Distance  //method value
    fmt.Println("1",distanceFromP(q)) 
    distance := Point.Distance  // method expression
    fmt.Println("2",distance(p, q))


    a := Person{"Robert", "Male", 33, "Beijing"}
    a.Grow()
    fmt.Printf("%v\n", a)   
}


type Person struct {
    Name    string
    Gender  string
    Age     uint8
    Address string
}

func (person Person) Grow() {
    person.Age++
}


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


func deferIt3() {
    f := func(i int) int {
        fmt.Printf("%d ",i)
        return i * 10
    }
    for i := 1; i < 5; i++ {
        defer fmt.Printf("%d ", f(i))
    }
}
func deferIt4() {
    for i := 1; i < 5; i++ {
        defer func() {
            fmt.Print(i)
        }()
    }
}

func f(x int) {
    fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
    defer fmt.Printf("defer %d\n", x)
    f(x - 1)
}
func Speak(ctx context.Context)  {
    for range time.Tick(time.Second){
        select {
        case <- ctx.Done():
            fmt.Println("我要闭嘴了")
            return
        default:
            fmt.Println("balabalabalabala")
        }
    }
}