package test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/vleango/config"
	"github.com/vleango/lib/auth"
	localDB "github.com/vleango/lib/datastores/dynamodb"
	"github.com/vleango/lib/models"
	"net/http"
	"time"
)

var svc = config.DynamoSvc

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
	url := fmt.Sprintf("%v/%v", config.ElasticSearchHost, config.DynamoArticlesTable)

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
		TableName: aws.String(config.DynamoArticlesTable),
	}

	_, err := svc.CreateTable(input)

	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
	}
}

func DefaultArticleModel() models.Article {
	return models.Article{
		Author:    "Tha",
		Title:     "this is my title",
		Body:      "this is my body",
		Tags:      []string{"ruby", "rails"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func CreateUserTable(newUsers ...map[string]interface{}) (tokens []string) {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("email"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("email"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(config.DynamoUsersTable),
	}

	_, err := svc.CreateTable(input)

	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
	}

	if len(newUsers) > 0 {
		for _, newUser := range newUsers {
			localDB.UserCreate(newUser["user"].(models.User), newUser["password"].(string), newUser["password"].(string))

			token, _ := auth.GenerateToken(newUser["user"].(models.User).Email, newUser["password"].(string))
			tokens = append(tokens, *token)
		}
	}

	return tokens
}
