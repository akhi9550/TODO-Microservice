package todo

import (
	"fmt"

	"github.com/akhi9550/pkg/config"
	"github.com/akhi9550/pkg/todo/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.TodoSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect:", err)
	}
	return pb.NewAuthServiceClient(cc)
}
