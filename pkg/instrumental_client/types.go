package instrumental

type (
	// InstrumentalMetric
	InstrumentalMetric struct {
		Version  int      `json:"version"`
		Flags    int      `json:"flags"`
		Response Response `json:"response"`
	}

	// Response
	Response struct {
		Metrics []Metric `json:"metrics"`
		Notices []Notice `json:"notices"`
	}

	// Metric
	Metric struct {
		ID         string `json:"id"`
		ProjectID  int    `json:"project_id"`
		Expression string `json:"expression"`
		Name       string `json:"name"`
		Type       string `json:"type"`
		CreatedAt  int    `json:"created_at"`
		UpdatedAt  int    `json:"updated_at"`
		Values     Value  `json:"values"`
	}

	// Notice
	Notice interface{}

	Value struct {
		Start      int    `json:"start"`
		Stop       int    `json:"stop"`
		Resolution int    `json:"resolution"`
		Duration   int    `json:"duration"`
		Data       []Data `json:"data"`
	}

	// Data
	Data struct {
		Sum     float64 `json:"s"`
		Count   int     `json:"c"`
		Average float64 `json:"a"`
	}
)
