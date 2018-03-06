#!/usr/bin/env bash

echo "**************************"
echo "*** Connecting to SSH"
echo "**************************"
echo ""
ssh -oStrictHostKeyChecking=no ubuntu@18.218.71.92 "bash -s" < run_mongoexport.sh

echo "**************************"
echo "*** Pulling Users"
echo "**************************"
echo ""
scp ubuntu@18.218.71.92:/home/ubuntu/data/users.json ./data

echo "**************************"
echo "*** Pulling Repos"
echo "**************************"
echo ""
scp ubuntu@18.218.71.92:/home/ubuntu/data/repos.json ./data

cd ./data

echo "**************************"
echo "*** Cleaning Users File"
echo "**************************"
echo ""
sed -i '$!s/$/,/' users.json
sed -i '1 i\[' users.json
sed -i "\$a]" users.json

echo "**************************"
echo "*** Cleaning Repos File"
echo "**************************"
echo ""
sed -i '$!s/$/,/' repos.json
sed -i '1 i\[' repos.json
sed -i "\$a]" repos.json

cd ..