package main

import (
	"fmt"
	"github.com/eapache/queue"
	"sync"
	"time"
)

var Cond = sync.NewCond(&sync.Mutex{})
var Queue = queue.New()

func FillQueue() {
	for {
		byts, err := GenerateRandomBytes(321)
		if err != nil {
			panic(err.Error())
		}
		Queue.Add(byts)

		fmt.Println("added:", string(byts), "queue length:", Queue.Length())
		go func() {
			if Queue.Length()%1024 == 0 {
				Cond.L.Lock()
				fmt.Println("allowing to produce...")
				allowToProduce = true
				Cond.Broadcast()
				Cond.L.Unlock()
			}
		}()
		time.Sleep(5 * time.Millisecond)
	}
}
