package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

//rpc协议标准
//1.字母必须大写
//2.必须要有2个导出类型的参数
//3.函数必须要有返回值
type Args struct {
	A, B int //两个参数
}

type Query struct {
	X, Y int //结果
}
type Last int //这里需要大写

//乘法
func (t *Last) Multiply(args Args, reply *int) error {
	*reply = args.A * args.B // 乘法
	fmt.Println(reply, "乘法已执行")
	return nil
}

//除法
func (t *Last) Divide(args Args, query *Query) error {
	if args.B == 0 {
		return errors.New("不能除以0")
	}
	query.X = args.A / args.B
	query.Y = args.A % args.B
	fmt.Println(query, "除法已执行")
	return nil
}
func main() {
	la := new(Last)
	fmt.Println(la, "=la")
	rpc.Register(la) //注册类型
	rpc.HandleHTTP() //设定http类型
	listen, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	http.Serve(listen, nil)
}
