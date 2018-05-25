#!/bin/bash

# functions
go test -v -p 1 -cover github.com/vleango/functions/... &&

# lib
go test -v -p 1 -cover github.com/vleango/lib/... &&

echo "All passed"
