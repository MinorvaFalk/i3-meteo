package meteo

import "encoding/json"

type Places []Place

func (p Places) MarshalBinary() (data []byte, err error) {
	return json.Marshal(p)
}

type Place struct {
	Name     string `json:"name"`
	PlaceId  string `json:"place_id"`
	AdmArea1 string `json:"adm_area1"`
	AdmArea2 string `json:"adm_area2"`
	Country  string `json:"country"`
	Lat      string `json:"lat"`
	Lon      string `json:"lon"`
	Timezone string `json:"timezone"`
	Type     string `json:"type"`
}

func (p Place) MarshalBinary() (data []byte, err error) {
	return json.Marshal(p)
}
