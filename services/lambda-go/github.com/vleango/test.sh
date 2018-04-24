#!/bin/bash

# functions
go test -v functions/articles/create/main_test.go &&
go test -v functions/articles/destroy/main_test.go &&
go test -v functions/articles/index/main_test.go &&
go test -v functions/articles/show/main_test.go &&
go test -v functions/articles/update/main_test.go &&

# models
go test -v lib/models/articles_test.go &&

# lib
go test -v lib/utils/arrays_test.go
