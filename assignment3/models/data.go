package models

type Data struct {
	WindSpeed    int
	WaterQuality int
	WindStatus   string
	WaterStatus  string
}

type Response struct {
	WindSpeed    int    `json:"wind_speed"`
	WaterQuality int    `json:"water_quality"`
	WindStatus   string `json:"wind_status"`
	WaterStatus  string `json:"water_status"`
}

func (d Data) ToResponse() Response {
	return Response{
		WindSpeed:    d.WindSpeed,
		WaterQuality: d.WaterQuality,
		WindStatus:   d.WindStatus,
		WaterStatus:  d.WaterStatus,
	}
}
