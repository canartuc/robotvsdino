package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
)

func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func CheckLowerUpperBound(row int, col int) bool {
	if (row < UPPERBOUND + 1 && row >= LOWERBOUND) || (col < UPPERBOUND + 1 && col >= LOWERBOUND) {
		return true
	}
	return false
}

func CheckCell(row int, column int) bool {
	if CheckLowerUpperBound(row, column) {
		s := strings.Split(Board[row][column], ":")
		if s[0] == "e" {
			return true
		}
	}
	return false
}

func CheckType(row int, column int) string {
	if CheckLowerUpperBound(row, column) {
		s := strings.Split(Board[row][column], ":")
		return s[0]
	}
	return ""
}

func InitBoard() {
	for i := 0; i < 50; i++ {
		Board[i] = make([]string, 50)
		for j := 0; j < 50; j++ {
			Board[i][j] = "e:e"
		}
	}
}

func CheckBoardInit() bool {
	res := false
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			if Board[i][j] == "e:e" {
				res = true
			} else {
				return false
			}
		}
	}
	return res
}

func ReturnKeyFromValueMap(m map[string]int, val int) (key string, err error) {
	for k, v := range m {
		if v == val {
			key = k
			err = nil
			return
		}
	}
	err = errors.New("could not found key from value")
	return
}

// MoveRobotState is the state machine of robot movement
func MoveRobotState(row int, column int, face string) (newRow int, newColumn int, newFace string, err error) {
	if CheckType(row, column) == "r" {
		if strings.ToLower(face) == "east" {
			newRow = row
			newColumn = column + 1
			newFace = "east"
			err = nil
		} else if strings.ToLower(face) == "west" {
			newRow = row
			newColumn = column - 1
			newFace = "west"
			err = nil
		} else if strings.ToLower(face) == "north" {
			newRow = row - 1
			newColumn = column
			newFace = "north"
			err = nil
		} else if strings.ToLower(face) == "south" {
			newRow = row + 1
			newColumn = column
			newFace = "south"
			err = nil
		} else {
			newRow = row
			newColumn = column
			newFace = face
			err = errors.New("unrecognized direction for movement")
		}
	} else {
		newRow = row
		newColumn = column
		newFace = face
		err = errors.New("only robots can move. There is no robot in the cell")
	}
	if CheckLowerUpperBound(newRow, newColumn) {
		newRow = row
		newColumn = column
		newFace = face
		err = errors.New("out of index boundries")
	}
	if !CheckCell(newRow, newColumn) {
		newRow = row
		newColumn = column
		newFace = face
		err = errors.New("robot can move to only empty cell")
	}
	return
}
