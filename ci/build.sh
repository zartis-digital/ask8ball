#!/bin/bash
REPO=$1
COMMIT=$2
BRANCH=$3
BUILD_NUMBER=$4
TAG="$REPO:$COMMIT"

echo "Building docker: $TAG"
docker build -f Dockerfile -t $TAG .