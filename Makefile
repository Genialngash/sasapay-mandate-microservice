.PHONY:

compose:
	echo "Starting docker compose"
	cd docker-set-up && \
	docker-compose --env-file .env up -d

#========KAFKA==========
create_topics:
	docker exec -it kafka "kafka-topics.sh --bootstrap-server kafka:9092 --create --topic sasapay-mandate-topic1 --partitions 3 --replication-factor 1"

#=========GOLANG=========
tidy:
	go mod tidy
	go mod vendor

	