#!/usr/bin/env bash

set -e

if [ -z $DOCKER_USER ]; then echo "• DOCKER_USER must be a set variable.
• try running: export DOCKER_USER={user} && ./bash.sh"; exit 1; fi;

# Build Gateway
cd gateway
./build.sh
cd -

ssh -oStrictHostKeyChecking=no ubuntu@18.218.71.92 "export GOENV=PROD; export DOCKER_USER=$DOCKER_USER; bash -s" < run.sh


