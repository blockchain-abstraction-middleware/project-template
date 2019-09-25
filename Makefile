install:
	go install -v

serve:
	go run cmd/serve/main.go

deps: 
	go get ./...