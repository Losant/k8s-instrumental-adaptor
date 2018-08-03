ARCH?=amd64
OUT_DIR?=build
PACKAGE=github.com/losant/k8s-instrumental-adapter
PREFIX?=gcr.io/structure-1104
TAG = v0.0.1
PKG := $(shell find pkg/* -type f)

.PHONY: build docker push test clean

build: build/adapter

build/adapter: main.go $(PKG)
	CGO_ENABLED=0 GOARCH=$(ARCH) go build -a -o $(OUT_DIR)/$(ARCH)/adapter main.go

docker: build/adapter
	docker build --pull -t ${PREFIX}/custom-metrics-instrumental-adapter:$(TAG) .

push: docker
	gcloud docker -- push ${PREFIX}/custom-metrics-instrumental-adapter:$(TAG)

test: $(PKG)
	CGO_ENABLED=0 go test ./pkg/...

clean:
	rm -rf build



# ARCH?=amd64
# OUT_DIR?=./_output

# .PHONY: all test verify-gofmt gofmt verify

# all: build
# build: vendor
# 	CGO_ENABLED=0 GOARCH=$(ARCH) go build -a -tags netgo -o $(OUT_DIR)/$(ARCH)/sample-adapter github.com/losant/k8s-instrumental-adaptor

# vendor: glide.lock
# 	glide install -v

# test: vendor
# 	CGO_ENABLED=0 go test ./pkg/...

# verify-gofmt:
# 	./hack/gofmt-all.sh -v

# gofmt:
# 	./hack/gofmt-all.sh

# verify: verify-gofmt test



# ARCH?=amd64
# OUT_DIR?=build
# PACKAGE=github.com/GoogleCloudPlatform/k8s-stackdriver/custom-metrics-stackdriver-adapter
# PREFIX?=gcr.io/google-containers
# TAG = v0.8.0
# PKG := $(shell find pkg/* -type f)

# .PHONY: build docker push test clean

# build: build/adapter

# build/adapter: adapter.go $(PKG)
# 	CGO_ENABLED=0 GOARCH=$(ARCH) go build -a -o $(OUT_DIR)/$(ARCH)/adapter adapter.go

# docker: build/adapter
# 	docker build --pull -t ${PREFIX}/custom-metrics-stackdriver-adapter:$(TAG) .

# push: docker
# 	gcloud docker -- push ${PREFIX}/custom-metrics-stackdriver-adapter:$(TAG)

# test: $(PKG)
# 	CGO_ENABLED=0 go test ./pkg/...

# clean:
# 	rm -rf build

