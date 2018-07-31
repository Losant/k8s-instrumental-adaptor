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
	"time"

	"github.com/golang/glog"
	"github.com/losant/k8s-instrumental-adaptor/pkg/provider"
	// corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	// apimeta "k8s.io/apimachinery/pkg/api/meta"

	instrumental "github.com/losant/k8s-instrumental-adaptor/pkg/instrumental_client"
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
	externalMetrics []externalMetric
}

client := &http.Client{
	Timeout: time.Second() * 10,
}
instrumentalClient := instrumental.NewClient(client, "asjkdlf")

// NewStackdriverProvider creates a StackdriverProvider
func NewStackdriverProvider() provider.ExternalMetricsProvider {
	return &InstrumentalProvider{
		instrumentalClient: instrumentalClient
		externalProvider: nil
	}
}

func (i *InstrumentalProvider) GetExternalMetric(namespace string, metricName string, metricSelector labels.Selector) (*external_metrics.ExternalMetricValueList, error) {
	// Build the query (ideally get metricName, duration?, resolution?, and time? from yaml)
	q := instrumental.Query{
		Path: "api/2/metrics",
		Duration: 120,
		Resolution: 60,
	}

	// This could all be done in the client itself ???
	req, err := i.instrumentalClient.NewQueryRequest(q)
	if err != nil {
		log.Printf("There was a problem requesting the Instrumental API: %v", err)
	}
	resp, err := instrumentalClient.HttpClient.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}

	defer resp.Body.Close()

	// Really, convert to struct and then get a specific average value out and put it
	// in externalMetricValueList{}.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("%T", body)

	fmt.Println(string(body))
	
	// Replace this with the external metric
	return &external_metrics.ExternalMetricValueList{
		Items: []ExternalMeticValue{}
	}
}

func (i *InstrumentalProvider) 	ListAllExternalMetrics() []provider.ExternalMetricInfo {
	return &provider.ExternalMetricInfo{}
}

// ################################################################################
// These are the interface functions for CustomMetricsProvider
// If we need to collect data on a per pods (or some other Kube object), then implement these
// ################################################################################

// func (i *InstrumentalProvider) GetRootScopedMetricByName(groupResource schema.GroupResource, name string, metricName string) (*custom_metrics.MetricValue, error) {
// 	return &custom_metrics.MetricValue{}
// }
// func (i *InstrumentalProvider) GetRootScopedMetricBySelector(groupResource schema.GroupResource, selector labels.Selector, metricName string) (*custom_metrics.MetricValueList, error) {
// 	return &custom_metrics.MetricValue{}
// }
// func (i *InstrumentalProvider) GetNamespacedMetricByName(groupResource schema.GroupResource, namespace string, name string, metricName string) (*custom_metrics.MetricValue, error) {
// 	return &custom_metrics.MetricValue{}
// }
// func (i *InstrumentalProvider) GetNamespacedMetricBySelector(groupResource schema.GroupResource, namespace string, selector labels.Selector, metricName string) (*custom_metrics.MetricValueList, error) {
// 	return &custom_metrics.MetricValue{}
// }
// func (i *InstrumentalProvider) ListAllMetrics() []CustomMetricInfo {
// 	return &custom_metrics.MetricValue{}
// }


// DELETE THIS - IT'S JUST FOR REFERENCE
// ################################################################################

// func (t *Translator) GetRespForExternalMetric(response *stackdriver.ListTimeSeriesResponse, metricName string) ([]external_metrics.ExternalMetricValue, error) {
// 	metricValues := []external_metrics.ExternalMetricValue{}
// 	for _, series := range response.TimeSeries {
// 		if len(series.Points) <= 0 {
// 			// This shouldn't happen with correct query to Stackdriver
// 			return nil, apierr.NewInternalError(fmt.Errorf("Empty time series returned from Stackdriver"))
// 		}
// 		// Points in a time series are returned in reverse time order
// 		point := series.Points[0]
// 		endTime, err := time.Parse(time.RFC3339, point.Interval.EndTime)
// 		if err != nil {
// 			return nil, apierr.NewInternalError(fmt.Errorf("Timeseries from Stackdriver has incorrect end time: %s", point.Interval.EndTime))
// 		}
// 		metricValue := external_metrics.ExternalMetricValue{
// 			Timestamp:    metav1.NewTime(endTime),
// 			MetricName:   metricName,
// 			MetricLabels: t.getMetricLabels(series),
// 		}
// 		value := *point.Value
// 		switch {
// 		case value.Int64Value != nil:
// 			metricValue.Value = *resource.NewQuantity(*value.Int64Value, resource.DecimalSI)
// 		case value.DoubleValue != nil:
// 			metricValue.Value = *resource.NewMilliQuantity(int64(*value.DoubleValue*1000), resource.DecimalSI)
// 		default:
// 			return nil, apierr.NewBadRequest(fmt.Sprintf("Expected metric of type DoubleValue or Int64Value, but received TypedValue: %v", value))
// 		}
// 		metricValues = append(metricValues, metricValue)
// 	}
// 	return metricValues, nil
// }

// func (t *Translator) getMetricLabels(series *stackdriver.TimeSeries) map[string]string {
// 	metricLabels := map[string]string{}
// 	for label, value := range series.Metric.Labels {
// 		metricLabels["metric.labels."+label] = value
// 	}
// 	metricLabels["resource.type"] = series.Resource.Type
// 	for label, value := range series.Resource.Labels {
// 		metricLabels["resource.labels."+label] = value
// 	}
// 	return metricLabels
// }

// type ExternalMetricValueList struct {
// 	metav1.TypeMeta `json:",inline"`
// 	metav1.ListMeta `json:"metadata,omitempty"`

// 	// value of the metric matching a given set of labels
// 	Items []ExternalMetricValue `json:"items"`
// }

// // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// // a metric value for external metric
// // A single metric value is identified by metric name and a set of string labels.
// // For one metric there can be multiple values with different sets of labels.
// type ExternalMetricValue struct {
// 	metav1.TypeMeta `json:",inline"`

// 	// the name of the metric
// 	MetricName string `json:"metricName"`

// 	// a set of labels that identify a single time series for the metric
// 	MetricLabels map[string]string `json:"metricLabels"`

// 	// indicates the time at which the metrics were produced
// 	Timestamp metav1.Time `json:"timestamp"`

// 	// indicates the window ([Timestamp-Window, Timestamp]) from
// 	// which these metrics were calculated, when returning rate
// 	// metrics calculated from cumulative metrics (or zero for
// 	// non-calculated instantaneous metrics).
// 	WindowSeconds *int64 `json:"window,omitempty"`

// 	// the value of the metric
// 	Value resource.Quantity `json:"value"`
// }

// ################################################################################


// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
// func getExternalMetricName(metricName string) string {
// 	return strings.Replace(metricName, "|", "/", -1)
// }

// func getCustomMetricName(metricName string) string {
// 	if strings.Contains(metricName, "|") {
// 		return getExternalMetricName(metricName)
// 	}
// 	return "custom.googleapis.com/" + metricName
// }
