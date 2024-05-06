package statistics

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

const delay = time.Second * 10
const maxIteration = 10000
const topic = "compression_statistics"
const ip = "127.0.0.1"
const port = 4150

// ip := "4.145.81.125"
func Start() {

	addr := fmt.Sprintf("%s:%d", ip, port)
	// wg := sync.WaitGroup{}

	config := nsq.NewConfig()
	p, err := nsq.NewProducer(addr, config)
	if err != nil {
		log.Panic(err)
	}
	idx := 0
	for idx < maxIteration {
		fmt.Println("idx = ", idx)

		payload := generatePayload()
		msg, err := json.Marshal(payload)
		if err != nil {
			log.Print(err)
		}
		fmt.Println(string(msg))
		fmt.Println("----------------------------- SENDING -----------------------------")

		p.Publish(topic, msg)
		// go p.Publish(topic, msg)
		// defer wg.Done()
		idx += 1
		time.Sleep(delay)
	}

	// wg.Wait()
}
