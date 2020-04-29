package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

var dbService *DatabaseService

func TestMain(m *testing.M) {
	var err error

	dbService, err = CreateDatabaseService()

	if (err != nil || dbService == nil) {
		log.Fatal("Could not create database service: " + err.Error())
	}
  
	code := m.Run()
 
	os.Exit(code)
}

func TestFindUnitsByOperatingYear(t *testing.T) {
	req, _ := http.NewRequest("GET", "/units/findByOperatingYear?operatingYear=2017", nil)
	response := httptest.NewRecorder()

	if err := dbService.database.Ping(); err != nil {
			log.Println("Database Error ", err)
		}

	findUnitsByOperatingYear(response, req, dbService)

	if isExpectedResponseCode(t, http.StatusOK, response.Code) {
		units := make([]Unit,0)
		err := json.Unmarshal(response.Body.Bytes(), &units)
		if err != nil {
			t.Errorf("Could not unmarshall JSON response: %s", err)
		}
	}
}

func TestFindUnitsBadOperatingYear(t *testing.T) {
	req, _ := http.NewRequest("GET", "/units/findByOperatingYear?operatingYear=AAAA", nil)
	response := httptest.NewRecorder()

	findUnitsByOperatingYear(response, req, dbService)

	isExpectedResponseCode(t, http.StatusBadRequest, response.Code)
}

func TestFindUnitsNoOperatingYear(t *testing.T) {
	req, _ := http.NewRequest("GET", "/units/findByOperatingYear", nil)
	response := httptest.NewRecorder()

	findUnitsByOperatingYear(response, req, dbService)

	isExpectedResponseCode(t, http.StatusBadRequest, response.Code)	
}

func isExpectedResponseCode(t *testing.T, expected, actual int) (bool) {
    if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
		return false
	}
	return true
}