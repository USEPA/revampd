package main

import (
	"net/http"
	"os"
	"encoding/json"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

// Ping endpoint for API
func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// FindUnitsByOperatingYear returns all of the units associated with the given year
func findUnitsByOperatingYear(w http.ResponseWriter, r *http.Request, dbService *DatabaseService) {

	years, ok := r.URL.Query()["operatingYear"]
	
	// Check to see if the operating year was supplied
	if !ok || len(years[0]) < 1 {
		log.Debug("Operating year is missing.")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Operating year is required"))
		return		
	}

	// Query()["key"] will return an array of items, 
	// we only want the single item.
	yearString := years[0]

	// Check to see if the operating year is a valid int
	year, err := strconv.Atoi(yearString)
	if err != nil {
		log.Debug("Can't convert year to int.")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Operating year is invalid"))
		return		
	}

	units, err := dbService.FindUnitsByOperatingYear(year)

	// Check the results
	if err == nil {
		if len(units) == 0 {
			log.Debug("No units found for year ", year)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Operating year is required"))
			return
		} else {
			jUnits, err := json.Marshal(units)
			if err == nil {
				w.Write(jUnits)
			}
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return		
	}
}

func main() {

	dbService, err := CreateDatabaseService()

	if err != nil {
		log.Fatal("Could create database service: " + err.Error())
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		ping(w, r)
	})

	http.HandleFunc("/units/findByOperatingYear", func(w http.ResponseWriter, r *http.Request) {
			findUnitsByOperatingYear(w, r, dbService)
	})

	log.Debug("Starting revAMPD API backend, listening on port " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
