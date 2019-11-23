#!/bin/bash

if [ "$1" == "up" ]; then
    echo "Starting app"
    docker-compose up -d
elif [ "$1" == "down" ]; then
    echo "Stopping app"
    docker-compose down
elif [ "$1" == "build" ]; then
    echo "Building app"
    docker-compose build
elif [ "$1" == "trigger" ]; then
    echo "Triggering camera"
    docker-compose exec camera /code/trigger.sh
fi
