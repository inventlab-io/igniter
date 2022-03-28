package http

import (
	"github.com/gin-gonic/gin"
)

func InitGin(routeBuilder func(e *gin.Engine)) {
	r := gin.Default()
	routeBuilder(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
