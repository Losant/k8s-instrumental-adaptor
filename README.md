# Kubernetes Instrumental Adaptor
This repo is an implementation of the Kubernetes ExternalMetrics API. Much of the code is pulled from these two very helpful OSS projects:

- [Google Stackdriver Adaptor](https://github.com/GoogleCloudPlatform/k8s-stackdriver/tree/master/custom-metrics-stackdriver-adapter)
- [custom-metrics](https://github.com/kubernetes-incubator/custom-metrics-apiserver)

Currently, this adaptor only implements the ExternalMetricsProvider. At the moment, all of the metrics we want are stored in Instrumental as "global" metrics and we didn't need to implement the CustomMetricsProvider. CustomMetricsProvider is a way to talk directly to the kubernetes objects, which we don't need at this time. However, they are all stubbed out. If you would like to contribute to that, please create an issue and let us know how you would like to integrate the kube objects with Instrumental

### Instrumental Client
As part of this project, we implemented a client that wraps an HTTP call to the Instrumental API. Currently it is fairly simple, but if it is something that could be used outside of this project, we could create a separate repo and add much more functionality to this client.

##### Usage
TODO: 
- add available functions
- add info about INSTRUMENTAL_TOKEN

### Adaptor Installation and Usage
To use this adaptor, you will need to deploy the yaml files in the ./deploy directory to your kubernetes cluster.  This will setup all the RBAC, service account, deployment, etc needed. The docker image provided by default is built for linux 386 machines. If you need a different build, you will need to clone the project and change the Makefile to use the GOOS and GOARCH or your chose. You will need to do this if you want to host the docker image in a different docker repository as well.

To deploy, first create a secret for your `INSTRUMENTAL_TOKEN`. Add your unique `INSTRUMENTAL_TOKEN` to the token.yaml file and run:

```bash
kubectl apply -f deploy/token.yaml
```

**NOTE: Be sure to keep your secrets secure. We recommend not committing any secrets to you repo.**

Then, create the deployment by running:

```bash
kubectl apply -f deploy/instrumental_adaptor.yaml
```

This will create all the adaptor resources in a namespace called `custom-metrics`. To see these resources run:

```bash
kubectl get all -n custom-metrics
```

**NOTE:** If you are running on GKE, you may need to enable cluster-admin for your user. To do so, run:

```bash
kubectl create clusterrolebinding cluster-admin-binding \
--clusterrole cluster-admin --user <USER_ACCOUNT>
```

Once the adaptor is setup, you will then need to create a deployment for your application. There is a sample yaml file in the examples directory. You will need to have an application that writes some Instrumental metrics within your app. When your app is ready to go, change the fields noted in the examples/simple.yaml file to match the metrics you want to scale on.

To check that everything is working, you can look at the logs produced by the adaptor. First, get the name of the po for the adaptor.

```bash
kubectl get po -n custom-metrics
```

```output
NAME                                                   READY     STATUS    RESTARTS   AGE
custom-metrics-instrumental-adapter-478374838-39ff1   1/1       Running   0          5m
```

Then follow the logs for this container:

```bash
kubectl logs custom-metrics-instrumental-adapter-478374838-39ff1 -f
```