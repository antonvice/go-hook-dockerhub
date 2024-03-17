#!/bin/sh
set -x # Print each command before executing it

IMAGE_NAME="${DOCKER_IMAGE:-<image_name>:latest}"

docker pull $IMAGE_NAME 2>&1
docker stop zimoveka 2>&1 || true # Use '|| true' to ignore errors if the container doesn't exist
docker rm zimoveka 2>&1 || true
docker run --name zimoveka -d $IMAGE_NAME 2>&1