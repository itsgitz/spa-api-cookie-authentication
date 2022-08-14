#!/usr/bin/env bash

PROJECT_NAME="gitz"

docker-compose -p $PROJECT_NAME \
    down -v

docker system prune -f

