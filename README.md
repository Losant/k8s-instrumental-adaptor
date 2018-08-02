# Kubernetes Instrumental Adaptor
This repo is an implementation of the Kubernetes ExternalMetrics API. Much of the code is pulled from these two very helpful OSS projects:

- [Google Stackdriver Adaptor](https://github.com/GoogleCloudPlatform/k8s-stackdriver/tree/master/custom-metrics-stackdriver-adapter)
- [custom-metrics](https://github.com/kubernetes-incubator/custom-metrics-apiserver)

Currently, this adaptor only implements the ExternalMetricsProvider. At the moment, all of the metrics we want are stored in Instrumental as "global" metrics and we didn't need to implement the CustomMetricsProvider. CustomMetricsProvider is a way to talk directly to the kubernetes objects, which we don't need at this time.  If you would like to contribute to that, please create an issue and let us know how you would like to integrate the kube objects with Instrumental

### Instrumental Client
As part of this project, we implemented a client wraps an HTTP call to the Instrumental API. Currently it is fairly simple, but if it is something that could be used outside of this project, we could create a separate repo and add much more functionality to this client.

##### Usage
TODO: 
- add available functions
- add info about INSTRUMENTAL_TOKEN

### Adaptor Installation and Usage