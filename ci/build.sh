#!/bin/bash
REPO=$1
COMMIT=$2
BRANCH=$3
BUILD_NUMBER=$4

docker build -f Dockerfile -t $REPO:$COMMIT .