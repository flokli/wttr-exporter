package main

type JSONResp struct {
	CurrentCondition []CurrentCondition `json:"current_condition"`
}

type CurrentCondition struct {
	FeelsLikeC float64 `json:"FeelsLikeC,string"`
	FeelsLikeF float64 `json:"FeelsLikeF,string"`
	CloudCover float64 `json:"cloudcover,string"`
	Humidity float64 `json:"humidity,string"`
	LocalObsDateTime string `json:"localObsDateTime"`
	ObservationTime string `json:"observation_time"`
	PrecipitationMM float64 `json:"precipMM,string"`
	PressureMBar float64 `json:"pressure,string"`
	TempC float64 `json:"temp_C,string"`
	TempF float64 `json:"temp_F,string"`
	UVIndex float64 `json:"uvIndex,string"`
	Visibility float64 `json:"visibility,string"`
	WinddirDegree float64 `json:"winddirDegree,string"`
	WindspeedKmpH float64 `json:"windspeedKmph,string"`
	WindspeedMiles float64 `json:"windspeedMiles,string"`
}
