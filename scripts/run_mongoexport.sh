#!/usr/bin/env bash

echo "**************************"
echo "*** Generating Repos Data"
echo "**************************"
echo ""
mongoexport --db git --collection repos --out ./data/repos.json

echo "**************************"
echo "*** Generating Users Data"
echo "**************************"
echo ""
mongoexport --db git --collection users --out ./data/users.json
