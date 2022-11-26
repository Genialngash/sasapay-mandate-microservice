package main

import (
	"log"

	kafkaclient "github.com/Genialngash/sasapay-mandate-microservice/kafkaClient"
	"github.com/Genialngash/sasapay-mandate-microservice/sasapay"
	"github.com/valyala/fasthttp"
)

func main() {
	kafkaConf := kafkaclient.Conf{
		GroupId:                  "sasapay-mandate",
		ConsumerBootstrapServers: []string{"localhost:9092"},
		ProducerBootstrapServers: []string{"localhost:9092"},
	}

	ioTopics := []string{

		"sasapay-mandate-topic",
	}
	ioPaymentResRecvr := kafkaclient.NewReceiver(&kafkaConf, ioTopics, sasapay.MandateRequest)

	//start receiver
	ioPaymentResRecvr.ReceiveAndHandle()

	log.Fatal(fasthttp.ListenAndServe(":20460", nil))
}
