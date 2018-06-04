package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/satori/go.uuid"
	"github.com/vleango/lib/models"
	"github.com/vleango/lib/password"
	"strings"
	"time"
)

func UserFindByEmail(email string) (*models.User, error) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(userTable),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	user := models.User{}
	dynamodbattribute.UnmarshalMap(result.Item, &user)

	if user.Email == "" {
		return nil, ErrRecordNotFound
	}

	return &user, nil
}

func UserCreate(item models.User, pass string, passwordConfirmation string) (*models.User, error) {
	if item.FirstName == "" || item.LastName == "" || item.Email == "" || pass == "" || passwordConfirmation == "" {
		return nil, fmt.Errorf("missing required params")
	}

	if len(pass) < 6 {
		return nil, fmt.Errorf("password min length is 6")
	}

	if pass != passwordConfirmation {
		return nil, fmt.Errorf("password does not match")
	}

	passwordDigest, err := password.HashPassword(pass)
	if err != nil {
		return nil, err
	}

	item.PasswordDigest = passwordDigest
	item.ID = fmt.Sprintf("%s", uuid.Must(uuid.NewV4(), nil))
	item.Admin = false
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	// save to DB
	av, err := dynamodbattribute.MarshalMap(item)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(userTable),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func UserAddArticle(item models.User, article models.Article) error {
	var updateExpression []string
	attributeValue := map[string]*dynamodb.AttributeValue{}
	articleList := map[string]*dynamodb.AttributeValue{}

	// get article list (minus same article)
	if len(item.Articles) > 0 {
		for key, value := range item.Articles {
			if key != article.ID {
				articleList[key] = &dynamodb.AttributeValue{S: aws.String(value)}
			}
		}
	}

	// add article to article list
	articleList[article.ID] = &dynamodb.AttributeValue{S: aws.String(article.Title)}

	// do dynamodb updateItem
	attributeValue[":articles"] = &dynamodb.AttributeValue{
		M: articleList,
	}
	updateExpression = append(updateExpression, "articles = :articles")

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(item.Email),
			},
		},
		ReturnValues:              aws.String("UPDATED_NEW"),
		TableName:                 aws.String(userTable),
		ExpressionAttributeValues: attributeValue,
		UpdateExpression:          aws.String("set " + strings.Join(updateExpression, ", ")),
	}

	_, err := svc.UpdateItem(input)

	return err
}
