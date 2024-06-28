package main

import (
	"log"

	"github.com/akhi9550/cmd/docs"
	server "github.com/akhi9550/pkg/api"
	"github.com/akhi9550/pkg/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func main() {
	docs.SwaggerInfo.Title = "TODOAPP_Microservice_CleanArchitecture"
	docs.SwaggerInfo.Description = "ToDo is a list of activities that is to be done by a particular individual"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "todoapp.zhooze.shop"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"https"}
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Routes(r, &c)
	r.Run(c.Port)
}
