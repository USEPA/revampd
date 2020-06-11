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
	//using default if there no limit or offset
	if len(limitString) < 1 {
		log.Debug("setting limit to 100")
		limitString = "100"
	}
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		log.Debug("Can't convert limit to int.")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid limit"))
		return
	}

	if len(offsetString) < 1 {
		log.Debug("setting offset to 0")
		offsetString = "0"
	}
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		log.Debug("Can't convert offset to int.")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid offset"))
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	//yearString := years[0]

	// Check to see if the operating year is a valid int
	year, err := strconv.Atoi(years)
	if err != nil {
		log.Debug("Can't convert year to int.")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("A valid operating year is required"))
		return
	}

	units, err := dbService.paginatedUnitsByOperatingYear(year, limit, offset)

	// Check the results
	if err == nil {
		if len(units) == 0 {
			log.Debug("No units found for year ", year)
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("No units were found"))
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
