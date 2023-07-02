#!/bin/sh

echo "start install confd"

# check installed?
if [ -d /etc/confd ]; then
    echo "confd has already installed"
    exit 1
fi

mkdir -p /etc/confd/{conf.d,templates,log}

cp -rf start.sh /etc/confd
chmod +x /etcd/confd/start.sh
cp -rf etcd /etc/confd
cp -rf bin/confd /usr/local/bin
chmod +x /usr/local/bin/confd

echo "confd install ok"
