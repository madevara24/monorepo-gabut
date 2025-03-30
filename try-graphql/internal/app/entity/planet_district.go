package entity

type PlanetDistrict struct {
	UUID       string `json:"uuid"`
	PlanetUUID string `json:"planet_uuid"`
	Type       string `json:"type"`
	Level      int    `json:"level"`
}

const (
	PLANET_DISTRICT_TYPE_CITY     = "CITY"
	PLANET_DISTRICT_TYPE_MINING   = "MINING"
	PLANET_DISTRICT_TYPE_ENERGY   = "ENERGY"
	PLANET_DISTRICT_TYPE_FARM     = "FARM"
	PLANET_DISTRICT_TYPE_INDUSTRY = "INDUSTRY"
)
