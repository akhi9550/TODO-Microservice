package client

import (
	"fmt"

	"github.com/akhi9550/pkg/config"
	"github.com/akhi9550/pkg/pb"
	"google.golang.org/grpc"
)

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.TodoSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect:", err)
	}
	return pb.NewAuthServiceClient(cc)
}
