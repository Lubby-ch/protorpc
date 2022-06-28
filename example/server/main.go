package main

import (
	"errors"
	"fmt"
	protorpc "github.com/Lubby-ch/protorpc"
	"log"
	"net"
	"net/rpc"
	"os"
	"server/pb"
)

// 算数运算结构体
type Arith struct {
}

// 乘法运算方法
func (this *Arith) Multiply(req *pb.ArithRequest, res *pb.ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// 除法运算方法
func (this *Arith) Divide(req *pb.ArithRequest, res *pb.ArithResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main() {
	server()
}

func server() {
	rpc.Register(new(Arith)) // 注册rpc服务

	lis, err := net.Listen("tcp", "127.0.0.1:8096")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}

	fmt.Fprintf(os.Stdout, "%s", "start connection")

	for {
		conn, err := lis.Accept() // 接收客户端连接请求
		if err != nil {
			continue
		}

		go func(conn net.Conn) { // 并发处理客户端请求
			fmt.Fprintf(os.Stdout, "%s", "new client in coming\n")
			protorpc.ServeConn(conn)
		}(conn)
	}
}

