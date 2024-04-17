package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	decodeConfig := nsq.NewConfig()
	topic := "compression"
	channel := "any"

	c, err := nsq.NewConsumer(topic, channel, decodeConfig)
	if err != nil {
		log.Panic("Could not create consumer")
	}

	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Println("OK READY TO ACCEPT NEW MESSAGE")
		fmt.Println(" ----------- ----------- ")
		log.Println("NSQ message received:")
		log.Println(string(message.Body))

		// aku belum bisa terima kirim pesan lagi ya
		c.ChangeMaxInFlight(0)

		// do proses nya yang amat sangat lama
		go doSomething(c)

		log.Println("Pesan Diterima Dengan Baik. TAPI BELUM SIAP TERIMA PESAN LAGI YA")
		log.Println("ADA YANG LAGI DIPROSES")

		// pesan yang kamu kirim udah aku terima
		return nil
	}))

	ip := "127.0.0.1"
	//ip := "4.145.81.125"
	port := 4150
	addr := fmt.Sprintf("%s:%d", ip, port)

	err = c.ConnectToNSQD(addr)
	if err != nil {
		log.Panic("Could not connect")
	}
	log.Printf("Awaiting messages from NSQ topic \"%s\"... \n", topic)
	wg.Wait()
}

func doSomething(c *nsq.Consumer) {
	fmt.Println("logic here")

	randomTime := rand.Intn(10)
	fmt.Println("do something for ", randomTime, " detik")

	time.Sleep(time.Second * time.Duration(randomTime))
	fmt.Println("XXXXXXXXXXXXX SELESAI XXXXXXXXXXXXX")

	c.ChangeMaxInFlight(1)
}
