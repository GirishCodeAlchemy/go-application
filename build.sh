set -euxo pipefail

# mkdir -p "$(pwd)/functions"
# GOBIN=$(pwd)/functions go install ./...
# # chmod +x "$(pwd)"/functions/*
# go env

mkdir -p "$(pwd)/functions"
GOOS=linux GOARCH=amd64 go build -o $(pwd)functions/main ./src/first.go