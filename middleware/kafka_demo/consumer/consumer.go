package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

func main() {
	// 创建新的消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Println("Failed to start consumer:", err)
		return
	}

	partitions, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Println("Failed to get list of partitions:", err)
		return
	}

	// 遍历分区
	wg := sync.WaitGroup{}
	for partition := range partitions {
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s", partition, err)
			return
		}

		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf(
					"Partition: %d, Offset: %d, Key: %s, Value: %s, Timestamp: %s\n",
					msg.Partition,
					msg.Offset,
					string(msg.Key),
					string(msg.Value),
					msg.Timestamp,
				)
			}
		}(pc)
	}

	wg.Wait()
}
