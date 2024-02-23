package routes

import (
	"context"
	"net/http"

	"github.com/akhi9550/pkg/todo/pb"
	"github.com/gin-gonic/gin"
)

type DeleteRequestBody struct {
	ID int64 `json:"id"`
}

func DeleteTodo(ctx *gin.Context, c pb.AuthServiceClient) {
	b := DeleteRequestBody{}
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.DeleteTodo(context.Background(), &pb.TodoIDRequest{
		ID: int64(b.ID),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)
}
