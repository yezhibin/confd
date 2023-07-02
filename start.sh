#!/bin/sh

etcd_username=root
etcd_password=qftech123!

kill -9 `pidof confd`
echo "start confd watch"

confd -watch -log-level debug -backend etcdv3 -client-ca-keys=/etc/confd/etcd/cert/ca.pem -node https://etcdserver1.com:42379 -node https://etcdserver2.com:42379 -node https://etcdserver3.com:42379 -username $etcd_username -password $etcd_password -basic-auth > /dev/null 2>&1 &

echo "confd start ok"
