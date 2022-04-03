package http

import (
	"github.com/gin-gonic/gin"
)

func InitGin(routeBuilder func(e *gin.Engine)) {
	e := gin.Default()
	routeBuilder(e)
	e.Run() // listen and serve on 0.0.0.0:8080
}
