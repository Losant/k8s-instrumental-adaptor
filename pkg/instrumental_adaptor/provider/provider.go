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
	"errors"
	"fmt"
	"log"

	instrumental "github.com/losant/k8s-instrumental-adaptor/pkg/instrumental_client"
	"github.com/losant/k8s-instrumental-adaptor/pkg/provider"
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
	externalMetrics []externalMetric
	translator      *Translator
}

// NewInstrumentalProvider creates the provider for interacting with Instrumental.
func NewInstrumentalProvider(token string, instrumentalClient *instrumental.Client) provider.MetricsProvider {
	fmt.Println("Creating NewInstrumentalProvider")
	translator := &Translator{
		instrumentalClient: instrumentalClient,
	}

	return &instrumentalProvider{
		externalMetrics: []externalMetric{},
		translator:      translator,
	}
}

// These functions are part of the CustomMetrics interface. If you need to interact with the k8s objects directly, implement these functions.
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

// The GetExternalMetric and ListAllExternalMetrics function implement the ExternalMetrics interface.
func (ip *instrumentalProvider) GetExternalMetric(namespace string, metricName string, metricSelector labels.Selector) (*external_metrics.ExternalMetricValueList, error) {

	// Convert metricName to Camelcase (apimachinery is calling ToLower on the metricName)
	camelMetricName := getCustomMetricName(metricName)
	q := instrumental.Query{
		Path:       "api/2/metrics/",
		Duration:   120,
		Resolution: 60,
		MetricName: camelMetricName,
	}
	metric, err := ip.translator.instrumentalClient.GetInstrumentalMetric(q)
	if err != nil {
		return nil, errors.New("The call to Instrumental returned an error")
	}
	log.Printf("Results returned from Instrumental: %v\n\n", metric)

	metrics, err := ip.translator.GetRespForExternalMetric(metric, camelMetricName)
	if err != nil {
		return nil, err
	}

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
