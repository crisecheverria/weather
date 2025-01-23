package types

type WeatherResponse struct {
	List []struct {
		DtTxt string `json:"dt_txt"`
		Main  struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"list"`
}
