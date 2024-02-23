package routes

import (
	"context"
	"net/http"

	"github.com/akhi9550/pkg/todo/pb"
	"github.com/gin-gonic/gin"
)

type ListRequestBody struct {
	page int64 
}

func ListAllTodo(ctx *gin.Context, c pb.AuthServiceClient) {
	b := ListRequestBody{}
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.ListTodo(context.Background(), &pb.ListTodoRequest{
		Page:b.page,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)
}
