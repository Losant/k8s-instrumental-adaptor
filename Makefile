ARCH?=386
OUT_DIR?=build
PACKAGE=github.com/losant/k8s-instrumental-adapter
PREFIX?=gcr.io/structure-1104
TAG = v0.0.1
PKG := $(shell find pkg/* -type f)

.PHONY: build docker push test clean

build: build/adapter

build/adapter: main.go $(PKG)
	CGO_ENABLED=0 GOARCH=$(ARCH) GOOS=linux go build -a -o $(OUT_DIR)/$(ARCH)/adapter main.go

docker: build/adapter
	docker build --pull -t ${PREFIX}/custom-metrics-instrumental-adapter:$(TAG) .

push: docker
	gcloud docker -- push ${PREFIX}/custom-metrics-instrumental-adapter:$(TAG)

test: $(PKG)
	CGO_ENABLED=0 go test ./pkg/...

clean:
	rm -rf build
