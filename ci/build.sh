#!/bin/bash
IMAGE="$REPO:$COMMIT"

echo "Building docker: $IMAGE"
docker build -f Dockerfile -t $IMAGE .