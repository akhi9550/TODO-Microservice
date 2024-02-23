package todo

import (
	"github.com/akhi9550/pkg/config"
	"github.com/akhi9550/pkg/todo/routes"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine,c *config.Config) *ServiceClient{
	svc:=&ServiceClient{
		Client: InitServiceClient(c),
	}
	routes:=r.Group("/todo")
	routes.POST("/create",svc.createTodo)
	routes.DELETE("/delete",svc.DeleteTodo)
	routes.GET("/",svc.GetTodo)
	routes.GET("/all",svc.ListAll)
	routes.PUT("/update",svc.updateTodo)
	return svc
}
func (svc *ServiceClient) createTodo(ctx *gin.Context){
	routes.CreateTodo(ctx,svc.Client)
}
func (svc *ServiceClient) DeleteTodo(ctx *gin.Context){
	routes.DeleteTodo(ctx,svc.Client)
}
func (svc *ServiceClient) GetTodo(ctx *gin.Context){
	routes.GetTodo(ctx,svc.Client)
}
func (svc *ServiceClient) ListAll(ctx *gin.Context){
	routes.ListAllTodo(ctx,svc.Client)
}
func (svc *ServiceClient) updateTodo(ctx *gin.Context){
	routes.UpdateTodo(ctx,svc.Client)
}