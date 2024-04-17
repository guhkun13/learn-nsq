package task

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/nsqio/go-nsq"
)

const StateStart = "start"
const ActionDownload = "download"
const StatusSuccess = "success"

type Payload struct {
	JobId  string `json:"job_id"`
	TaskId string `json:"task_id"`
	State  string `json:"state"`
	Action string `json:"action"`
	Status string `json:"status"`
}

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

	idx := 0
	for idx < 3 {
		fmt.Println("idx = ", idx)
		// wg.Add(1)

		payload := Payload{
			JobId:  uuid.NewString(),
			TaskId: uuid.NewString(),
			State:  StateStart,
			Action: ActionDownload,
			Status: "",
		}

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
	}

	// wg.Wait()
}
