package instrumental

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Client struct {
	BaseURL           string
	InstrumentalToken string

	HttpClient *http.Client
}

// Query is used to build the URL string passed to the request.Query.Query
// Duration, Resolution, and Time are optional and will default to Instrumental's defaults
type Query struct {
	Path       string
	MetricName string
	Duration   int
	Resolution int
	Time       int
}

const DefaultBaseURL = "https://instrumentalapp.com/"

// NewClient creates a client with the Intrumental token. Optionally, pass in your own client.
func NewClient(httpClient *http.Client, token string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{
		HttpClient: httpClient,
	}

	c.BaseURL = DefaultBaseURL
	c.InstrumentalToken = token

	return c
}

// GetInstrumentalMetric takes a Query object, makes the API cal to Instrumental, and
// returns an *InstrumentalMetric.
func (c *Client) GetInstrumentalMetric(query Query) (*InstrumentalMetric, error) {
	var im InstrumentalMetric
	req, err := c.NewQueryRequest(query)
	if err != nil {
		log.Printf("%v", err)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&im)
	if err != nil {
		log.Println(err)
	}

	return &im, nil
}

// NewQueryRequest sets up Instrumental query and returns an *http.Request and an error
func (c *Client) NewQueryRequest(query Query) (*http.Request, error) {
	url := buildBaseURL(c, query)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Instrumental-Token", c.InstrumentalToken)

	q := req.URL.Query()
	if query.Duration != 0 {
		q.Add("duration", strconv.Itoa(query.Duration))
	}
	if query.Resolution != 0 {
		q.Add("resolution", strconv.Itoa(query.Resolution))
	}
	if query.Time != 0 {
		q.Add("time", strconv.Itoa(query.Time))
	}
	req.URL.RawQuery = q.Encode()

	return req, nil
}

// buildBaseURL creates the base of the request string.
// Ex: https://instrumentalapp.com/api/2/metrics/<metricName>
// "api/2/metrics" and metricName come from the Query object.
func buildBaseURL(c *Client, query Query) string {

	path := query.Path
	metric := query.MetricName

	qs := c.BaseURL + path + metric

	return qs
}
