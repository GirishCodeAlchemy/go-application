package main

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler for API Gateway requests
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// Serve files from the "static" directory
	fs := http.FileServer(http.Dir("static"))

	// Create a request that can be handled by the file server
	req, err := http.NewRequest(request.HTTPMethod, request.Path, nil)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}

	// Capture the response using a ResponseRecorder
	rr := &responseRecorder{ResponseWriter: http.NewResponseWriter(), statusCode: 200}
	fs.ServeHTTP(rr, req)

	return &events.APIGatewayProxyResponse{
		StatusCode: rr.statusCode,
		Body:       rr.body,
		Headers:    map[string]string{"Content-Type": rr.header.Get("Content-Type")},
	}, nil
}

func main() {
	// Ensure the "static" directory exists
	if _, err := os.Stat("static"); os.IsNotExist(err) {
		os.Mkdir("static", os.ModePerm)
	}

	lambda.Start(handler)
}

// Custom response recorder to capture the HTTP response
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       string
	header     http.Header
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *responseRecorder) Write(body []byte) (int, error) {
	r.body = string(body)
	return r.ResponseWriter.Write(body)
}

func (r *responseRecorder) Header() http.Header {
	if r.header == nil {
		r.header = http.Header{}
	}
	return r.header
}
