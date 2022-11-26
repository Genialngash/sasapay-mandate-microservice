package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Conf encapsulates Mongo DB login credentials and the target database name
type Conf struct {
	Uri string `yaml:"mongo-uri"`
	Db  string `yaml:"mongo-db"`
}

func New(conf *Conf) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Uri))
	defer cancel()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return client, nil

}
