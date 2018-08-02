/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"fmt"
	"log"
	"net/http"
	"time"

	instrumental "github.com/losant/k8s-instrumental-adaptor/pkg/instrumental_client"
	"github.com/losant/k8s-instrumental-adaptor/pkg/provider"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/metrics/pkg/apis/external_metrics"
)

// TODO(kawych):
// * Support long responses from Stackdriver (pagination).

// type clock interface {
// 	Now() time.Time
// }

// type realClock struct{}

// func (c realClock) Now() time.Time {
// 	return time.Now()
// }

type externalMetric struct {
	info  provider.ExternalMetricInfo
	value external_metrics.ExternalMetricValue
}

// InstrumentalProvider is a provider of custom metrics from Instrumental.
type InstrumentalProvider struct {
	instrumentalClient *instrumental.Client
	externalMetrics    []externalMetric
}

// NewStackdriverProvider creates a StackdriverProvider
func NewStackdriverProvider() provider.ExternalMetricsProvider {
	var client *http.Client
	client = &http.Client{
		Timeout: time.Second * 10,
	}
	instrumentalClient := instrumental.NewClient(client, "asjkdlf")
	return &InstrumentalProvider{
		instrumentalClient: instrumentalClient,
		// externalMetrics:    []externalMetric{},
	}
}

func (i *InstrumentalProvider) GetExternalMetric(namespace string, metricName string, metricSelector labels.Selector) (*external_metrics.ExternalMetricValueList, error) {
	// Build the query (ideally get metricName, duration?, resolution?, and time? from yaml)
	q := instrumental.Query{
		Path:       "api/2/metrics",
		Duration:   120,
		Resolution: 60,
	}

	metric, err := i.instrumentalClient.GetInstrumentalMetric(q)
	if err != nil {
		log.Printf("Unable to get Instrumental metrics: %v\n", err)
	}
	fmt.Println(metric)

	metricValues := []external_metrics.ExternalMetricValue{}
	// for _, v := metric.Response.Metrics[0] {
	// 	if len(v.Values.Data) <= 0 {
	// 		return apierr.NewInternalError(fmt.Errorf("Empty time series returned from Stackdriver"))
	// 	}

	// 	// Get the second to last data point.  The last data point is very volitile and changes often.
	// 	dataLength = len(v.Values.Data)
	// 	point := v.Values.Data[dataLength - 1]
	// 	endTime, err := time.Parse(time.PRC3339, v.Values.Stop)
	// 	if err != nil {
	// 		return nil, apierr.NewInternalError(fmt.Errorf("Timeseries from Instrumental has incorrect end time: %s", val.Stop))
	// 	}
	// 	metricValue := external_metrics.ExternalMetricValue{
	// 		Timestamp: metav1.NewTime(endTime),
	// 		MetricName: metricName,
	// 		MetricLabels: map[string]string{
	// 			metricLabels["resource.type"]: v.Type,
	// 			metricLabels["resource.name"]: v.Name,
	// 		}
	// 	}
	// 	value := point.A
	// 	switch {
	// 	case value.int64 != nil:
	// 		metricValue.Value = *resource.NewQuantity(value, resource.DecimalSI)
	// 	case value.float64 != nil:
	// 		metricValue.Value = *resource.NewMilliQuantity(int64(value*1000), resource.DecimalSI)
	// 	default:
	// 		return nil, apierr.NewBadRequest(fmt.Sprintf("Expected metric of type DoubleValue or Int64Value, but received TypedValue: %v", value))
	// 	}
	// 	metricValues = append(metricValues, metricValue)
	// }

	return &external_metrics.ExternalMetricValueList{
		Items: metricValues,
	}, nil
}

// ListAllExternalMetrics returns a list of available external metrics.
// Not implemented (currently returns empty list).
func (i *InstrumentalProvider) ListAllExternalMetrics() []provider.ExternalMetricInfo {
	return []provider.ExternalMetricInfo{}
}

/*
These are the interface functions for CustomMetricsProvider
If we need to collect data on a per pods (or some other Kube object), then implement these.

Note: If/when the CustomMetricsProvider is implemented, you will need to change the
NewStackdriverProvider to return provider.MetricsProvider instead of provider.ExternalMetricsProvider.

provider.MetricsProvider is simply an interface that implements both the CustomMetricsProvider
and the ExternalMetricsProvider.
*/

/*
func (i *InstrumentalProvider) GetRootScopedMetricByName(groupResource schema.GroupResource, name string, metricName string) (*custom_metrics.MetricValue, error) {
	return &custom_metrics.MetricValue{}
}
func (i *InstrumentalProvider) GetRootScopedMetricBySelector(groupResource schema.GroupResource, selector labels.Selector, metricName string) (*custom_metrics.MetricValueList, error) {
	return &custom_metrics.MetricValue{}
}
func (i *InstrumentalProvider) GetNamespacedMetricByName(groupResource schema.GroupResource, namespace string, name string, metricName string) (*custom_metrics.MetricValue, error) {
	return &custom_metrics.MetricValue{}
}
func (i *InstrumentalProvider) GetNamespacedMetricBySelector(groupResource schema.GroupResource, namespace string, selector labels.Selector, metricName string) (*custom_metrics.MetricValueList, error) {
	return &custom_metrics.MetricValue{}
}
func (i *InstrumentalProvider) ListAllMetrics() []CustomMetricInfo {
	return &custom_metrics.MetricValue{}
}
*/
