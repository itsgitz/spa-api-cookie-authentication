#!/usr/bin/env bash

PROJECT_NAME="gitz"
PROJECT_FILE="docker-compose.yaml"

docker-compose -f $PROJECT_FILE \
    -p $PROJECT_NAME \
    down -v

docker system prune -f

