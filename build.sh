#!/bin/bash

appName=confd

echo "$appName build start."

appName=confd
GIT_SHA=`git rev-parse --short HEAD || echo`
mkdir -p bin

go build -ldflags "-X main.GitSHA=${GIT_SHA}" -o bin/$appName -mod=vendor

echo "$appName build succeed."