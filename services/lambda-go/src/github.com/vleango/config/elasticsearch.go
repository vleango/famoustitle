package config

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/sha1sum/aws_signing_client"
	"net/http"
	"os"
)

var ElasticSearchHost string
var ESClient *http.Client

func init() {
	ESClient = &http.Client{}

	switch os.Getenv("APP_ENV") {
	case "development":
		ElasticSearchHost = "http://datastore-es:9200"
	case "test":
		ElasticSearchHost = "http://datastore-es-test:9200"
	case "ci":
		ElasticSearchHost = "http://localhost:9200"
	case "production":
		ElasticSearchHost = os.Getenv("ELASTICSEARCH_HOST_URL")

		var myClient *http.Client
		signer := v4.NewSigner(credentials.NewStaticCredentials(os.Getenv("AWS_ID"), os.Getenv("AWS_SECRET"), ""))
		var awsClient, err = aws_signing_client.New(signer, myClient, "es", DefaultRegion)
		if err != nil {
			fmt.Println(err.Error())
		}
		ESClient = awsClient
	}
}
