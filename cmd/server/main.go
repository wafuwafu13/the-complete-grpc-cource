package main

import (
	"flag"
	"fmt"
	pb "github.com/wafuwafu13/the-complete-grpc-course/pb/proto"
	"github.com/wafuwafu13/the-complete-grpc-course/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")

	laptopServer := service.NewLaptopServer(laptopStore, imageStore)
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("can not start server:", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("can not start server:", err)
	}
}
