gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

test:
	go test --cover -v ./service/

server:
	go run cmd/server/main.go

client:
	go run cmd/client/main.go

.PHONY: gen test client server

