#!/bin/bash

appName=confd
echo "start build $appName"
rm -rf bin/*
sh build.sh

version=v1.0.0
dst=${appName}_${version}

echo "start pack $appName $version"
mkdir -p $dst/bin
cp -rf bin/confd $dst/bin
cp -rf install.sh $dst
cp -rf start.sh $dst
cp -rf etcd $dst
tar -czvf $dst.tar.gz $dst
rm -rf $dst

echo "pack $dst.tar.gz ok"
