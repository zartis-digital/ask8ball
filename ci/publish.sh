#!/bin/bash
REPO=$1
COMMIT=$2
BRANCH=$3
BUILD_NUMBER=$4

if [ "$BRANCH" = "master" ]; then
  docker tag $REPO:$COMMIT $REPO:latest
  # TODO: Add version tag
elif [ "$BRANCH" = "develop" ]; then
  docker tag $REPO:$COMMIT $REPO:edge
else
  docker tag $REPO:$COMMIT $REPO:$BRANCH-$BUILD_NUMBER
fi

docker push $REPO