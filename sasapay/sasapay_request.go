package sasapay

import (
	"fmt"
	"log"
	"time"

	kafkaclient "github.com/Genialngash/sasapay-mandate-microservice/kafkaClient"
	"github.com/robfig/cron/v3"
)

var kafkaConf = kafkaclient.Conf{
	GroupId:                  "sasapay-mandate",
	ConsumerBootstrapServers: []string{"localhost:9092"},
	ProducerBootstrapServers: []string{"localhost:9092"},
}

// kafka sender instance
var kafkaSEnder = kafkaclient.NewSender(&kafkaConf)

func MandateRequest(request []byte) {
	fmt.Println(request)

	log.Printf(string(request))
	fmt.Sprintf(string(request))
}

func checkForJobs() {
	log.Println("Starting...")
	// Define a cron runner
	c := cron.New()
	// Timing for 5 seconds, execute every 5 seconds
	c.AddFunc("@every 2s ", sendCheckQue)

	// Start
	c.Start()

	time.Sleep(time.Minute * 5)
	defer c.Stop()
}
func sendCheckQue() {
	checkMessage := "hello ! are there jobs"
	sendKafkaErr := kafkaSEnder.Send("sasapay-mandate-topic", []byte(checkMessage))
	if sendKafkaErr != nil {
		log.Println(sendKafkaErr)

	}
}
