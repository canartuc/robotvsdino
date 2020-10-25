package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RouteIndex is main dummy index
func RouteIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello World",
	})
}
