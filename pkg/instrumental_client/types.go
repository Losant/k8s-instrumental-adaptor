package instrumental

import "time"

type InstrumentalMetric struct {
	Version  int      `json:"version"`
	Flags    int      `json:"flags"`
	Response Response `json:"response"`
	Notices  []Notice `json:"notices"`
}

// Notice - Not sure what should go in here yet ???
type Notice struct{}

type Response struct {
	Metrics []Metric `json:"metrics"`
}

type Metric struct {
	ID         string    `json:"id"`
	ProjectID  int       `json:"projectId"`
	Expression string    `json:"expression"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Values     Values    `json:"values"`
}

type Values struct {
	Start      time.Time    `json:"start"`
	Stop       time.Time    `json:"stop"`
	Resolution int          `json:"resolution"`
	Duration   int          `json:"duration"`
	Data       []DataPoints `json:"data"`
}

type DataPoints struct {
	Sum     int64   `json:"s"`
	Count   int64   `json:"c"`
	Average float64 `json:"a"`
}
