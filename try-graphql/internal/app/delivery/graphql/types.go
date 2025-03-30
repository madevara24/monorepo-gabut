package graphql

type Planet struct {
	UUID      string              `json:"uuid"`
	Name      string              `json:"name"`
	Size      int                 `json:"size"`
	Type      string              `json:"type"`
	Districts []*PlanetDistrict   `json:"districts"`
	Features  []*PlanetaryFeature `json:"features"`
}

type PlanetDistrict struct {
	UUID      string            `json:"uuid"`
	Type      string            `json:"type"`
	Level     int               `json:"level"`
	Buildings []*PlanetBuilding `json:"buildings"`
}

type PlanetBuilding struct {
	UUID  string `json:"uuid"`
	Type  string `json:"type"`
	Level int    `json:"level"`
}

type PlanetaryFeature struct {
	UUID string `json:"uuid"`
	Type string `json:"type"`
}
