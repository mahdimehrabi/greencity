package main

func main() {
	go FillQueue()
	ProduceKafka()
}
