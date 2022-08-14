#!/usr/bin/env bash

PROJECT_NAME="gitz"

docker-compose -p $PROJECT_NAME \
    up -d --force-recreate
