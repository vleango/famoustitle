#!/bin/bash

# functions
go test -v github.com/vleango/functions/articles/create/ &&
go test -v github.com/vleango/functions/articles/destroy/ &&
go test -v github.com/vleango/functions/articles/index/ &&
go test -v github.com/vleango/functions/articles/show/ &&
go test -v github.com/vleango/functions/articles/update/ &&
go test -v github.com/vleango/functions/articles/archives/index/ &&

# datastores
go test -v github.com/vleango/lib/datastores/dynamodb/ &&
go test -v github.com/vleango/lib/datastores/elasticsearch &&

# models
go test -v github.com/vleango/lib/models/ &&

# lib
go test -v github.com/vleango/lib/utils/ &&

echo "All passed"