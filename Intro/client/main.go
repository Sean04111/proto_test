package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	pb "test/Intro/client/Test"
	"time"
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
	SendAndCheck(Carstream)
	go ListenAndWrite(Carstream)
	time.Sleep(5 * time.Second)
}
func ListenAndWrite(stream pb.OrderSend_OrderDelyClient) {
	for {
		cars, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		fmt.Println("Cars item :", cars.Item)
		fmt.Println("Cars id : ", cars.Carid)
		fmt.Println("order ids:")
		for _, b := range cars.Orders {
			if b != nil {
				fmt.Println(b.Id)
			}
		}
	}
}
func SendAndCheck(stream pb.OrderSend_OrderDelyClient) {
	Iphone := new(pb.OrderInfor)
	Iphone.Id = "0"
	Iphone.Item = "phone"
	Samsung := new(pb.OrderInfor)
	Samsung.Id = "3"
	Samsung.Item = "phone"
	Apple := new(pb.OrderInfor)
	Apple.Id = "2"
	Apple.Item = "fruit"
	Banana := new(pb.OrderInfor)

	Banana.Id = "1"
	Banana.Item = "fruit"
	err := stream.Send(Iphone)
	if err != nil {
		fmt.Println("Send error : ", err)
	} else {
		fmt.Println("Send Iphone done")
	}
	//time.Sleep(time.Second * 1)
	err = stream.Send(Apple)
	if err != nil {
		fmt.Println("Send error : ", err)
	} else {
		fmt.Println("Send Apple done")
	}
	//	time.Sleep(time.Second * 1)
	err = stream.Send(Samsung)
	if err != nil {
		fmt.Println("Send error : ", err)
	} else {
		fmt.Println("Send Samsung done")
	}
	err = stream.Send(Banana)
	if err != nil {
		fmt.Println("Send error : ", err)
	} else {
		fmt.Println("Send Banana done")
	}
	err_ := stream.CloseSend()
	if err_ != nil {
		fmt.Println("Close Send error : ", err_)
	}
}
