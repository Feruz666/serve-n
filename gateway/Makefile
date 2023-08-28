build:
	@go build -o bin/gateway

run: build
	@./bin/gateway

topics: 
	@docker exec broker kafka-topics --bootstrap-server broker:9092 --create --topic account-topic
	@docker exec broker kafka-topics --bootstrap-server broker:9092 --create --topic survey-topic
	@docker exec broker kafka-topics --bootstrap-server broker:9092 --create --topic map-topic

topic-list:
	@docker exec broker kafka-topics --list --bootstrap-server localhost:9092
	
proto:
	@protoc --go_out=paths=source_relative:./protos/account --go-grpc_out=paths=source_relative:./protos/account ./protos/account.proto

.PHONY: account survey maps


# docker exec bff4a89ffd63 kafka-topics --list --bootstrap-server localhost:9092
# docker exec brokerContainerName kafka-topics --bootstrap-server brokerName:port --create --topic topicName