package main

import (
	"fmt"
	"log"
	"net"

	"github.com/akhi9550/todo-service/pkg/api/services"
	"github.com/akhi9550/todo-service/pkg/config"
	"github.com/akhi9550/todo-service/pkg/db"
	"github.com/akhi9550/todo-service/pkg/pb"
	"github.com/akhi9550/todo-service/pkg/repository"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	h := db.Init(c.DBUrl)
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	fmt.Println("Todo Svc on", c.Port)

	repo := repository.NewDBServer(h)
	s := &services.Server{
		Repository: repo,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
