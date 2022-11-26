package kafkaclient


// Conf holds all Kafka consumer and producer configurations
type Conf struct {
	GroupId                  string   `yaml:"kafka-group-id"`
	ConsumerBootstrapServers []string `yaml:"kafka-consumer-bootstrap-servers"`
	ProducerBootstrapServers []string `yaml:"kafka-producer-bootstrap-servers"`
}
