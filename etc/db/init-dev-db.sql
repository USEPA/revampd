-- Create the unit universe table
\i /data/unit_universe_ddl.sql

-- Copy the test data into the table
COPY UNIT_UNIVERSE(
  OP_YEAR, 
  UNIT_ID, 
  EPA_REGION, 
  STATE, 
  FACILITY_NAME, 
  ORIS_CODE, 
  UNITID, 
  STACK_IDS, 
  OP_STATUS, 
  PROGRAM_CODE, 
  UNIT_TYPE_DESCRIPTION, 
  PRIMARY_FUEL_TYPE_DESC, 
  PRIMARY_FUEL_GROUP,
  AN_COUNT_OP_TIME,
  AN_GLOAD,
  AN_SLOAD,
  AN_HEAT_INPUT,
  AN_CO2_MASS,
  AN_SO2_MASS,
  AN_NOX_MASS 
  )
FROM '/data/Y2017-2019 unit-universe-data.csv' DELIMITER ',' CSV HEADER;
