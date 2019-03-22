#!/bin/bash
BUILD_NUMBER=$TRAVIS_BUILD_NUMBER
IMAGE="$REPO:$COMMIT"

# Build from current branch
BRANCH="$TRAVIS_PULL_REQUEST_BRANCH"
if [ "$BRANCH" = "" ]; then
  BRANCH="$TRAVIS_TAG"
fi 
if [ "$BRANCH" = "" ]; then
  BRANCH="$TRAVIS_BRANCH"
fi 

if [ "$BRANCH" = "master" ]; then
  echo "Tagging image $IMAGE -> $REPO:latest"
  docker tag $IMAGE $REPO:latest
  # TODO: Add version tag
elif [ "$BRANCH" = "develop" ]; then
  echo "Tagging image $IMAGE -> $REPO:edge"
  docker tag $IMAGE $REPO:edge
else
  echo "Tagging image $IMAGE -> $REPO:$BRANCH-$BUILD_NUMBER"
  docker tag $IMAGE $REPO:$BRANCH-$BUILD_NUMBER
fi

docker push $REPO