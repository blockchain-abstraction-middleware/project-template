VERSION := 0.0.1

DOCKER_REG = bamdockerhub
DOCKER_IMAGE = project-template
DOCKER_IMAGE_TAG = $(VERSION)
USER = $(shell whoami)

docker-build:
	docker build -t $(DOCKER_REG)/$(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG) .

docker-push:
	docker push $(DOCKER_REG)/$(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG)

docker-build-dev:
	docker build -t $(DOCKER_REG)/$(DOCKER_IMAGE):$(USER) .

docker-push-dev:
	docker push $(DOCKER_REG)/$(DOCKER_IMAGE):$(USER)

install:
	go install -v

serve:
	cd cmd/serve; go run main.go

build:
	go build ./...

deps:
	go get ./...

test:
	go test ./... -cover

lint: 
	gofmt -s -w .
	gofmt -r '(a) -> a' -l *.go .
	gofmt -r '(a) -> a' -w *.go .
	gofmt -r 'α[β:len(α)] -> α[β:]' -w $GOROOT/src .

.PHONY: \
	install \
	build \
	serve \
	deps \
	lint \
	test \