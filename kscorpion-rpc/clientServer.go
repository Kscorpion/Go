package main

import (
	"fmt"
	"net/rpc"
)

type ArgsX struct {
	A, B int //两个参数
}

type QueryX struct {
	X, Y int //结果
}

func main() {
	servepIp := "127.0.0.1:1234"
	client, err := rpc.DialHTTP("tcp", servepIp)
	if err != nil {
		fmt.Println(err)
	}
	a := 10
	b := 5
	args := &ArgsX{a, b}
	var reply int
	err = client.Call("Last.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(args.A, args.B, reply) //乘法
	var qu QueryX
	err = client.Call("Last.Divide", args, &qu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(args.A, args.B, qu) //除法
}
