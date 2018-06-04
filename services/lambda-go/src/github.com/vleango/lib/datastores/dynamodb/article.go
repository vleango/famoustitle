package dynamodb

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/satori/go.uuid"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/utils"
	"strings"
	"time"
)

// TODO need to separate the tags sanitized so update can use it too
func ArticleCreate(item models.Article, author string) (models.Article, error) {

	if item.Title == "" || item.Body == "" {
		return models.Article{}, ErrTitleBodyNotProvided
	}

	item.ID = fmt.Sprintf("%s", uuid.Must(uuid.NewV4(), nil))
	item.Author = author
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	var sanitizedTags []string

	if len(item.Tags) > 0 {
		for _, tag := range item.Tags {
			if tag != "" {
				trimmed := strings.TrimSpace(tag)
				sanitizedTags = append(sanitizedTags, strings.ToLower(trimmed))
			}
		}
	}
	sanitizedTags = utils.RemoveStringDuplicatesUnordered(sanitizedTags)
	item.Tags = sanitizedTags

	av, err := dynamodbattribute.MarshalMap(item)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(articleTable),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return models.Article{}, err
	}

	return item, nil
}

func ArticleDestroy(item models.Article) (models.Article, error) {
	// since deleteItem doesn't return an error, need to verify delete
	_, err := ArticleFind(item.ID)
	if err != nil {
		return item, ErrRecordNotFound
	}

	_, err = svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(articleTable),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(item.ID),
			},
		},
	})

	if err != nil {
		return models.Article{}, err
	}

	return item, nil
}

func ArticleFindAll() ([]models.Article, error) {

	params := &dynamodb.ScanInput{
		TableName: aws.String(articleTable),
	}

	result, err := svc.Scan(params)
	if err != nil {
		return []models.Article{}, err
	}

	var articles []models.Article
	articles = []models.Article{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &articles)
	if err != nil {
		return []models.Article{}, err
	}

	return articles, nil
}

func ArticleFind(id string) (models.Article, error) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(articleTable),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return models.Article{}, err
	}

	article := models.Article{}
	dynamodbattribute.UnmarshalMap(result.Item, &article)

	if article.ID == "" {
		return models.Article{}, ErrRecordNotFound
	}

	return article, nil
}

func ArticleUpdate(item models.Article) (models.Article, error) {
	var sanitizedTags []string

	if item.Title == "" && item.Body == "" {
		return models.Article{}, errors.New("title and/or body is blank")
	}

	attributeValue := map[string]*dynamodb.AttributeValue{}
	var updateExpression []string

	if item.Title != "" {
		attributeValue[":title"] = &dynamodb.AttributeValue{S: aws.String(item.Title)}
		updateExpression = append(updateExpression, "title = :title")
	}

	if item.Body != "" {
		attributeValue[":body"] = &dynamodb.AttributeValue{S: aws.String(item.Body)}
		updateExpression = append(updateExpression, "body = :body")
	}

	if len(item.Tags) > 0 {
		for _, tag := range item.Tags {
			if tag != "" {
				trimmed := strings.TrimSpace(tag)
				sanitizedTags = append(sanitizedTags, strings.ToLower(trimmed))
			}
		}
	}
	sanitizedTags = utils.RemoveStringDuplicatesUnordered(sanitizedTags)

	if len(sanitizedTags) > 0 {
		attributeValue[":tags"] = &dynamodb.AttributeValue{SS: aws.StringSlice(sanitizedTags)}
	} else {
		attributeValue[":tags"] = &dynamodb.AttributeValue{NULL: aws.Bool(true)}
	}
	updateExpression = append(updateExpression, "tags = :tags")

	attributeValue[":updated_at"] = &dynamodb.AttributeValue{S: aws.String(time.Now().Format(time.RFC3339))}
	updateExpression = append(updateExpression, "updated_at = :updated_at")

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(item.ID),
			},
		},
		ReturnValues:              aws.String("UPDATED_NEW"),
		TableName:                 aws.String(articleTable),
		ExpressionAttributeValues: attributeValue,
		UpdateExpression:          aws.String("set " + strings.Join(updateExpression, ", ")),
	}

	_, err := svc.UpdateItem(input)

	if err != nil {
		return models.Article{}, err
	}

	updatedArticle, _ := ArticleFind(item.ID)
	return updatedArticle, nil
}
