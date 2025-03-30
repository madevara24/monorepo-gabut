package entity

type Planet struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Size int    `json:"size"`
	Type string `json:"type"`
}

const (
	PLANET_TYPE_TROPICAL    = "TROPICAL"
	PLANET_TYPE_OCEAN       = "OCEAN"
	PLANET_TYPE_CONTINENTAL = "CONTINENTAL"
	PLANET_TYPE_SAVANNA     = "SAVANNA"
	PLANET_TYPE_DESERT      = "DESERT"
	PLANET_TYPE_ARID        = "ARID"
	PLANET_TYPE_TUNDRA      = "TUNDRA"
	PLANET_TYPE_ARCTIC      = "ARCTIC"
	PLANET_TYPE_ALPINE      = "ALPINE"
)
