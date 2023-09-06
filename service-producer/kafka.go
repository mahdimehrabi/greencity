package main

var allowToProduce = false

func Producer(topic string, partition int) {
	for Queue.Length() > 0 {
		Cond.L.Lock()
		if !allowToProduce {
			Cond.Wait()
		}
		data := Queue.Peek()
		//produce to kafka
		//remove if produced successfully to kafka
		Queue.Remove()
		Cond.L.Unlock()
	}
}

func ProduceKafka() {
	defer func() {
		if r := recover(); r != nil {
			//panic happened ( queue became empty)
			//stopping produce
			Cond.L.Lock()
			allowToProduce = false
			Cond.Broadcast()
			Cond.L.Unlock()
		}
	}()

	go Producer("green-city", 0)
	go Producer("green-city", 1)
	go Producer("green-city", 2)
	go Producer("green-city", 3)
}
