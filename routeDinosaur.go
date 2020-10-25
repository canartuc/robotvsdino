package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func CreateDinosaur(c *gin.Context) {
	row := c.Params.ByName("row")
	column := strings.ToUpper(c.Params.ByName("column"))

	rowInt, err := strconv.Atoi(row)
	if err == nil {
		rowInt = rowInt - 1
	}
	columnInt := Column[column]

	if CheckCell(rowInt, columnInt) {
		Board[rowInt][columnInt] = "d:e"
		c.JSON(http.StatusCreated, gin.H{
			"msg":             fmt.Sprintf("Dinosaur created at row %s, column %s", row, column),
			"dinosaur_row":    rowInt,
			"dinosaur_column": columnInt,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Dinosaur cannot be created. The cell may not be empty, outside the boundry or wrong parameters",
		})
	}
}
