#!/bin/bash

export APP_IMAGE_NAME="myapp:current"

COMMAND=$1

case $COMMAND in
  "build")
    echo "Building the Docker image..."
    docker compose -f ./build/builder.yml build --no-cache
    ;;
  *)
    echo "Usage: $0 {build}"
    exit 1
    ;;
esac

