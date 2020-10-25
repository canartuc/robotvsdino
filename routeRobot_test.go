package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCreateRobot tests robot creation by sending dummy request and manipulating response
func TestCreateRobot(t *testing.T) {
	InitBoard()

	ts := httptest.NewServer(startGinServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s%s%s/create/%d/%s/%s",
		ts.URL, REQUESTGROUP, ROBOTGROUP, 39, "ac", "right"))

	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %v", http.StatusCreated, resp.StatusCode)
	}

	// TODO: Rather than hiding ResponseStruct here in test, we can use it globally for responses rather than hardcoded
	type ResponseStruct struct {
		Msg         string `json:"msg"`
		RobotRow    int    `json:"robot_row"`
		RobotColumn int    `json:"robot_column"`
		RobotFace   string `json:"robot_face"`
	}

	var resMap ResponseStruct
	err = json.NewDecoder(resp.Body).Decode(&resMap)
	if err != nil {
		t.Fatalf("Cannot decode json to map, %v", err)
	}

	expectedResRow := 38
	expectedResCol := 28
	expectedResFace := "east"

	if resMap.RobotColumn != expectedResCol ||
		resMap.RobotRow != expectedResRow ||
		resMap.RobotFace != expectedResFace {
		t.Fatalf("Expected (row:col:face) %d:%d:%s / Got (row:col:face) %d:%d:%s",
			expectedResRow, expectedResCol, expectedResFace,
			resMap.RobotRow, resMap.RobotColumn, resMap.RobotFace)
	}

}

// TODO: Extend test coverage with TestMoveRobot
