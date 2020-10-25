package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func CreateRobot(c *gin.Context) {
	row := c.Params.ByName("row")
	column := strings.ToUpper(c.Params.ByName("column"))
	face := strings.ToUpper(c.Params.ByName("face"))

	rowInt, err := strconv.Atoi(row)
	if err == nil {
		rowInt = rowInt - 1
	}
	columnInt := Column[column]

	if CheckCell(rowInt, columnInt) {
		faceDir := strings.ToLower(Face[face])
		Board[rowInt][columnInt] = fmt.Sprintf("r:%s", faceDir)
		c.JSON(http.StatusCreated, gin.H{
			"msg":          fmt.Sprintf("Robot created at row %s, column %s and face %s", row, column, face),
			"robot_row":    rowInt,
			"robot_column": columnInt,
			"robot_face":   faceDir,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Robot cannot be created. The cell may not be empty, outside the boundry or wrong parameters",
		})
	}
}

func MoveRobot(c *gin.Context) {
	row := c.Params.ByName("row")
	column := strings.ToUpper(c.Params.ByName("column"))
	face := strings.ToUpper(c.Params.ByName("face"))

	rowInt, err := strconv.Atoi(row)
	if err == nil {
		rowInt = rowInt - 1
	}
	columnInt := Column[column]
	faceDir := strings.ToLower(Face[face])

	moveRow, moveCol, moveFace, err := MoveRobotState(rowInt, columnInt, faceDir)
	moveColLetter, errColLetter := ReturnKeyFromValueMap(Column, moveCol)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":          err.Error(),
			"robot_row":    moveRow,
			"robot_column": moveCol,
			"robot_face":   moveFace,
		})
	} else if errColLetter != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":          errColLetter.Error(),
			"robot_row":    moveRow,
			"robot_column": moveCol,
			"robot_face":   moveFace,
		})
	} else {
		// check if the cell is empty because robots can only move to empty cells
		if CheckCell(moveRow, moveCol) {
			Board[rowInt][columnInt] = "e:e"
			Board[moveRow][moveCol] = fmt.Sprintf("r:%s", moveFace)

			c.JSON(http.StatusCreated, gin.H{
				"msg": fmt.Sprintf("Robot moved to row %d, column %s and face %s",
					moveRow+1, moveColLetter, face),
				"robot_row":    rowInt,
				"robot_column": moveCol,
				"robot_face":   faceDir,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":          fmt.Sprintf("Cell is not empty, robot cannot move"),
				"robot_row":    moveRow,
				"robot_column": moveCol,
				"robot_face":   faceDir,
			})
		}
	}
}

func AttackRobot(c *gin.Context) {
	row := c.Params.ByName("row")
	column := strings.ToUpper(c.Params.ByName("column"))
	face := strings.ToUpper(c.Params.ByName("face"))

	rowInt, err := strconv.Atoi(row)
	if err == nil {
		rowInt = rowInt - 1
	}
	columnInt := Column[column]
	faceDir := strings.ToLower(Face[face])

	moveRow, moveCol, moveFace, err := MoveRobotState(rowInt, columnInt, faceDir)
	moveColLetter, errColLetter := ReturnKeyFromValueMap(Column, moveCol)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":          err.Error(),
			"robot_row":    moveRow,
			"robot_column": moveCol,
			"robot_face":   moveFace,
		})
	} else if errColLetter != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":          errColLetter.Error(),
			"robot_row":    moveRow,
			"robot_column": moveCol,
			"robot_face":   moveFace,
		})
	} else if CheckType(moveRow, moveCol) != "d" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":          "Robot can attack only on dinosaur, no dinosaur found",
			"robot_row":    moveRow,
			"robot_column": moveCol,
			"robot_face":   moveFace,
		})
	} else {
		Board[rowInt][columnInt] = "e:e"
		Board[moveRow][moveCol] = fmt.Sprintf("r:%s", moveFace)

		c.JSON(http.StatusCreated, gin.H{
			"msg": fmt.Sprintf("Robot attacked to row %d, column %s and face %s",
				moveRow+1, moveColLetter, face),
			"robot_row":    rowInt,
			"robot_column": moveCol,
			"robot_face":   faceDir,
		})
	}
}
