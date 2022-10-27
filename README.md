# gRPC双向流通信测试项目
## 订单通信-项目以客户向服务端发送订单信息；服务端向客户装载订单货物的货车信息
<img width="724" alt="image" src="https://user-images.githubusercontent.com/96430610/198008231-7a47fc2c-77ee-4d3f-820b-a0b9ff6c7a5a.png">
服务端goroutine函数：<br>

```

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

```

服务端服务函数重写：

```

func (this *server) OrderDely(stream pb.OrderSend_OrderDelyServer) error {
	conn := make(chan *pb.OrderInfor, 5)
	go Listen(stream, conn)
	return this.Return(stream, conn)
}

```

客户端服务函数见源码<br>
值得注意的是，在处理对方的发送过来的流信息的时候，需要通过一个for循环来处理，使用err==io.EOF来判断对方的流信息是否发送完毕，若为真，则已经发送完毕，若为假，则每一次for的order都是得到信息（没有错误发生）,同时，为了实现服务端和客户端的收发互不干扰，使用goroutine来调度，但是在客户端中主线程需要sleep来等待子线程，而在服务端可以用channel阻塞来解决这个问题<br>
运行结果：<br>
<img width="838" alt="image" src="https://user-images.githubusercontent.com/96430610/198266513-ce5ae83c-0362-4076-9b11-4bf5d4164191.png">

<br>注意，此处的Carid为每一个货车的id，在生成的时候，使用了<br>this.Trucks[order.Item].Carid = strconv.Itoa(time.Now().Second())<br>,这是为了模拟在实际中货车的id不是一个服务器端能够决定的内容，所以这里输出的都是相同的秒数，因为服务端在处理流的时候时间间隔很短。<br>

