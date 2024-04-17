package statistics

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func Start() {
	ip := "127.0.0.1"
	// ip := "4.145.81.125"

	port := 4150
	addr := fmt.Sprintf("%s:%d", ip, port)
	//topic := "compression"
	// topic := "task"
	topic := "compression_statistics"

	// wg := sync.WaitGroup{}

	config := nsq.NewConfig()
	p, err := nsq.NewProducer(addr, config)
	if err != nil {
		log.Panic(err)
	}

	maxIteration := 1000
	idx := 0
	for idx < maxIteration {
		fmt.Println("idx = ", idx)

		payload := generatePayload()

		msg, err := json.Marshal(payload)
		if err != nil {
			log.Print(err)
		}
		fmt.Println(string(msg))

		fmt.Println("kirim")

		p.Publish(topic, msg)
		// go p.Publish(topic, msg)
		// defer wg.Done()
		idx += 1
		time.Sleep(time.Second * 1)
	}

	// wg.Wait()
}
