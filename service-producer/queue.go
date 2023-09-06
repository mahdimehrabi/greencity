package main

import (
	"github.com/eapache/queue"
	"sync"
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
		if Queue.Length()%1024 == 0 {
			Cond.L.Lock()
			allowToProduce = true
			Cond.Broadcast()
			Cond.L.Unlock()
		}
	}
}
