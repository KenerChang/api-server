#!/bin/sh

# Exit immediately if a command exits with a non-zero status
set -e

DIR="$( cd -P "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Get docker image tags and export to environment variables
set -a

BUILDROOT="${DIR}/.."

# Build docker image
DATE=`date +%Y%m%d`
GIT_HEAD="$(git rev-parse --short HEAD)"
TAG="$DATE-$GIT_HEAD"
IMAGE="api-base:$TAG"

docker build -t $IMAGE -f Dockerfile-base  $BUILDROOT