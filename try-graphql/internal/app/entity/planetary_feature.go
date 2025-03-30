package entity

type PlanetaryFeature struct {
	UUID       string `json:"uuid"`
	PlanetUUID string `json:"planet_uuid"`
	Type       string `json:"type"`
}

const (
	PLANETARY_FEATURE_TYPE_MINING_1            = "MINING_1"
	PLANETARY_FEATURE_TYPE_MINING_2            = "MINING_2"
	PLANETARY_FEATURE_TYPE_MINING_3            = "MINING_3"
	PLANETARY_FEATURE_TYPE_MINING_UNDERWATER_1 = "MINING_U_1"
	PLANETARY_FEATURE_TYPE_MINING_UNDERWATER_2 = "MINING_U_2"
	PLANETARY_FEATURE_TYPE_MINING_UNDERWATER_3 = "MINING_U_3"
	PLANETARY_FEATURE_TYPE_ENERGY_1            = "ENERGY_1"
	PLANETARY_FEATURE_TYPE_ENERGY_2            = "ENERGY_2"
	PLANETARY_FEATURE_TYPE_ENERGY_3            = "ENERGY_3"
	PLANETARY_FEATURE_TYPE_ENERGY_UNDERWATER_1 = "ENERGY_U_1"
	PLANETARY_FEATURE_TYPE_ENERGY_UNDERWATER_2 = "ENERGY_U_2"
	PLANETARY_FEATURE_TYPE_ENERGY_UNDERWATER_3 = "ENERGY_U_3"
	PLANETARY_FEATURE_TYPE_FARM_1              = "FARM_1"
	PLANETARY_FEATURE_TYPE_FARM_2              = "FARM_2"
	PLANETARY_FEATURE_TYPE_FARM_3              = "FARM_3"
)
