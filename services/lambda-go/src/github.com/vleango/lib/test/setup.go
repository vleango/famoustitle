package test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/vleango/database"
	"github.com/vleango/lib/models"
)

var svc = database.DynamoSvc

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
		Title: "this is my title",
		Body:  "this is my body",
		Tags:  []string{"ruby", "rails"},
	}
}
