# famoustitle Lambda-Go

## Creating new template
```
serverless sam export -o template.yml
```


## Adding extra event for options
```
Events:
  Event1:
    Type: Api
    Properties:
      Path: /articles
      Method: post
      RestApiId:
        Ref: famoustitle
  Event2:
    Type: Api
    Properties:
      Path: /articles
      Method: options
      RestApiId:
        Ref: famoustitle
```
