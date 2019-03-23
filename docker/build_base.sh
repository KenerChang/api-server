#!/bin/sh

# Exit immediately if a command exits with a non-zero status
set -e

DIR="$( cd -P "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Get docker image tags and export to environment variables
set -a

BUILDROOT="${DIR}/.."

# Build docker image
GIT_HEAD="$(git rev-parse --short HEAD)"
IMAGE="api-base:$GIT_HEAD"

docker build -t $IMAGE -f Dockerfile-base  $BUILDROOT