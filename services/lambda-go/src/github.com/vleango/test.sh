#!/bin/bash

if [ "$1" != "" ] && [ $1 = "--single" ] || [ "$1" != "" ] && [ $1 = "-s" ]; then
    # sample: testing individual
    go test -v -p 1 -cover github.com/vleango/functions/articles/destroy && echo "All passed"
else
    if [ "$APP_ENV" = "ci" ]; then
        # CI Parallel Tests
        if [ "$CIRCLE_NODE_INDEX" = 0 ]; then
            go test -v -p 1 -cover github.com/vleango/functions/... && echo "All passed"
        else
            if [ "$CIRCLE_NODE_INDEX" = 1 ]; then
                go test -v -p 1 -cover github.com/vleango/lib/... && echo "All passed"
            fi
        fi
    else
        # regular test
        go test -v -p 1 -cover github.com/vleango/functions/... &&
        go test -v -p 1 -cover github.com/vleango/lib/... &&
        echo "All passed"
    fi
fi
