package models

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/satori/go.uuid"
	"github.com/vleango/database"
	"strings"
	"time"
)

var (
 tableName = "articles"
 svc = database.DynamoSvc
 ErrTitleBodyNotProvided = errors.New("missing title and/or body in the HTTP body")
 ErrRecordNotFound = errors.New("record not found")
)

type Article struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
  Tags      []string    `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ArticleCreate(item Article) (Article, error) {

  if item.Title == "" || item.Body == "" {
    return Article{}, ErrTitleBodyNotProvided
  }

	item.ID = fmt.Sprintf("%s", uuid.Must(uuid.NewV4(), nil))
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	av, err := dynamodbattribute.MarshalMap(item)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return Article{}, err
	}

	return item, nil
}

func ArticleDestroy(item Article) (Article, error) {
  // since deleteItem doesn't return an error, need to verify delete
  _, err := ArticleFind(item.ID)
  if err != nil {
    return item, ErrRecordNotFound
  }

	_, err = svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(item.ID),
			},
		},
	})

  if err != nil {
    return Article{}, err
  }

	return item, nil
}

func ArticleFindAll() ([]Article, error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := svc.Scan(params)
	if err != nil {
		return []Article{}, err
	}

	articles := []Article{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &articles)
	if err != nil {
		return []Article{}, err
	}

	return articles, nil
}

func ArticleFind(id string) (Article, error) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return Article{}, err
	}

	article := Article{}
	dynamodbattribute.UnmarshalMap(result.Item, &article)

  if article.ID == "" {
    return Article{}, ErrRecordNotFound
  }

	return article, nil
}

func ArticleUpdate(article Article) (Article, error) {
	if article.Title == "" || article.Body == "" {
		return Article{}, errors.New("title and/or body is blank")
	}

  article.UpdatedAt = time.Now()

	updateExpression := []string{}
	attributeValue := map[string]*dynamodb.AttributeValue{}

	if article.Title != "" {
		updateExpression = append(updateExpression, "title = :title")
		attributeValue[":title"] = &dynamodb.AttributeValue{S: aws.String(article.Title)}
	}

	if article.Body != "" {
		updateExpression = append(updateExpression, "body = :body")
		attributeValue[":body"] = &dynamodb.AttributeValue{S: aws.String(article.Body)}
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: attributeValue,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(article.ID),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		TableName:        aws.String(tableName),
		UpdateExpression: aws.String("set " + strings.Join(updateExpression, ", ")),
	}

	_, err := svc.UpdateItem(input)

	if err != nil {
		return Article{}, err
	}

	return article, nil
}
