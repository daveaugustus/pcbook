package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/davetweetlive/pcbook/pb"
	"github.com/davetweetlive/pcbook/service"
	"google.golang.org/grpc"
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
	lisener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start the server: ", err)
	}

	err = grpcServer.Serve(lisener)
	if err != nil {
		log.Fatal("cannot start the server: ", err)
	}
}
