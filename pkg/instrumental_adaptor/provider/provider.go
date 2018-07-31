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
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	instrumental "github.com/losant/k8s-instrumental-adaptor/pkg/instrumental_client"

	// "github.com/GoogleCloudPlatform/k8s-stackdriver/custom-metrics-stackdriver-adapter/pkg/config"
	// "github.com/GoogleCloudPlatform/k8s-stackdriver/custom-metrics-stackdriver-adapter/pkg/provider"

	apimeta "k8s.io/apimachinery/pkg/api/meta"
)

// TODO(kawych):
// * Support long responses from Stackdriver (pagination).

type clock interface {
	Now() time.Time
}

type realClock struct{}

func (c realClock) Now() time.Time {
	return time.Now()
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
	
	
	// Replace this with the external metric
	return &external_metrics.ExternalMetricValueList{}
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
