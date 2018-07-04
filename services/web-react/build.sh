#!/bin/bash

if [ "$1" != "" ] && [ $1 = "--stage=staging" ]; then
    echo "Moving env staging to production"
    mv .env.production .env.production.bk
    cp .env.staging .env.production

    echo "building staging app..."
    yarn build
    echo "build complete"

    echo "doing cleanup..."
    mv .env.production.bk .env.production
    echo "cleanup complete"

else
  if [ "$1" != "" ] && [ $1 = "--stage=production" ] ; then
    echo "building production app..."
    yarn build
    echo "build complete"
  else
    echo "Error: please supply --stage:staging or --stage:production option"
    exit 1
  fi
fi
