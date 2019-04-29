package main

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.NoResponse

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)

	if err != nil {
		log.Fatalln(err)
	}
	//close the producer
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	startTime := time.Now()
	for {
		currentTime := time.Now()

		//写 kafka 五分钟吧
		if currentTime.Sub(startTime).Minutes() >= 5 {
			os.Exit(0)
		}

		messages := produceMessages()

		//发送消息
		sendErr := producer.SendMessages(messages)

		if sendErr != nil {
			log.Printf("FAILED to send message: %s\n", sendErr)
		} else {
			//log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
			log.Printf("Messages sent to partition somewhere\n")
		}

		time.Sleep(time.Second * 1)
	}
}

func produceMessages() []*sarama.ProducerMessage {
	var messages []*sarama.ProducerMessage

	for i := 0; i < 10; i++ {
		//要发送的消息
		msg := &sarama.ProducerMessage{
			Topic: "test", Value: sarama.StringEncoder("test message " + strconv.Itoa(i)),
		}
		messages = append(messages, msg)
	}

	return messages
}
