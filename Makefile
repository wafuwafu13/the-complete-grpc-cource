gen:
	protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false proto/*.proto

clean:
	rm proto/*.go

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

test:
	go test -cover -race ./...
