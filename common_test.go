package main

import (
	"fmt"
	"testing"
)

func TestMoveRobotState(t *testing.T) {
	InitBoard()
	// Test edge cases for movement - This is not a regular test so we are forcing to fail here
	var tests = []struct {
		row, col int
		face     string
		errMes   string
	}{
		{LOWERBOUND, LOWERBOUND, "west", "out of index boundries"},
		{LOWERBOUND, LOWERBOUND, "north", "out of index boundries"},
		{LOWERBOUND, UPPERBOUND, "east", "out of index boundries"},
		{LOWERBOUND, UPPERBOUND, "north", "out of index boundries"},
		{UPPERBOUND, LOWERBOUND, "west", "out of index boundries"},
		{UPPERBOUND, LOWERBOUND, "south", "out of index boundries"},
		{UPPERBOUND, UPPERBOUND, "east", "out of index boundries"},
		{UPPERBOUND, UPPERBOUND, "south", "out of index boundries"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Edge case (row:col:move): %d:%d:%s", tt.row, tt.col, tt.face)
		t.Run(testName, func(t *testing.T) {
			Board[tt.row][tt.col] = fmt.Sprintf("r:%s", tt.face)
			_, _, _, err := MoveRobotState(tt.row, tt.col, tt.face)
			if err != nil {
				if err.Error() != tt.errMes {
					t.Errorf("Edge case failed (row:col:move): %d:%d:%s", tt.row, tt.col, tt.face)
				}
			}
		})
	}
}
