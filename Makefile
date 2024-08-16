.PHONY: run clean docs

build:
	mkdir -p functions
	cp -r ./api/static functions/
	cp ./api/go.mod go.mod
	go install ./...
	# GOBIN=./api go install ./...
	GOOS=linux GOARCH=amd64 go build -o functions/main ./api/lambda_handler.go