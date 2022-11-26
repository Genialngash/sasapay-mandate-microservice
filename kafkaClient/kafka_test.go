package kafkaclient

import (
	"fmt"
	"log"
	"testing"

)

var kafkaConf = Conf{
	GroupId:                  "sasapay-mandate",
	ConsumerBootstrapServers: []string{"localhost:9092"},
	ProducerBootstrapServers: []string{"localhost:9092"},
}

// kafka sender instance
var kafkaSEnder = NewSender(&kafkaConf)

func TestKafkaSender(t *testing.T) {

	checkMessage := "hello ! micoro service, do the following jobs"
	sendKafkaErr := kafkaSEnder.Send("sasapay-mandate-topic", []byte(checkMessage))
	if sendKafkaErr != nil {
		log.Println(sendKafkaErr)

	}
}

func TestKafkaReceiver(t *testing.T) {

	ioTopics := []string{

		"sasapay-mandate-topic",
	}
	ioPaymentResRecvr := NewReceiver(&kafkaConf, ioTopics,printReceivedMessages)

	
	ioPaymentResRecvr.ReceiveAndHandle()

	
}

func printReceivedMessages( message []byte){

	fmt.Sprintf(string(message))


}
