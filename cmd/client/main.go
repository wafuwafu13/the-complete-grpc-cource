package main

import (
	"context"
	"flag"
	pb "github.com/wafuwafu13/the-complete-grpc-course/pb/proto"
	"github.com/wafuwafu13/the-complete-grpc-course/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func main() {
	serverAdress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAdress)

	conn, err := grpc.Dial(*serverAdress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}

	log.Printf("crated laptop with id: %s", res.Id)
}
