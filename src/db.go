package main

import (
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

//DatabaseService type
type DatabaseService struct {
	database *sqlx.DB
}

//CreateDatabaseService for the API
func CreateDatabaseService() (*DatabaseService, error) {

	var db *sqlx.DB
	var err error

	retries := 5
	log.Debug("Getting a database connection. ", os.Getenv("DATABASE_URL"))

	//Sometinme the API starts before the database is up
	for i := 0; i < retries; i++ {
		db, err = sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Debug("Waiting to retry connection...")
			time.Sleep(time.Duration(5 * i * 1e9))
		}
	}

	if err != nil {
		log.Debug("Could not connect to the database.")
		return nil, err
	}

	dbService := DatabaseService {
		database: db,
	}

	return &dbService, nil
}

//FindUnitsByOperatingYear returns the units for the given year
func (dbService *DatabaseService) FindUnitsByOperatingYear(year int) ([]Unit, error) {
	log.Debug("Finding units for year ", year)
	rows, err := dbService.database.Queryx("SELECT unitid, op_year, facility_name, oris_code, state, epa_region, unit_type_description, stack_ids, op_status, program_code, primary_fuel_type_desc, primary_fuel_group, an_count_op_time, an_gload, an_sload, an_heat_input, an_co2_mass, an_so2_mass, an_nox_mass FROM UNIT_UNIVERSE where op_year = $1 limit 100", year)

	units := []Unit{}
		
	for rows.Next() {
		var unit Unit
		err = rows.StructScan(&unit)

		units = append(units, unit)

		if err != nil {
			return nil, err
		}
	}

	return units, nil
}
