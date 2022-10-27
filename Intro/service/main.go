package main

import (
	pb "Intro/service/Test"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"net"
	"strconv"
	"time"
)

type server struct {
	pb.UnimplementedOrderSendServer
	Trucks map[string]*pb.Car
}

func (this *server) OrderDely(stream pb.OrderSend_OrderDelyServer) error {
	conn := make(chan *pb.OrderInfor, 5)
	go Listen(stream, conn)
	return this.Return(stream, conn)
}
func main() {
	trucks := make(map[string]*pb.Car)
	conn, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Listen error : ", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderSendServer(s, &server{Trucks: trucks})
	err = s.Serve(conn)
	if err != nil {
		fmt.Println("Serve error : ", err)
	}
}
func Listen(stream pb.OrderSend_OrderDelyServer, pip chan *pb.OrderInfor) {
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		pip <- order
	}
	close(pip)
}
func (this *server) Return(stream pb.OrderSend_OrderDelyServer, pip chan *pb.OrderInfor) error {
	for {
		order, ok := <-pip
		if ok == false {
			for _, b := range this.Trucks {
				if er := stream.Send(b); er != nil {
					fmt.Println("[server]Send error : ", er)
					return er
				} else {
					continue
				}
			}
		}
		id, _ := strconv.Atoi(order.Id)
		if a, b := this.Trucks[order.Item]; b {
			a.Orders[id] = order
		} else {
			this.Trucks[order.Item] = new(pb.Car)
			this.Trucks[order.Item].Orders = make([]*pb.OrderInfor, 10)
			this.Trucks[order.Item].Item = order.Item
			this.Trucks[order.Item].Orders[id] = order
			this.Trucks[order.Item].Carid = strconv.Itoa(time.Now().Second())
		}
	}
}
