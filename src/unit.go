package main

//The Unit entity
type Unit struct {
	OperationYear *int `db:"op_year"`
	UnitID *string `db:"unitid"` 
	EPARegion *int `db:"epa_region"` 
	State *string `db:"state"` 
	FacilityName *string `db:"facility_name"` 
	ORISCode *int `db:"oris_code"` 
	StackIDs *string `db:"stack_ids"`
	OperationStatus	*string `db:"op_status"` 
	ProgramCode *string `db:"program_code"` 
	UnitType *string `db:"unit_type_description"` 
	PrimaryFuel *string `db:"primary_fuel_type_desc"`
	PrimaryFuelGroup *string `db:"primary_fuel_group"` 
	OperationTime *int64 `db:"an_count_op_time"` 
	GrossLoad *int64 `db:"an_gload"` 
	SteamLoad *int64 `db:"an_sload"` 
	HeatInput *float64 `db:"an_heat_input"` 
	CO2Mass *float64 `db:"an_co2_mass"`
	SO2Mass *float64 `db:"an_so2_mass"` 
	NOXMass *float64 `db:"an_nox_mass"`
}

