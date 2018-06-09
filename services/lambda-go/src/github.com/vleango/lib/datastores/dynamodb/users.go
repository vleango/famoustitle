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

var (
	ErrArticleDoesNotBelong = fmt.Errorf("resource does not belong to the user")
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

func UserCreate(user models.User, pass string, passwordConfirmation string) (*models.User, error) {
	if user.FirstName == "" || user.LastName == "" || user.Email == "" || pass == "" || passwordConfirmation == "" {
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

	user.PasswordDigest = passwordDigest
	user.ID = fmt.Sprintf("%s", uuid.Must(uuid.NewV4(), nil))
	user.Admin = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// save to DB
	av, err := dynamodbattribute.MarshalMap(user)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(userTable),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UserAddArticle(user models.User, article models.Article) error {
	var updateExpression []string
	attributeValue := map[string]*dynamodb.AttributeValue{}
	articleList := map[string]*dynamodb.AttributeValue{}

	// get article list (minus same article)
	if len(user.Articles) > 0 {
		for key, value := range user.Articles {
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
				S: aws.String(user.Email),
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

// TODO right now only destroyable/updatable by same author, might need to all any admin to do so as well
func UserArticleDestroy(user models.User, article models.Article) (*models.Article, error) {
	for id := range user.Articles {
		if id == article.ID {
			return ArticleDestroy(article)
		}
	}

	return nil, ErrArticleDoesNotBelong
}

func UserArticleUpdate(user models.User, article models.Article) (*models.Article, error) {
	for id := range user.Articles {
		if id == article.ID {
			return ArticleUpdate(article)
		}
	}

	return nil, ErrArticleDoesNotBelong
}
