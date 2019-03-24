#!/bin/sh

# Exit immediately if a command exits with a non-zero status
set -e

DIR="$( cd -P "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Get docker image tags and export to environment variables
set -a

BUILD_CONTEXT="${DIR}/.."
DOCKER_FILE="${DIR}/Dockerfile"

# Build docker image
GIT_HEAD="$(git rev-parse --short HEAD)"
IMAGE="go-api-server:$GIT_HEAD"
docker-compose -f ${DIR}/docker-compose.yml build