package main

import (
	"client/pb"
	"fmt"
	"github.com/Lubby-ch/protorpc"
	"log"
)

//// 算数运算结构体
//type Arith struct {
//}
//
//// 乘法运算方法
//func (this *Arith) Multiply(req *pb.ArithRequest, res *pb.ArithResponse) error {
//	res.Pro = req.A * req.B
//	return nil
//}
//
//// 除法运算方法
//func (this *Arith) Divide(req *pb.ArithRequest, res *pb.ArithResponse) error {
//	if req.B == 0 {
//		return errors.New("divide by zero")
//	}
//	res.Quo = req.A / req.B
//	res.Rem = req.A % req.B
//	return nil
//}

func main() {
	conn, err := protorpc.Dial("tcp", "127.0.0.1:8096")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}

	req := &pb.ArithRequest{A: 9, B: 2}

	var res pb.ArithResponse


	err = conn.Call("Arith.Multiply", req, &res) // 乘法运算
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)

	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)

}

