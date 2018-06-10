#!/bin/bash

if [ "$1" != "" ] && [ $1 = "--single" ] || [ "$1" != "" ] && [ $1 = "-s" ]; then
    # sample: testing individual
    go test -v -p 1 -cover github.com/vleango/functions/articles/destroy && echo "All passed"
else
    # functions
    go test -v -p 1 -cover github.com/vleango/functions/... &&

    # lib
    go test -v -p 1 -cover github.com/vleango/lib/... &&

    echo "All passed"
fi
