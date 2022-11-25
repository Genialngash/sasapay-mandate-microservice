package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

func main() {
	kafkaConf := kafkaclient.Conf{
		GroupId:                  "tanda-integrations",
		ConsumerBootstrapServers: []string{"localhost:9092"},
		ProducerBootstrapServers: []string{"localhost:9092"},
	}
	//LOAD env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// mongoConf := mongodb.Conf{
	// 	Uri: os.Getenv("MONGO_DB_URL"),
	// 	Db:  "payu",
	// }

	// mongoClient, err := mongodb.New(&mongoConf)
	// if err != nil {
	// 	panic(err)
	// }
	// defer mongoClient.Disconnect(context.Background())

	// db := mongoClient.Database(mongoConf.Db)

	// rlStorage := mongodb.NewRequestLogStorage(db)
	// prlStorage := mongodb.NewPendingRequestLogStorage(db)
	eventSender := kafka.NewSender(&kafkaConf)

	

	ioReqsHander := payukafka.NewIoPaymentRequesthandler(eventSender, rlStorage, prlStorage, payuRest)

	ioTopics := []string{

		"IO-01-PAYU-TopupFix-DSTV-IoPaymentRequest-v1",
		"IO-01-PAYU-TopupFix-GOTV-IoPaymentRequest-v1",
	}
	ioPaymentResRecvr := kafka.NewReceiver(&kafkaConf, ioTopics, ioReqsHander.Handle)

	//start receiver
	ioPaymentResRecvr.ReceiveAndHandle()

	log.Fatal(fasthttp.ListenAndServe(":20460", nil))
}
