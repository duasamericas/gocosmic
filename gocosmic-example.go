

package main

import (
	"errors"
  "log"
	"github.com/aslanvaroqua/gocosmic/gocosmic"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

// Handler is your Lambda function handler # ToDo make event from cosmicjs
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID

  return GoCosmic.Get(url){
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(Handler)
}
