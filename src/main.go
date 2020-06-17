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

	years, ok := r.URL.Query()["operatingYear"]

	// Check to see if the operating year was supplied
	if !ok || len(years[0]) < 1 {
		log.Debug("A valid operating year is required")
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
		w.Write([]byte("A valid operating year is required"))
		return
	}

	units, err := dbService.FindUnitsByOperatingYear(year)

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

//Enable Cross-Origin Resource Sharing (CORS)
func enableCors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	dbService, err := CreateDatabaseService()

	if err != nil {
		log.Fatal("Could create database service: " + err.Error())
	}

	http.HandleFunc("/units/findByOperatingYear", func(w http.ResponseWriter, r *http.Request) {
		findUnitsByOperatingYear(w, r, dbService)
	})

	fs := http.FileServer(http.Dir("./api-spec"))
	http.Handle("/swagger/", http.StripPrefix("/swagger", fs))

	log.Debug("Starting revAMPD API backend, listening on port " + os.Getenv("PORT"))
	//Enable CORS at the server level
	http.ListenAndServe(":"+os.Getenv("PORT"), enableCors(http.DefaultServeMux))
}
