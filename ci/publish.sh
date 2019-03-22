#!/bin/bash
REPO=$1
COMMIT=$2
BRANCH=$3
BUILD_NUMBER=$4
TAG="$REPO:$COMMIT"

if [ "$BRANCH" = "master" ]; then
  echo "Tagging image $TAG -> $REPO:latest"
  docker tag $TAG $REPO:latest
  # TODO: Add version tag
elif [ "$BRANCH" = "develop" ]; then
  echo "Tagging image $TAG -> $REPO:edge"
  docker tag $TAG $REPO:edge
else
  echo "Tagging image $TAG -> $REPO:$BRANCH-$BUILD_NUMBER"
  docker tag $TAG $REPO:$BRANCH-$BUILD_NUMBER
fi

docker push $REPO