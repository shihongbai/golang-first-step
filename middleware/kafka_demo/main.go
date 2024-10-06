package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// kafka client demo
func main() {

	// 1 生产者配置
	config := sarama.NewConfig()
	// 设置ack的确认方式
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 设置发送消息的分区，随机发送
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 成功交付的信息
	config.Producer.Return.Successes = true

	// 2 连接Kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}

	defer client.Close()

	// 3 封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "shopping"
	msg.Value = sarama.StringEncoder("hello world")

	// 4 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message err:", err)
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
