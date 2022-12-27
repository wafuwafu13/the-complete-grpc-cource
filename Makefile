gen:
	protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false proto/*.proto

clean:
	rm proto/*.go

server1:
	go run cmd/server/main.go -port 50051

server2:
	go run cmd/server/main.go -port 50052

server1-tls:
	go run cmd/server/main.go -port 50051 -tls

server2-tls:
	go run cmd/server/main.go -port 50052 -tls

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

client-tls:
	go run cmd/client/main.go -address 0.0.0.0:8080 -tls

test:
	go test -cover -race ./...

cert:
	cd cert; ./gen.sh; cd ..

.PHONY: clean gen server client test cert
