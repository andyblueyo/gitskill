#!/usr/bin/env bash
mkdir -p tls
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj "/CN=localhost" -keyout tls/privkey.pem -out tls/fullchain.pem
