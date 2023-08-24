build:
	@go build -o bin/gateway

run: build
	@./bin/gateway


account:
	@protoc --go_out=paths=source_relative:./protos/account --go-grpc_out=paths=source_relative:./protos/account ./protos/account.proto

.PHONY: account survey maps


# docker exec bff4a89ffd63 kafka-topics --list --bootstrap-server localhost:9092
# docker exec brokerContainerName kafka-topics --bootstrap-server brokerName:port --create --topic topicName