package main

import (
	"errors"
	"strings"
)

// CheckLowerUpperBound checks the lower and upper bound of the board as defined in consdef.go file
func CheckLowerUpperBound(row int, col int) bool {
	if (row < UPPERBOUND+1 && row >= LOWERBOUND) || (col < UPPERBOUND+1 && col >= LOWERBOUND) {
		return true
	}
	return false
}

// CheckCell returns true if the cell is empty (e) or false if not
func CheckCell(row int, column int) bool {
	if CheckLowerUpperBound(row, column) {
		s := strings.Split(Board[row][column], ":")
		if s[0] == "e" {
			return true
		}
	}
	return false
}

// CheckType returns the type of the cell as e: empty, r: robot and d: dinosaur
func CheckType(row int, column int) string {
	if CheckLowerUpperBound(row, column) {
		s := strings.Split(Board[row][column], ":")
		return s[0]
	}
	return ""
}

// InitBoard initialize the board with full of empty values
// Robots have face so first is type (e: empty, r: robot, d: dinosaur) and the second is face direction
// (e: empty, west, north, east, south)
func InitBoard() {
	for i := 0; i < UPPERBOUND + 1; i++ {
		Board[i] = make([]string, UPPERBOUND + 1)
		for j := 0; j < UPPERBOUND + 1; j++ {
			Board[i][j] = "e:e"
		}
	}
}

// CheckBoardInit checks the initialization of the board because we should not initialize it again after the
// initial initialization. Initialization makes every cells empty so we will lose the movements if we don't check
// the initialization
func CheckBoardInit() bool {
	res := false
	for i := 0; i < UPPERBOUND + 1; i++ {
		for j := 0; j < UPPERBOUND + 1; j++ {
			if Board[i][j] == "e:e" {
				res = true
			} else {
				return false
			}
		}
	}
	return res
}

// ReturnKeyFromValueMap checks value of the map and return its key
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
func MoveRobotState(row int, column int, face string, attack bool) (newRow int, newColumn int, newFace string, err error) {
	// Move robot to the next cell without checking any condition. Next cell may be north, south, west and east cell
	// and only one step. Nothing is movable other than robots so checking if it is robot or not
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

	// After movement defined and state machine defined, check the movement if it is inside the board
	if CheckLowerUpperBound(newRow, newColumn) {
		newRow = row
		newColumn = column
		newFace = face
		err = errors.New("out of index boundries")
	}

	// Robots can also attack. Normally robot cannot move to any other cell if the cell is not empty but in attack mode
	// robot can move to non-empty cell which is allocated with dinosaur
	if attack {
		if CheckType(newRow, newColumn) == "d" {
			err = nil
		}

	} else {
		// In case robot is not in attack mode, checking the cell if it is empty or not to perform movement or not
		if !CheckCell(newRow, newColumn) {
			newRow = row
			newColumn = column
			newFace = face
			err = errors.New("robot can move to only empty cell")
		}
	}

	return
}
