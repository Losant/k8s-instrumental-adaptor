package provider

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	instrumental "github.com/losant/k8s-instrumental-adaptor/pkg/instrumental_client"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/external_metrics"
)

var endTime int = 1533674340

func TestGetRespForExternalMetric(t *testing.T) {
	// Create fake instrumentalProvider
	translator := newFakeTranslator()
	fmt.Println(translator)

	// Create fake value - This is the metric to expect???
	var value float64 = 200.2

	// Create fake response
	response := generateFakeResponse()

	// fmt.Println(value)
	// fmt.Println(response)

	// Call GetRespForExternalMetric
	metrics, err := translator.GetRespForExternalMetric(response, "metric.fake.latency")
	if err != nil {
		t.Errorf("Failed to translate the metric to type []external_metrics.ExternalMetricValue: %v\n", err)
	}

	// Check if external_metrics.ExternalMetricValueList is correct (expect)
	expectedMetrics := []external_metrics.ExternalMetricValue{
		{
			Value:      *resource.NewMilliQuantity(int64(value*1000), resource.DecimalSI),
			Timestamp:  metav1.NewTime(time.Unix(int64(endTime), 0)),
			MetricName: "metric.fake.latency",
			MetricLabels: map[string]string{
				"resource.type": "gauge",
				"resource.name": "metric.fake.latency",
			},
		},
	}

	if len(metrics) != len(expectedMetrics) {
		t.Errorf("Unexpected result. Expected %d metrics, received %d", len(expectedMetrics), len(metrics))
	}
	for i := range metrics {
		if !reflect.DeepEqual(metrics[i], expectedMetrics[i]) {
			t.Errorf("Unexpected result. Expected: \n%v,\n received: \n%v", expectedMetrics[i], metrics[i])
		}
	}
}

func newFakeTranslator() *Translator {
	instrumentalClient := instrumental.NewClient(http.DefaultClient, "faketoken")

	return &Translator{
		instrumentalClient: instrumentalClient,
	}
}

func generateFakeResponse() *instrumental.InstrumentalMetric {
	var data []instrumental.Data
	d1 := instrumental.Data{
		Sum:     400.5,
		Count:   500,
		Average: 200.2,
	}
	d2 := instrumental.Data{
		Sum:     302.3,
		Count:   450,
		Average: 230.1,
	}
	data = append(data, d1)
	data = append(data, d2)

	value := instrumental.Value{
		Start:      1533672540,
		Stop:       endTime,
		Resolution: 60,
		Duration:   180,
		Data:       data,
	}

	var metrics []instrumental.Metric
	metric := instrumental.Metric{
		ID:         "metric.fake.latency",
		ProjectID:  111,
		Expression: "metric.fake.latency",
		Name:       "metric.fake.latency",
		Type:       "gauge",
		CreatedAt:  0,
		UpdatedAt:  0,
		Values:     value,
	}
	metrics = append(metrics, metric)

	response := instrumental.Response{
		Metrics: metrics,
		Notices: []instrumental.Notice{},
	}

	im := instrumental.InstrumentalMetric{
		Version:  2,
		Flags:    77777,
		Response: response,
	}

	return &im
}
