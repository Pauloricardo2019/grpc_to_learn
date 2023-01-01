package main

import (
	"google.golang.org/grpc"
	pb "grpc-demo/proto"
	"log"
	"net"
)

type helloServer struct {
	pb.GreetServiceServer
}

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the server %v", err)
	}

 
	grpcServer := grpc.NewServer()
log.Println("start grpc server...")
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("falied to start: %v", err)
	}

}
