package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getConfigTemplateProxy(c *gin.Context) {

	c.String(http.StatusOK, "hello")
}
