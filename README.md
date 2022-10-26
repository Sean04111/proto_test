# grpc双向流通信测试项目
## 项目以客户向服务端发送订单信息；服务端向客户装载订单货物的货车信息
<img width="724" alt="image" src="https://user-images.githubusercontent.com/96430610/198008231-7a47fc2c-77ee-4d3f-820b-a0b9ff6c7a5a.png">
关键源码展示：
<img width="733" alt="image" src="https://user-images.githubusercontent.com/96430610/198010970-9e5f71f4-91a7-4c94-aaba-d37c60bf33f9.png">
值得注意的是，在处理对方的发送过来的流信息的时候，需要通过一个for循环来处理，使用err==io.EOF来判断对方的流信息是否发送完毕，若为真，则已经发送完毕，若为假，则每一次for的order都是得到信息（没有错误发生）<br>
运行结果：
<img width="638" alt="image" src="https://user-images.githubusercontent.com/96430610/198008731-8932aad0-5750-4ae5-beae-1c8b84e80c56.png">

