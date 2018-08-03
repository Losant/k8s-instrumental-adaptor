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
	"strconv"
	"time"

	instrumental "github.com/losant/k8s-instrumental-adaptor/pkg/instrumental_client"
	"github.com/losant/k8s-instrumental-adaptor/pkg/provider"
	apierr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/metrics/pkg/apis/custom_metrics"
	"k8s.io/metrics/pkg/apis/external_metrics"
)

type externalMetric struct {
	info  provider.ExternalMetricInfo
	value external_metrics.ExternalMetricValue
}

type instrumentalProvider struct {
	instrumentalClient *instrumental.Client
	externalMetrics    []externalMetric
}

// type fakeInstrumentalCustomProvider struct{}

// func NewFakeInstrumentalCustomProvider() provider.CustomMetricsProvider {
// 	return &fakeInstrumentalCustomProvider{}
// }

func NewInstrumentalProvider(token string) provider.MetricsProvider {
	fmt.Println("Creating NewInstrumentalProvider")
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	instrumentalClient := instrumental.NewClient(client, token)
	return &instrumentalProvider{
		instrumentalClient: instrumentalClient,
		externalMetrics:    []externalMetric{},
	}
}

func (ip *instrumentalProvider) GetRootScopedMetricByName(groupResource schema.GroupResource, name string, metricName string) (*custom_metrics.MetricValue, error) {
	return &custom_metrics.MetricValue{}, nil
}
func (ip *instrumentalProvider) GetRootScopedMetricBySelector(groupResource schema.GroupResource, selector labels.Selector, metricName string) (*custom_metrics.MetricValueList, error) {
	return &custom_metrics.MetricValueList{}, nil
}
func (ip *instrumentalProvider) GetNamespacedMetricByName(groupResource schema.GroupResource, namespace string, name string, metricName string) (*custom_metrics.MetricValue, error) {
	return &custom_metrics.MetricValue{}, nil
}
func (ip *instrumentalProvider) GetNamespacedMetricBySelector(groupResource schema.GroupResource, namespace string, selector labels.Selector, metricName string) (*custom_metrics.MetricValueList, error) {
	return &custom_metrics.MetricValueList{}, nil
}

func (ip *instrumentalProvider) ListAllMetrics() []provider.CustomMetricInfo {
	fcmi := []provider.CustomMetricInfo{}
	return fcmi
}

func (ip *instrumentalProvider) GetExternalMetric(namespace string, metricName string, metricSelector labels.Selector) (*external_metrics.ExternalMetricValueList, error) {
	fmt.Println("Calling GetExternalMetric func")
	fmt.Printf("%v\n", metricName)
	metrics := []external_metrics.ExternalMetricValue{}
	metricLabels := map[string]string{}

	q := instrumental.Query{
		Path:       "api/2/metrics/",
		Duration:   120,
		Resolution: 60,
		MetricName: metricName,
	}
	metric, err := ip.instrumentalClient.GetInstrumentalMetric(q)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	fmt.Printf("Results returned from Instrumental: %v\n", metric)

	for i, v := range metric.Response.Metrics {
		values := metric.Response.Metrics[i].Values

		l := len(values.Data)
		point := values.Data[l-1]
		endTime, err := time.Parse(time.RFC3339, strconv.Itoa(values.Stop))
		if err != nil {
			return nil, apierr.NewInternalError(fmt.Errorf("Timeseries from Instrumental has incorrect end time: %v", values.Stop))
		}

		value := point.Average
		metricValue := external_metrics.ExternalMetricValue{
			Timestamp:  metav1.NewTime(endTime),
			MetricName: metricName,
			MetricLabels: map[string]string{
				metricLabels["resource.type"]: v.Type,
				metricLabels["resource.name"]: v.Name,
			},
		}
		metricValue.Value = *resource.NewMilliQuantity(int64(value*1000), resource.DecimalSI)
		metrics = append(metrics, metricValue)
	}

	fmt.Println(metrics)

	return &external_metrics.ExternalMetricValueList{
		Items: metrics,
	}, nil
}

func (ip *instrumentalProvider) ListAllExternalMetrics() []provider.ExternalMetricInfo {
	externalMetricsInfo := []provider.ExternalMetricInfo{}
	for _, metric := range ip.externalMetrics {
		externalMetricsInfo = append(externalMetricsInfo, metric.info)
	}
	return externalMetricsInfo
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
