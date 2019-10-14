install:
	go install -v

serve:
	go run cmd/serve/main.go

build:
	go build ./...

deps:
	go get ./...

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