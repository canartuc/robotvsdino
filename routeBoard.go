package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBoard(c *gin.Context) {
	if CheckBoardInit() {
		InitBoard()
		c.JSON(http.StatusCreated, gin.H{
			"msg": "50x50 Board has been created",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Board has already been created!",
		})
	}
}

func ShowBoardStatus(c *gin.Context) {
	bs := GetBoardStatus()
	c.JSON(http.StatusOK, bs)
}
