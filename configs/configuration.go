package configuration

import (
	"os"
)

var SatanaSelected = os.Getenv("SATANA_SELECTED") == "true"
var ReselectSatana = os.Getenv("RESELECT_SATANA") == "true"
var DBFileLocation = os.Getenv("DB_FILE")
