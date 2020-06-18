package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func findUnitsByOperatingYear(w http.ResponseWriter, r *http.Request, dbService *DatabaseService) {

	years := r.URL.Query().Get("operatingYear")
	limitString := r.URL.Query().Get("limit")
	offsetString := r.URL.Query().Get("offset")

	// Check to see if the operating year was supplied
	if len(years) < 1 {
		log.Debug("A valid operating year is required")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Operating year is required"))
		return
	}

	// Check to see if the operating year is a valid int
	year, err := strconv.Atoi(years)
	if err != nil || year < 0 {
		log.Debug("Can't convert year to int.")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("A valid operating year is required"))
		return
	}
	//using default if there no limit or offset
	if len(limitString) < 1 {
		log.Debug("setting limit to 100")
		limitString = "100"
	}
	limit, err := strconv.Atoi(limitString)
	if err != nil || limit < 0 {
		log.Debug("Can't convert limit to int.")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid limit. Is it a positive integer?"))
		return
	}

	if len(offsetString) < 1 {
		log.Debug("setting offset to 0")
		offsetString = "0"
	}
	offset, err := strconv.Atoi(offsetString)
	if err != nil || offset < 0 {
		log.Debug("Can't convert offset to int. Is it a positive integer?")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid offset. Is it a positive integer?"))
		return
	}

	units, err := dbService.paginatedUnitsByOperatingYear(year, limit, offset)
	total, err2 := dbService.getTotalNumberOfRows(year)
	// Check the results
	if err == nil && err2 == nil {
		if len(units) == 0 {
			log.Debug("No units found for year ", year)
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("No units were found"))
			return
		} else {
			//jUnits, err := json.Marshal(units)
			var payload *Payload
			payload = new(Payload)
			payload.Units = units
			payload.MetaData.Retrieved = strconv.Itoa(len(units))
			payload.MetaData.Total = strconv.Itoa(total)
			jPayload, err := json.Marshal(payload)
			if err == nil {
				w.Write(jPayload)
			}
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//Allows all sites to access the API. This can be fine tuned
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func main() {

	dbService, err := CreateDatabaseService()

	if err != nil {
		log.Fatal("Could create database service: " + err.Error())
	}

	http.HandleFunc("/units/findByOperatingYear", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		findUnitsByOperatingYear(w, r, dbService)
	})

	log.Debug("Starting revAMPD API backend, listening on port " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
