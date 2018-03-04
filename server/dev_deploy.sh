#!/usr/bin/env bash
#set -e
cd gateway
./build.sh
cd -

export TLSCERT=/tls/fullchain.pem
export TLSKEY=/tls/privkey.pem
export SESSIONKEY=shoulddesignerscode?

docker rm -f git-gateway
#docker rm -f dev-mongosvr
#
## DBs
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
-v $(pwd)/gateway/tls:/tls:ro \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
-e GIT_CLIENT_ID=$GIT_CLIENT_ID \
-e GIT_CLIENT_SECRET=$GIT_CLIENT_SECRET \
-e GHTOKENS=$GHTOKENS \
-e DBADDR=dev-mongosvr:27017 \
-e SESSIONKEY=$SESSIONKEY \
-e GOENV=PROD \
$DOCKER_USER/git-gateway
