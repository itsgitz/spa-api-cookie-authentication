#!/usr/bin/env bash

PROJECT_NAME="gitz"
PROJECT_FILE="docker-compose.yml"

docker-compose -f $PROJECT_FILE \
    -p $PROJECT_NAME \
    up -d --build

sleep 5

echo "Waiting for database service ..."
echo "Build go application locally ..."

GOOS=linux go build

echo "Run go application ..."
./go-api
