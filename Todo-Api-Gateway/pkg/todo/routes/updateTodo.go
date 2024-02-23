package routes

import (
	"context"
	"net/http"

	"github.com/akhi9550/pkg/todo/pb"
	"github.com/gin-gonic/gin"
)

type UpdateRequestBody struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func UpdateTodo(ctx *gin.Context, c pb.AuthServiceClient) {
	b := UpdateRequestBody{}
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.UpdateTodo(context.Background(), &pb.UpdateTodoRequest{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)
}
