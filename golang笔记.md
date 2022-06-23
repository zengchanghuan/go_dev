#golang 笔记
##
##Go 中的方法是一种特殊类型的函数，但存在一个简单的区别：你必须在函数名称之前加入一个额外的参数。 此附加参数称为接收方。

##将 channel 用作通信机制
###1.不是通过共享内存通信，而是通过通信共享内存。
###2.需要将值从一个 goroutine 发送到另一个时，可以使用通道channel。
###虽然我们在 Go 语言中也能使用共享内存加互斥锁进行通信，但是 Go 语言提供了一种不同的并发模型，即通信顺序进程（Communicating sequential processes，CSP）1。Goroutine 和 Channel 分别对应 CSP 中的实体和传递信息的媒介，Goroutine 之间会通过 Channel 传递数据。
![](https://img.draveness.me/2020-01-28-15802171487080-channel-and-goroutines.png)
上图中的两个 Goroutine，一个会向 Channel 中发送数据，另一个会从 Channel 中接收数据，它们两者能够独立运行并不存在直接关联，但是能通过 Channel 间接完成通信。

###目前的 Channel 收发操作均遵循了先进先出的设计，具体规则如下：
* 先从 Channel 读取数据的 Goroutine 会先接收到数据；
* 先向 Channel 发送数据的 Goroutine 会得到先发送数据的权利；