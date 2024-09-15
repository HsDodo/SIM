package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/zeromicro/go-zero/core/logx"
)

type MsgData struct {
	Msg        string `json:"msg"`
	SenderName string `json:"senderName"`
	RecverName string `json:"recverName"`
}

func RecvMsgFromTopic(topic string) {
	// 从kafka中消费消息, 开一个协程就行
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // 从最老的消息开始消费
	kafkaConsumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		logx.Errorf("消费者创建失败:%s", err)
		return
	}
	defer kafkaConsumer.Close()
	partitionList, err := kafkaConsumer.Partitions(topic)
	if err != nil {
		logx.Errorf("获取分区列表失败：%s", err)
		return
	}
	fmt.Println("partition:", partitionList)

	consumers := make([]sarama.PartitionConsumer, len(partitionList))

	for i, partition := range partitionList {
		consumers[i], err = kafkaConsumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			logx.Errorf("创建分区消费者失败：%s", err)
			return
		}
		go func(consumer sarama.PartitionConsumer) {
			for {
				select {
				case m := <-consumer.Messages():
					var msgModel MsgData
					err := json.Unmarshal(m.Value, &msgModel)
					if err != nil {
						logx.Errorf("消息解析失败: %s", err.Error())
						continue
					}
					fmt.Println(msgModel)
				}
			}
		}(consumers[i])
	}
	select {}
}

func main() {
	go RecvMsgFromTopic("test_topic")
	select {}
}
