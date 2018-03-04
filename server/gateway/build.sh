#!/usr/bin/env bash
set -e

echo "*******************************************************"
echo "**** Building Go API Gateway Server"
echo "*******************************************************"

unamestr=`uname`
if [ $unamestr == 'Linux' ]; then
    CGO_ENABLED=0 go build -a .
else
    GOOS=linux go build
fi
docker build -t $DOCKER_USER/git-gateway .
docker push $DOCKER_USER/git-gateway
go clean

