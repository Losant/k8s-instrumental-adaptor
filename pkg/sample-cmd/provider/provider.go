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
	"net/http"
	"time"

	"k8s.io/apimachinery/pkg/labels"

	"github.com/kubernetes-incubator/custom-metrics-apiserver/pkg/provider"
	instrumental "github.com/losant/k8s-instrumental-adaptor/pkg/instrumental_client"
	"k8s.io/metrics/pkg/apis/external_metrics"
)

type externalMetric struct {
	info  provider.ExternalMetricInfo
	value external_metrics.ExternalMetricValue
}

type testingProvider struct {
	instrumentalClient *instrumental.Client
	// client dynamic.Interface
	// mapper apimeta.RESTMapper

	// values          map[provider.CustomMetricInfo]int64
	externalMetrics []externalMetric
}

func NewFakeProvider() provider.ExternalMetricsProvider {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	instrumentalClient := instrumental.NewClient(client, "ajsdklf")
	return &testingProvider{
		instrumentalClient: instrumentalClient,
		// client:          client,
		// mapper:          mapper,
		// values:          make(map[provider.CustomMetricInfo]int64),
		externalMetrics: []externalMetric{},
	}
}

// func (p *testingProvider) valueFor(groupResource schema.GroupResource, metricName string, namespaced bool) (int64, error) {
// 	info := provider.CustomMetricInfo{
// 		GroupResource: groupResource,
// 		Metric:        metricName,
// 		Namespaced:    namespaced,
// 	}

// 	info, _, err := info.Normalized(p.mapper)
// 	if err != nil {
// 		return 0, err
// 	}

// 	value := p.values[info]
// 	value += 1
// 	p.values[info] = value

// 	return value, nil
// }

// func (p *testingProvider) metricFor(value int64, groupResource schema.GroupResource, namespace string, name string, metricName string) (*custom_metrics.MetricValue, error) {
// 	kind, err := p.mapper.KindFor(groupResource.WithVersion(""))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &custom_metrics.MetricValue{
// 		DescribedObject: custom_metrics.ObjectReference{
// 			APIVersion: groupResource.Group + "/" + runtime.APIVersionInternal,
// 			Kind:       kind.Kind,
// 			Name:       name,
// 			Namespace:  namespace,
// 		},
// 		MetricName: metricName,
// 		Timestamp:  metav1.NewTime(time.Now()), // Should to endTime
// 		Value:      *resource.NewMilliQuantity(value*100, resource.DecimalSI),
// 	}, nil
// }

// func (p *testingProvider) metricsFor(totalValue int64, groupResource schema.GroupResource, metricName string, list runtime.Object) (*custom_metrics.MetricValueList, error) {
// 	if !apimeta.IsListType(list) {
// 		return nil, fmt.Errorf("returned object was not a list")
// 	}

// 	res := make([]custom_metrics.MetricValue, 0)

// 	err := apimeta.EachListItem(list, func(item runtime.Object) error {
// 		objMeta := item.(metav1.Object)
// 		value, err := p.metricFor(0, groupResource, objMeta.GetNamespace(), objMeta.GetName(), metricName)
// 		if err != nil {
// 			return err
// 		}
// 		res = append(res, *value)

// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	for i := range res {
// 		res[i].Value = *resource.NewMilliQuantity(100*totalValue/int64(len(res)), resource.DecimalSI)
// 	}

// 	//return p.metricFor(value, groupResource, "", name, metricName)
// 	return &custom_metrics.MetricValueList{
// 		Items: res,
// 	}, nil
// }

// func (p *testingProvider) GetRootScopedMetricByName(groupResource schema.GroupResource, name string, metricName string) (*custom_metrics.MetricValue, error) {
// 	value, err := p.valueFor(groupResource, metricName, false)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return p.metricFor(value, groupResource, "", name, metricName)
// }

// func (p *testingProvider) GetRootScopedMetricBySelector(groupResource schema.GroupResource, selector labels.Selector, metricName string) (*custom_metrics.MetricValueList, error) {
// 	fullReses, err := p.mapper.ResourcesFor(groupResource.WithVersion(""))
// 	if err != nil || len(fullReses) == 0 {
// 		glog.Errorf("unable to get prefered GVRs for GR to list matching resource names: %v", err)
// 		// don't leak implementation details to the user
// 		return nil, apierr.NewInternalError(fmt.Errorf("unable to list matching resources"))
// 	}

// 	totalValue, err := p.valueFor(groupResource, metricName, false)
// 	if err != nil {
// 		return nil, err
// 	}

// 	matchingObjectsRaw, err := p.client.Resource(fullReses[0]).
// 		List(metav1.ListOptions{LabelSelector: selector.String()})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return p.metricsFor(totalValue, groupResource, metricName, matchingObjectsRaw)
// }

// func (p *testingProvider) GetNamespacedMetricByName(groupResource schema.GroupResource, namespace string, name string, metricName string) (*custom_metrics.MetricValue, error) {
// 	value, err := p.valueFor(groupResource, metricName, true)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return p.metricFor(value, groupResource, namespace, name, metricName)
// }

// func (p *testingProvider) GetNamespacedMetricBySelector(groupResource schema.GroupResource, namespace string, selector labels.Selector, metricName string) (*custom_metrics.MetricValueList, error) {
// 	fullReses, err := p.mapper.ResourcesFor(groupResource.WithVersion(""))
// 	if err != nil || len(fullReses) == 0 {
// 		glog.Errorf("unable to get prefered GVRs for GR to list matching resource names: %v", err)
// 		// don't leak implementation details to the user
// 		return nil, apierr.NewInternalError(fmt.Errorf("unable to list matching resources"))
// 	}

// 	totalValue, err := p.valueFor(groupResource, metricName, true)
// 	if err != nil {
// 		return nil, err
// 	}

// 	matchingObjectsRaw, err := p.client.Resource(fullReses[0]).Namespace(namespace).
// 		List(metav1.ListOptions{LabelSelector: selector.String()})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return p.metricsFor(totalValue, groupResource, metricName, matchingObjectsRaw)
// }

// func (p *testingProvider) ListAllMetrics() []provider.CustomMetricInfo {
// 	// TODO: maybe dynamically generate this?
// 	return []provider.CustomMetricInfo{
// 		{
// 			GroupResource: schema.GroupResource{Group: "", Resource: "pods"},
// 			Metric:        "packets-per-second",
// 			Namespaced:    true,
// 		},
// 		{
// 			GroupResource: schema.GroupResource{Group: "", Resource: "services"},
// 			Metric:        "connections-per-second",
// 			Namespaced:    true,
// 		},
// 		{
// 			GroupResource: schema.GroupResource{Group: "", Resource: "namespaces"},
// 			Metric:        "queue-length",
// 			Namespaced:    false,
// 		},
// 	}
// }
func (p *testingProvider) GetExternalMetric(namespace string, metricName string, metricSelector labels.Selector) (*external_metrics.ExternalMetricValueList, error) {
	metrics := []external_metrics.ExternalMetricValue{}
	// for _, metric := range p.externalMetrics {
	// 	if metric.info.Metric == metricName &&
	// 		metricSelector.Matches(labels.Set(metric.info.Labels)) {
	// 		metricValue := metric.value
	// 		metricValue.Timestamp = metav1.Now()
	// 		metrics = append(metrics, metricValue)
	// 	}
	// }
	return &external_metrics.ExternalMetricValueList{
		Items: metrics,
	}, nil
}

func (p *testingProvider) ListAllExternalMetrics() []provider.ExternalMetricInfo {
	externalMetricsInfo := []provider.ExternalMetricInfo{}
	for _, metric := range p.externalMetrics {
		externalMetricsInfo = append(externalMetricsInfo, metric.info)
	}
	return externalMetricsInfo
}
