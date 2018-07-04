build:
#	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/articles_index functions/articles/index/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/articles_create functions/articles/create/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/articles_show functions/articles/show/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/articles_update functions/articles/update/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/articles_destroy functions/articles/destroy/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/articles_archives_index functions/articles/archives/index/main.go

	env GOOS=linux go build -ldflags="-s -w" -o bin/tokens_create functions/tokens/create/main.go

	env GOOS=linux go build -ldflags="-s -w" -o bin/users_create functions/users/create/main.go