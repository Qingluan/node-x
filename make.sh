#!/bin/bash

openssl req -x509 -newkey rsa:2048 -keyout Res/key.pem -out Res/cert.pem -days 3560 -subj "//O=Org\CN=Test" -nodes

go-bindata -o=./asset/asset.go -pkg=asset Res/...
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o node-x

