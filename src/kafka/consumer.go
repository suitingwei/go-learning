package main

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"time"
)

func main() {

	//配置 broker 集群,以及特定的 config
	consumer, err := sarama.NewConsumer([]string{"localhost:9092", "localhost:9093"}, nil)

	if err != nil {
		log.Fatalln("fail to create the consumer")
	}

	//获取 broker集群的 topics
	topics, err := consumer.Topics()

	if err != nil {
		log.Fatalln("fail to retrieve the topics from the broker cluster")
	}

	log.Printf("Current available topics:%v", topics)

	testPartitions, err := consumer.Partitions("test")

	if err != nil {
		log.Fatalln("fail to retrieve the topics from the broker cluster")
	}

	log.Printf("Current available partitions in topic:[test] => %v", testPartitions)

	partitionConsumer, err := consumer.ConsumePartition("test", 0, 100)

	if err != nil {
		log.Fatalln("fail to create the partition consumer on topic:[test],partition[0]")
	}

	messages := partitionConsumer.Messages()

	for {
		singleMessage := <-messages

		log.Printf("Successfully retrieve a message from kafka\n"+
			"\ttopic:%s,partition:%d\n\tkey:%v,value:%q,offset:%d\n", singleMessage.Topic, singleMessage.Partition, singleMessage.Key, singleMessage.Value, singleMessage.Offset)

		time.Sleep(time.Millisecond * 100)
	}

	os.Exit(0)
}
