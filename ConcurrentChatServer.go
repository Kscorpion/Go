package main

import (
	"net"
	"fmt"
	"strings"
	"time"
)

type Client struct {
	C chan string //用于发送数据的管道
	Name string //用户名
	Addr string //网络地址
}

//保存在线用户 cliAddr==>
var onlineMap map[string]Client

//
var message = make(chan string)

func WriteMsgToClient(cli Client,conn net.Conn){
	for msg:=range cli.C {//给当前客户端发送消息
		conn.Write([]byte(msg))
	}
}

func MakeMsg(cli Client,msg string)(buf string){
	message<- "["+cli.Addr+"]"+cli.Name+":"+msg
	return
}
func HandleConn(conn net.Conn){
	defer conn.Close()
	//获取客户端网络地址
	cliAddr :=conn.RemoteAddr().String()
	//创建一个结构体 默认用户名和网络地址一样
	cli := Client{make(chan string),cliAddr,cliAddr}
	//把结构体添加到map
	onlineMap[cliAddr] = cli
	//新开一个协程专门给客户端发送信息
	go WriteMsgToClient(cli,conn)
	//广播某个人在线
	//message <-MakeMsg(cli,"login")
	//提示我是谁
	message <- MakeMsg(cli,"I am comming !")
	isQuit := make(chan bool)//用于判断用户是否主动退出
	hasData := make(chan bool)//用于判断用户是否有数据发送
	//新开一个协程接收用户发送过来的数据
	go func() {
		buf := make([]byte,2048)
		for {
			n,err := conn.Read(buf)
			if n== 0{//对方端口或出问题
				isQuit <- true
				fmt.Println("conn.Read err = ",err)
				return
			}
			msg:=string(buf[:n-2])
			//fmt.Println(len(msg))
			if len(msg) == 3 && msg == "who"{
				conn.Write([]byte("user list:"))
				//遍历列表给当前用户发送所有成员
				for _,tmp := range onlineMap{
					msg = tmp.Addr+":"+tmp.Name+"\n"
					conn.Write([]byte(msg))
				}
			}else if len(msg)>=8 && msg[:6]=="rename"{
				name := strings.Split(msg,"|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename success"))
			}else {
				//转发此内容
				message<-MakeMsg(cli,msg)
			}
			hasData<-true
		}

	}()
	for{
		//通过select 检测chan 的流动
		select {
			case <-isQuit:
				delete(onlineMap,cliAddr)//当前用户从map移除
				message<-MakeMsg(cli,"login out") //广播谁下线了
				return
		case <-hasData://有数据不做任何处理

		case <-time.After(60*time.Second):
			delete(onlineMap,cliAddr)
			message<-MakeMsg(cli,"time out leave out")//超时退出
			return
		}
	}
}

func Manager(){
	//给map分配空间
	onlineMap = make(map[string]Client)
	for{
		msg := <-message //没有消息前这里会阻塞
		for _,cli :=range onlineMap  {
			cli.C <-msg
		}
	}
}

func main() {
	//监听
	Listener,err := net.Listen("tcp","127.0.0.1:9494")
	if err != nil{
		fmt.Println("net.Listen err = ",err)
		return
	}
	defer Listener.Close()
	//新开一个协程用于转发消息，只要有消息来就遍历map，给每个成员都发送消息
	go Manager()
	//主协程循环等待
	for  {
		conn,err:= Listener.Accept()
		if err != nil{
			fmt.Println("Listener.Accept err = ",err)
			continue
		}
		go HandleConn(conn)//处理用户连接
	}

}
