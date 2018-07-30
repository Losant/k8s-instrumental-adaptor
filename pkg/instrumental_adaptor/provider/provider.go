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
	kubeClient *corev1.CoreV1Client
	// instrumentalURL     string
	// config              *config.GceConfig
	rateInterval time.Duration
	// translator          *Translator
	useNewResourceModel bool
}

// NewStackdriverProvider creates a StackdriverProvider
func NewStackdriverProvider(kubeClient *corev1.CoreV1Client, mapper apimeta.RESTMapper, rateInterval, alignmentPeriod time.Duration, useNewResourceModel bool) provider.MetricsProvider {
	// gceConf, err := config.GetGceConfig()
	if err != nil {
		glog.Fatalf("Failed to retrieve GCE config: %v", err)
	}

	return &InstrumentalProvider{
		kubeClient: kubeClient,
		// stackdriverService: stackdriverService,
		// config:             gceConf,
		rateInterval: rateInterval,
		// translator: &Translator{
		// 	service:             stackdriverService,
		// 	config:              gceConf,
		// 	reqWindow:           rateInterval,
		// 	alignmentPeriod:     alignmentPeriod,
		// 	clock:               realClock{},
		// 	mapper:              mapper,
		// 	useNewResourceModel: useNewResourceModel,
		// },
	}
}

// GetRootScopedMetricByName queries Stackdriver for metrics identified by name and not associated
// with any namespace. Current implementation doesn't support root scoped metrics.
// func (p *StackdriverProvider) GetRootScopedMetricByName(groupResource schema.GroupResource, name string, escapedMetricName string) (*custom_metrics.MetricValue, error) {
// 	if !p.translator.useNewResourceModel {
// 		return nil, provider.NewOperationNotSupportedError("Get root scoped metric by name")
// 	}
// 	if groupResource.Resource != nodeResource {
// 		return nil, provider.NewOperationNotSupportedError(fmt.Sprintf("Get root scoped metric by name for resource %q", groupResource.Resource))
// 	}
// 	matchingNode, err := p.kubeClient.Nodes().Get(name, metav1.GetOptions{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	metricKind, err := p.translator.GetMetricKind(getCustomMetricName(escapedMetricName))
// 	if err != nil {
// 		return nil, err
// 	}
// 	stackdriverRequest, err := p.translator.GetSDReqForNodes(&v1.NodeList{Items: []v1.Node{*matchingNode}}, getCustomMetricName(escapedMetricName), metricKind)
// 	if err != nil {
// 		return nil, err
// 	}
// 	stackdriverResponse, err := stackdriverRequest.Do()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return p.translator.GetRespForSingleObject(stackdriverResponse, groupResource, escapedMetricName, "", name)
// }

// // GetRootScopedMetricBySelector queries Stackdriver for metrics identified by selector and not
// // associated with any namespace. Current implementation doesn't support root scoped metrics.
// func (p *StackdriverProvider) GetRootScopedMetricBySelector(groupResource schema.GroupResource, selector labels.Selector, escapedMetricName string) (*custom_metrics.MetricValueList, error) {
// 	if !p.translator.useNewResourceModel {
// 		return nil, provider.NewOperationNotSupportedError("Get root scoped metric by selector")
// 	}
// 	if groupResource.Resource != nodeResource {
// 		return nil, provider.NewOperationNotSupportedError(fmt.Sprintf("Get root scoped metric by selector for resource %q", groupResource.Resource))
// 	}
// 	matchingNodes, err := p.kubeClient.Nodes().List(metav1.ListOptions{LabelSelector: selector.String()})
// 	if err != nil {
// 		return nil, err
// 	}
// 	metricKind, err := p.translator.GetMetricKind(getCustomMetricName(escapedMetricName))
// 	if err != nil {
// 		return nil, err
// 	}
// 	result := []custom_metrics.MetricValue{}
// 	for i := 0; i < len(matchingNodes.Items); i += oneOfMax {
// 		nodesSlice := &v1.NodeList{Items: matchingNodes.Items[i:min(i+oneOfMax, len(matchingNodes.Items))]}
// 		stackdriverRequest, err := p.translator.GetSDReqForNodes(nodesSlice, getCustomMetricName(escapedMetricName), metricKind)
// 		if err != nil {
// 			return nil, err
// 		}
// 		stackdriverResponse, err := stackdriverRequest.Do()
// 		if err != nil {
// 			return nil, err
// 		}
// 		slice, err := p.translator.GetRespForMultipleObjects(stackdriverResponse, p.translator.getNodeItems(matchingNodes), groupResource, escapedMetricName)
// 		if err != nil {
// 			return nil, err
// 		}
// 		result = append(result, slice...)
// 	}
// 	return &custom_metrics.MetricValueList{Items: result}, nil
// }

// // GetNamespacedMetricByName queries Stackdriver for metrics identified by name and associated
// // with a namespace.
// func (p *StackdriverProvider) GetNamespacedMetricByName(groupResource schema.GroupResource, namespace string, name string, escapedMetricName string) (*custom_metrics.MetricValue, error) {
// 	if groupResource.Resource != podResource {
// 		return nil, provider.NewOperationNotSupportedError(fmt.Sprintf("Get namespaced metric by name for resource %q", groupResource.Resource))
// 	}
// 	matchingPod, err := p.kubeClient.Pods(namespace).Get(name, metav1.GetOptions{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	metricKind, err := p.translator.GetMetricKind(getCustomMetricName(escapedMetricName))
// 	if err != nil {
// 		return nil, err
// 	}
// 	stackdriverRequest, err := p.translator.GetSDReqForPods(&v1.PodList{Items: []v1.Pod{*matchingPod}}, getCustomMetricName(escapedMetricName), metricKind, namespace)
// 	if err != nil {
// 		return nil, err
// 	}
// 	stackdriverResponse, err := stackdriverRequest.Do()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return p.translator.GetRespForSingleObject(stackdriverResponse, groupResource, escapedMetricName, namespace, name)
// }

// // GetNamespacedMetricBySelector queries Stackdriver for metrics identified by selector and associated
// // with a namespace.
// func (p *StackdriverProvider) GetNamespacedMetricBySelector(groupResource schema.GroupResource, namespace string, selector labels.Selector, escapedMetricName string) (*custom_metrics.MetricValueList, error) {
// 	if groupResource.Resource != podResource {
// 		return nil, provider.NewOperationNotSupportedError(fmt.Sprintf("Get namespaced metric by selector for resource %q", groupResource.Resource))
// 	}
// 	matchingPods, err := p.kubeClient.Pods(namespace).List(metav1.ListOptions{LabelSelector: selector.String()})
// 	if err != nil {
// 		return nil, err
// 	}
// 	metricKind, err := p.translator.GetMetricKind(getCustomMetricName(escapedMetricName))
// 	if err != nil {
// 		return nil, err
// 	}
// 	result := []custom_metrics.MetricValue{}
// 	for i := 0; i < len(matchingPods.Items); i += oneOfMax {
// 		podsSlice := &v1.PodList{Items: matchingPods.Items[i:min(i+oneOfMax, len(matchingPods.Items))]}
// 		stackdriverRequest, err := p.translator.GetSDReqForPods(podsSlice, getCustomMetricName(escapedMetricName), metricKind, namespace)
// 		if err != nil {
// 			return nil, err
// 		}
// 		stackdriverResponse, err := stackdriverRequest.Do()
// 		if err != nil {
// 			return nil, err
// 		}
// 		slice, err := p.translator.GetRespForMultipleObjects(stackdriverResponse, p.translator.getPodItems(matchingPods), groupResource, escapedMetricName)
// 		if err != nil {
// 			return nil, err
// 		}
// 		result = append(result, slice...)
// 	}
// 	return &custom_metrics.MetricValueList{Items: result}, nil
// }

// // ListAllMetrics returns all custom metrics available from Stackdriver.
// // List only pod metrics
// func (p *StackdriverProvider) ListAllMetrics() []provider.CustomMetricInfo {
// 	metrics := []provider.CustomMetricInfo{}
// 	stackdriverRequest := p.translator.ListMetricDescriptors()
// 	response, err := stackdriverRequest.Do()
// 	if err != nil {
// 		glog.Errorf("Failed request to stackdriver api: %s", err)
// 		return metrics
// 	}
// 	return p.translator.GetMetricsFromSDDescriptorsResp(response)
// }

// // GetExternalMetric queries Stackdriver for external metrics.
// func (p *StackdriverProvider) GetExternalMetric(namespace string, metricNameEscaped string, metricSelector labels.Selector) (*external_metrics.ExternalMetricValueList, error) {
// 	metricName := getExternalMetricName(metricNameEscaped)
// 	metricKind, err := p.translator.GetMetricKind(metricName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	stackdriverRequest, err := p.translator.GetExternalMetricRequest(metricName, metricKind, metricSelector)
// 	if err != nil {
// 		return nil, err
// 	}
// 	stackdriverResponse, err := stackdriverRequest.Do()
// 	if err != nil {
// 		return nil, err
// 	}
// 	externalMetricItems, err := p.translator.GetRespForExternalMetric(stackdriverResponse, metricNameEscaped)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &external_metrics.ExternalMetricValueList{
// 		Items: externalMetricItems,
// 	}, nil
// }

// // ListAllExternalMetrics returns a list of available external metrics.
// // Not implemented (currently returns empty list).
// func (p *StackdriverProvider) ListAllExternalMetrics() []provider.ExternalMetricInfo {
// 	return []provider.ExternalMetricInfo{}
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
