package provider

import (
	"errors"
	"log"
	"regexp"
	"strings"
	"time"

	instrumental "github.com/losant/k8s-instrumental-adaptor/pkg/instrumental_client"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/external_metrics"
)

type Translator struct {
	instrumentalClient *instrumental.Client
}

// GetRespForExternalMetric takes the response from the Instrumental client and returns a slice of ExternalMetricValue or an error.
func (t *Translator) GetRespForExternalMetric(response *instrumental.InstrumentalMetric, metricName string) ([]external_metrics.ExternalMetricValue, error) {
	metrics := []external_metrics.ExternalMetricValue{}

	for i, v := range response.Response.Metrics {
		values := response.Response.Metrics[i].Values
		l := len(values.Data)
		// Because Instrumental forces Duration to be at least twice the Resolution,
		// (i.e. there with be at least two points) This should be reflected in the tests.
		point := values.Data[l-2]
		endTime := time.Unix(int64(values.Stop), 0)

		value := point.Average
		log.Printf("\n\tValue (Average): %f\n\n", value)
		if value <= 0 {
			// This shouldn't happen with correct query to Stackdriver
			return nil, errors.New("Empty time series returned from Instrumental")
		}
		metricValue := external_metrics.ExternalMetricValue{
			Timestamp:  metav1.NewTime(endTime),
			MetricName: metricName,
			MetricLabels: map[string]string{
				"resource.type": v.Type,
				"resource.name": v.Name,
			},
		}

		metricValue.Value = *resource.NewMilliQuantity(int64(value*1000), resource.DecimalSI)
		metrics = append(metrics, metricValue)
	}

	return metrics, nil
}

func getExternalMetricName(metricName string) string {
	re := regexp.MustCompile(`\|[a-z]?`)
	out := re.ReplaceAllStringFunc(metricName, strings.ToUpper)
	out = strings.ReplaceAll(out, "|", "")
	return out;
}

func getCustomMetricName(metricName string) string {
	if strings.Contains(metricName, "|") {
		return getExternalMetricName(metricName)
	}
	return metricName
}
