package entity

type PlanetBuilding struct {
	UUID         string `json:"uuid"`
	DistrictUUID string `json:"district_uuid"`
	Type         string `json:"type"`
	Level        int    `json:"level"`
}

const (
	PLANET_BUILDING_TYPE_CAPITAL        = "CAPITAL"
	PLANET_BUILDING_TYPE_RESEARCH       = "RESEARCH"
	PLANET_BUILDING_TYPE_ADMINISTRATION = "ADMINISTRATION"
	PLANET_BUILDING_TYPE_FOUNDRY        = "FOUNDRY"
	PLANET_BUILDING_TYPE_FACTORY        = "FACTORY"
	PLANET_BUILDING_TYPE_ENTERTAINMENT  = "ENTERTAINMENT"
	PLANET_BUILDING_TYPE_COMMERCE       = "COMMERCE"
	PLANET_BUILDING_TYPE_MINING         = "MINING"
	PLANET_BUILDING_TYPE_ENERGY         = "ENERGY"
	PLANET_BUILDING_TYPE_FARM           = "FARM"
)
