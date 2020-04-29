package main

//The Unit entity
type Unit struct {
	OperatingYear *int `db:"op_year" json:"operatingYear"`
	UnitID *string `db:"unitid" json:"unitID"` 
	EPARegion *int `db:"epa_region"` 
	EPAState *string `db:"state"` 
	FacilityName *string `db:"facility_name" json:"facilityName"` 
	OrisCode *int `db:"oris_code" json:"orisCode"` 
	StackId *string `db:"stack_ids" json:"stackId"`
	OperatingStatus	*string `db:"op_status" json:"operatingStatus"` 
	ProgramCode *string `db:"program_code" json:"programCode"` 
	UnitType *string `db:"unit_type_description" json:"unitType"` 
	PrimaryFuel *string `db:"primary_fuel_type_desc" json:"primaryFuel"`
	PrimaryFuelGroup *string `db:"primary_fuel_group" json:"primaryFuelGroup"` 
	OperatingTime *int64 `db:"an_count_op_time" json:"operatingTime"` 
	GrossLoad *int64 `db:"an_gload" json:"grossLoad"` 
	SteamLoad *int64 `db:"an_sload" json:"steamLoad"` 
	HeatInput *float64 `db:"an_heat_input" json:"heatInput"` 
	Co2Mass *float64 `db:"an_co2_mass" json:"co2Mass"`
	So2Mass *float64 `db:"an_so2_mass" json:"so2Mass"` 
	NoxMass *float64 `db:"an_nox_mass" json:"noxMass"`
}

