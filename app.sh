#!/bin/bash

if [ "$1" == "up" ]; then
    echo "Creating networks"
    docker network create app
    echo "Stating mattermost"
    docker-compose -f docker-compose-mattermost.yml up -d
    echo "Starting app"
    docker-compose up -d
elif [ "$1" == "down" ]; then
    echo "Stopping mattermost"
    docker-compose -f docker-compose-mattermost.yml down
    echo "Stopping app"
    docker-compose down
    echo "Removing networks"
    docker network remove app
elif [ "$1" == "build" ]; then
    echo "Building mattermost"
    docker-compose -f docker-compose-mattermost.yml build
    echo "Building app"
    docker-compose build
fi
