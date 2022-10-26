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
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			for _, b := range this.Trucks {
				if er := stream.Send(b); er != nil {
					return er
				}
			}
		}
		if err != nil {
			return err
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
