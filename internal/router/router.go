package router

import (
	"cotton-gin-test/internal/controller"
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	engine.GET("/test/sum", controller.Sum)
}
