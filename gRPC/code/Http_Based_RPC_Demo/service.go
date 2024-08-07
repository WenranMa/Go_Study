// rpc demo/service.go

package main

import "fmt"

type Args struct {
	X, Y int
}

// ServiceA 自定义一个结构体类型
type ServiceA struct{}

// Add 为ServiceA类型增加一个可导出的Add方法
func (s *ServiceA) Add(args *Args, reply *int) error {
	fmt.Println("Add function called: ", args)
	*reply = args.X + args.Y
	return nil
}
