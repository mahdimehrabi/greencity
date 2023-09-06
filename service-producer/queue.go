package main

import "github.com/eapache/queue"

var Queue = queue.New()

func FillQueue() {
	for {
		byts, err := GenerateRandomBytes(15)
		if err != nil {
			panic(err.Error())
		}
		Queue.Add(byts)
	}
}
