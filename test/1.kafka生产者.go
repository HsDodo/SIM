package main

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/zeromicro/go-zero/core/logx"
)

type Msg struct {
	Msg        string `json:"msg"`
	SenderName string `json:"senderName"`
	RecverName string `json:"recverName"`
}

func SendMsgToTopic(topic string, msg Msg) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal

	kafkaProducer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		logx.Error(err)
		return
	}
	defer kafkaProducer.Close()
	msgBytes, _ := json.Marshal(msg)
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msgBytes),
	}
	_, _, err = kafkaProducer.SendMessage(kafkaMsg)
	if err != nil {
		logx.Errorf("kafka 消息发送失败: %s", err.Error())
		return
	}
}

func main() {
	SendMsgToTopic("test_topic", Msg{
		Msg:        "森哥你好",
		SenderName: "森",
		RecverName: "小小森",
	})

	SendMsgToTopic("test_topic", Msg{
		Msg:        "啦啦啦啦啦",
		SenderName: "森",
		RecverName: "小小森",
	})

}
