package test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/vleango/config"
	"github.com/vleango/lib/models"
	"net/http"
	"time"
)

var svc = config.DynamoSvc
var clusterName = "tech_writer_article"

func CleanDataStores() {
	CleanDB()
	CleanElasticSearch()
}

func CleanDB() {
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println(err)
	}
	for _, n := range result.TableNames {
		input := &dynamodb.DeleteTableInput{
			TableName: aws.String(*n),
		}
		_, err := svc.DeleteTable(input)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func CleanElasticSearch() {
	url := fmt.Sprintf("%v/%v", config.ElasticSearchHost, clusterName)

	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}

func CreateArticlesTable() {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("tech_writer_articles"),
	}

	_, err := svc.CreateTable(input)

	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
	}
}

func DefaultArticleModel() models.Article {
	return models.Article{
		Title:     "this is my title",
		Body:      "this is my body",
		Tags:      []string{"ruby", "rails"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
