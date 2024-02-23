package routes

import (
	"context"
	"net/http"

	"github.com/akhi9550/pkg/todo/pb"
	"github.com/gin-gonic/gin"
)

type CreateRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateTodo(ctx *gin.Context, c pb.AuthServiceClient) {
	b := CreateRequestBody{}
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.CreateTodo(context.Background(), &pb.AddTodoRequest{
		Title:       b.Title,
		Description: b.Description,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}
