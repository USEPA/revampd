package main

//The Unit entity
type Unit struct {
	operatingYear *int `db:"op_year"`
	unitID *string `db:"unitid"` 
	EPARegion *int `db:"epa_region"` 
	EPAState *string `db:"state"` 
	facilityName *string `db:"facility_name"` 
	orisCode *int `db:"oris_code"` 
	stackId *string `db:"stack_ids"`
	operatingStatus	*string `db:"op_status"` 
	programCode *string `db:"program_code"` 
	unitType *string `db:"unit_type_description"` 
	primaryFuel *string `db:"primary_fuel_type_desc"`
	primaryFuelGroup *string `db:"primary_fuel_group"` 
	operatingTime *int64 `db:"an_count_op_time"` 
	grossLoad *int64 `db:"an_gload"` 
	steamLoad *int64 `db:"an_sload"` 
	heatInput *float64 `db:"an_heat_input"` 
	co2Mass *float64 `db:"an_co2_mass"`
	so2Mass *float64 `db:"an_so2_mass"` 
	noxMass *float64 `db:"an_nox_mass"`
}

