package main

import (
	"context"
	"net/http"
	"os"
	"bytes"
	"fmt"
	"path/filepath"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func PrintDirTree(path string, indent string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for i, entry := range entries {
		prefix := "├── "
		if i == len(entries)-1 {
			prefix = "└── "
		}
		fmt.Println(indent + prefix + entry.Name())
		if entry.IsDir() {
			PrintDirTree(filepath.Join(path, entry.Name()), indent+"    ")
		}
	}
}

// Handler for API Gateway requests
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("Starting the handler code")
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}
	fmt.Println("Current Directory:", cwd)

	// Print the directory tree
	PrintDirTree(cwd, "")
	// Serve files from the "static" directory
	fs := http.FileServer(http.Dir("static"))

	// Create a custom ResponseRecorder
	rr := &responseRecorder{header: http.Header{}}
	fmt.Println("Starting the main code", rr)
	// Create a new HTTP request
	req, err := http.NewRequest(request.HTTPMethod, request.Path, nil)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}

	// Serve the HTTP request
	fs.ServeHTTP(rr, req)

	return &events.APIGatewayProxyResponse{
		StatusCode: rr.statusCode,
		Body:       rr.body.String(),
		Headers:    map[string]string{"Content-Type": rr.header.Get("Content-Type")},
	}, nil
}

func main() {
	fmt.Println("Starting the main code")
	// Ensure the "static" directory exists
	if _, err := os.Stat("static"); os.IsNotExist(err) {
		fmt.Println("Static file not present the main code")
		os.Mkdir("static", os.ModePerm)
	}
	lambda.Start(handler)
}

// Custom response recorder to capture the HTTP response
type responseRecorder struct {
	header     http.Header
	body       bytes.Buffer
	statusCode int
}

func (r *responseRecorder) Header() http.Header {
	return r.header
}

func (r *responseRecorder) Write(body []byte) (int, error) {
	return r.body.Write(body)
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
}
