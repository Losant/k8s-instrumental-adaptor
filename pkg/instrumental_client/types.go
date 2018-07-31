package instrumental

type InstrumentalMetric struct {
	Version  int `json:"version"`
	Flags    int `json:"flags"`
	Response struct {
		Metrics []struct {
			ID         string `json:"id"`
			ProjectID  int    `json:"project_id"`
			Expression string `json:"expression"`
			Name       string `json:"name"`
			Type       string `json:"type"`
			CreatedAt  int    `json:"created_at"`
			UpdatedAt  int    `json:"updated_at"`
			Values     struct {
				Start      int `json:"start"`
				Stop       int `json:"stop"`
				Resolution int `json:"resolution"`
				Duration   int `json:"duration"`
				Data       []struct {
					Sum     float64 `json:"s"`
					Count   int     `json:"c"`
					Average float64 `json:"a"`
				} `json:"data"`
			} `json:"values"`
		} `json:"metrics"`
		Notices []interface{} `json:"notices"`
	} `json:"response"`
}
