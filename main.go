package main

import (
	"cotton-gin-test/internal/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	router.Init(engine)
	log.Printf("server listen at :%d \n", 10001)
	engine.Run(":10001")
}
