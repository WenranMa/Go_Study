package main

import (
	"container/ring"
	"fmt"
)

func main() {
	ring := ring.New(5)
	fmt.Println(ring)
}
