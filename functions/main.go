package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"png2jpg/image"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ReqBody struct {
	URL       string `json:"url"`
	ImageName string `json:"image_name"`
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	body := ReqBody{}
	json.Unmarshal([]byte(request.Body), &body)

	out, err := os.CreateTemp("", body.ImageName)
	defer os.Remove(out.Name())
	if err != nil {
		return Response{StatusCode: 503, Body: err.Error()}, err
	}
	defer out.Close()

	resp, err := http.Get(body.URL)
	if err != nil {
		return Response{StatusCode: 503, Body: err.Error()}, err
	}
	defer resp.Body.Close()

	imageCotent, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{StatusCode: 503, Body: err.Error()}, err
	}
	jpegContent, err := image.ToJpeg(image.Imager{}, imageCotent)
	if err != nil {
		return Response{StatusCode: 503, Body: err.Error()}, err
	}

	b64img := base64.StdEncoding.EncodeToString(jpegContent)

	return Response{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":   "image/jpeg",
			"Content-Length": fmt.Sprintf("%d", len(b64img)),
		},
		Body:            b64img,
		IsBase64Encoded: true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
