package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:19092"
	kafkaTopic = "user_topic_001"
)

// For producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// For consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		GroupID:     groupID,
		Topic:       topic,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		StartOffset: kafka.FirstOffset,
	})

}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

//Mua ban chung khoan

func NewStock(msg, typeMessage string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMessage
	return &s
}

func ActionStock(c *gin.Context) {
	s := NewStock(c.Query("msg"), c.Query("type"))

	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)

	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "Action Successfully",
	})
}

func RegisterConsumerATC(id int) {
	kafkaGroupId := "consumer-group-"
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer(%d) Hong Phien ATC::\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer(%d) Hong Phien ATC:%v\n", id, err)
		}
		fmt.Printf("Consumer(%d) hong topic:%v, partition:%v, offset:%v, time:%d %s = %s\n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("action/stock", ActionStock)

	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)

	r.Run(":8999")
}
