package main

import (
	"log"

	"github.com/akhi9550/pkg/config"
	"github.com/akhi9550/pkg/todo"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	r := gin.Default()
	todo.Routes(r, &c)
	r.Run(c.Port)
}
