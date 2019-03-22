#!/bin/bash
IMAGE="$REPO:$1"

echo "Building docker: $IMAGE"
docker build -f Dockerfile -t $IMAGE .