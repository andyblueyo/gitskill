#!/usr/bin/env bash

echo $DOCKER_USER

source $(pwd)/secrets.sh

export TLSCERT=/etc/letsencrypt/live/api.gitin.tech/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/api.gitin.tech/privkey.pem

export SESSIONKEY=shoulddesignerscode?

docker network create dev

# Gateway
docker rm -f git-gateway
#docker rm -f dev-mongosvr

# DBs
#docker run -d \
#--name dev-mongosvr  \
#-p 27017:27017 \
#--network dev \
#mongo

docker pull $DOCKER_USER/git-gateway
docker run -d \
--name git-gateway \
--network dev \
-p 443:443 \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
-e GIT_CLIENT_ID=$GIT_CLIENT_ID \
-e GIT_CLIENT_SECRET=$GIT_CLIENT_SECRET \
-e GHTOKEN=$GHTOKEN \
-e DBADDR=dev-mongosvr:27017 \
-e SESSIONKEY=$SESSIONKEY \
-e GOENV=PROD \
$DOCKER_USER/git-gateway

