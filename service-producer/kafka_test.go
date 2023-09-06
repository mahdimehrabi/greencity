package main

import (
	"github.com/IBM/sarama/mocks"
	"sync"
	"testing"
)

func TestProducer(t *testing.T) {
	producer := mocks.NewAsyncProducer(t, nil)
	producer.ExpectInputAndSucceed()
	wg := new(sync.WaitGroup)
	Queue.Add([]byte{'m', 'a', 'h', 'd', 'i'})

	// Use a test function to call Producer and capture the panic
	testFunc := func() {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("Expected a panic, but no panic occurred")
			}
		}()
		Producer(producer, "green-city", 0, wg)
	}

	// Execute the test function and expect a panic
	testFunc()

	if Queue.Length() > 0 {
		t.Fatalf("queue length must decrease by 1 after producer")
	}
}
