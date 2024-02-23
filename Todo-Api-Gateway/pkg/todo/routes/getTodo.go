package routes

import (
	"context"
	"net/http"

	"github.com/akhi9550/pkg/todo/pb"
	"github.com/gin-gonic/gin"
)

type GetRequestBody struct {
	ID int64 `json:"id"`
}

func GetTodo(ctx *gin.Context, c pb.AuthServiceClient) {
	b := GetRequestBody{}
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.GetTodoByID(context.Background(), &pb.TodoIDRequest{
		ID: int64(b.ID),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)
}
