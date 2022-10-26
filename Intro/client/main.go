package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	pb "test/Intro/client/Test"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial error : ", err)
	}
	cli := pb.NewOrderSendClient(conn)
	Carstream, err := cli.OrderDely(context.Background())
	if err != nil {
		fmt.Println("OrderDely error : ", err)
	}
	Iphone := new(pb.OrderInfor)
	Iphone.Id = "0"
	Iphone.Item = "phone"
	Sumsung := new(pb.OrderInfor)
	Sumsung.Id = "3"
	Sumsung.Item = "phone"
	Apple := new(pb.OrderInfor)
	Apple.Id = "2"
	Apple.Item = "fruit"
	Banana := new(pb.OrderInfor)
	Banana.Id = "1"
	Banana.Item = "fruit"
	err = Carstream.Send(Iphone)
	if err != nil {
		fmt.Println("Send error : ", err)
	}
	//time.Sleep(time.Second * 1)
	err = Carstream.Send(Apple)
	if err != nil {
		fmt.Println("Send error : ", err)
	}
	//	time.Sleep(time.Second * 1)
	err = Carstream.Send(Sumsung)
	if err != nil {
		fmt.Println("Send error : ", err)
	}
	err = Carstream.Send(Banana)
	if err != nil {
		fmt.Println("Send error : ", err)
	}
	err_ := Carstream.CloseSend()
	if err_ != nil {
		fmt.Println("Close Send error : ", err_)
	}
	for {
		cars, err := Carstream.Recv()
		if err == io.EOF {
			fmt.Println("Delivery over !")
		}
		if err != nil {
			fmt.Println("Recv error : ", err)
		}
		fmt.Println("Cars item :", cars.Item)
		fmt.Println("order ids:")
		for _, b := range cars.Orders {
			if b != nil {
				fmt.Println(b.Id)
			}
		}
	}
}
