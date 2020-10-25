package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//CreateBoard creates empty board depends on UPPERBOUND limit
func CreateBoard(c *gin.Context) {
	if CheckBoardInit() {
		InitBoard()
		c.JSON(http.StatusCreated, gin.H{
			"msg": fmt.Sprintf("%dx%d Board has been created", UPPERBOUND +1, UPPERBOUND +1),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Board has already been created!",
		})
	}
}

// ShowBoardStatus shows board status which is coming from common function
func ShowBoardStatus(c *gin.Context) {
	bs := GetBoardStatus()
	c.JSON(http.StatusOK, bs)
}
