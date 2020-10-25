package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDinosaur(t *testing.T) {
	InitBoard()

	ts := httptest.NewServer(startGinServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s%s%s/create/%d/%s",
		ts.URL, REQUESTGROUP, DINOSAURGROUP, 39, "ad"))

	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %v", http.StatusCreated, resp.StatusCode)
	}

	// TODO: Rather than hiding ResponseStruct here in test, we can use it globally for responses rather than hardcoded
	type ResponseStruct struct {
		Msg            string `json:"msg"`
		DinosaurRow    int    `json:"dinosaur_row"`
		DinosaurColumn int    `json:"dinosaur_column"`
	}

	var resMap ResponseStruct
	err = json.NewDecoder(resp.Body).Decode(&resMap)
	if err != nil {
		t.Fatalf("Cannot decode json to map, %v", err)
	}

	expectedResRow := 38
	expectedResCol := 29

	if resMap.DinosaurColumn != expectedResCol ||
		resMap.DinosaurRow != expectedResRow {
		t.Fatalf("Expected (row:col) %d:%d / Got (row:col) %d:%d",
			expectedResRow, expectedResCol, resMap.DinosaurRow, resMap.DinosaurColumn)
	}

}
