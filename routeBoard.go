package main

import (
	"fmt"
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

	fmt.Println(Board)
	fmt.Println(len(Board))
}
