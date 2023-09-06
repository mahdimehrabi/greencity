package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"sync"
	"time"
)

var allowToProduce = false
var kafka = NewKafkaProducer([]string{"localhost:29092", "localhost:29093"})
var queueMutex = sync.Mutex{}

func Producer(topic string, partition int, wg *sync.WaitGroup) {
	defer func() {
		recover()
		wg.Done()
	}()
	for {
		var data []byte
		func() {
			defer queueMutex.Unlock()
			queueMutex.Lock()
			data = Queue.Peek().([]byte)
			Queue.Remove()
		}()

		//produce to kafka
		kafka.Input() <- &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(data),
		}
		fmt.Println("sent ", string(data), "to kafka asynchronously")
		//remove if produced successfully to kafka
	}
}

func ProduceKafka() {
	for {
		Cond.L.Lock()
		if !allowToProduce {
			Cond.Wait()
		}
		Cond.L.Unlock()
		wg := new(sync.WaitGroup)
		wg.Add(4)
		go Producer("green-city", 0, wg)
		go Producer("green-city", 1, wg)
		go Producer("green-city", 2, wg)
		go Producer("green-city", 3, wg)
		wg.Wait()

		//stopping produce
		Cond.L.Lock()
		allowToProduce = false
		Cond.L.Unlock()
	}
}

func NewKafkaProducer(brokersList []string) sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 10
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages

	producer, err := sarama.NewAsyncProducer(brokersList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	go func() {
		for err := range producer.Errors() {
			//logging out to STDOUT
			log.Println("Warning: failed to produce:", err)

			//adding message to queue again on error
			byts, er := err.Msg.Value.Encode()
			if er != nil {
				fmt.Println("Critical: failed to encode message:", err.Msg, "we lost it,err:", er.Error())
				continue
			}
			Queue.Add(byts)
		}
	}()
	go func() {
		for message := range producer.Successes() {
			byts, err := message.Value.Encode()
			if err != nil {
				fmt.Println("Warning: failed to encode message:", err.Error())
				continue
			}
			//adding message to redis on successful storing on kafka
			log.Println("successfully produced:", string(byts))
		}
	}()
	return producer
}
