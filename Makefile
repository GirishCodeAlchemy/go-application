.PHONY: run clean docs

netlify-build:
	mkdir -p functions
	cp -r ./api/static functions/
	cp ./api/go.mod go.mod
	cp ./api/go.sum go.sum
	go install ./...
	# GOBIN=./api go install ./...
	GOOS=linux GOARCH=amd64 go build -o functions/main ./api/lambda_handler.go

build:
	cp ./api/go.mod go.mod
	cp ./api/go.sum go.sum
	go install ./...
	GOOS=linux GOARCH=amd64 go build ./api/main.go