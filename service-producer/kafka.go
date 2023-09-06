package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"sync"
	"time"
)

var allowToProduce = false
var kafka = NewKafkaProducer([]string{"localhost:29092", "localhost:39092"})
var queueMutex = sync.Mutex{}

func Producer(topic string, partition int32, wg *sync.WaitGroup) {
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
			Topic:     topic,
			Value:     sarama.ByteEncoder(data),
			Partition: partition,
		}
		fmt.Println("sent ", string(data), "to kafka asynchronously")
		//remove if produced successfully to kafka
	}
}

func ProduceKafka() {
	go handleSuccesses()
	go handleErrors()
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
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewAsyncProducer(brokersList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	return producer
}

func handleSuccesses() {
	for message := range kafka.Successes() {
		byts, err := message.Value.Encode()
		if err != nil {
			fmt.Println("Warning: failed to encode message:", err.Error())
			continue
		}
		//adding message to redis on successful storing on kafka
		log.Println("successfully produced:", string(byts))
	}
}

func handleErrors() {
	for err := range kafka.Errors() {
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
}
