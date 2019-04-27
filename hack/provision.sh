#!/usr/bin/env bash

apt-get update
apt-get upgrade -y
apt-get install -y unzip git build-essential

wget -q https://get.docker.com -O - | sh
usermod -aG docker vagrant

wget -q https://storage.googleapis.com/gvisor/releases/nightly/latest/runsc -O /usr/local/bin/runsc
chmod a+x /usr/local/bin/runsc

wget -q https://dl.google.com/go/go1.12.4.linux-amd64.tar.gz -O go.tar.gz
tar -C /usr/local -xzf go.tar.gz

cat <<EOF >> ~/.bashrc
export PATH=\$PATH:/usr/local/go/bin
export GOPATH=~/go
export GO111MODULE=on
export PATH=\$PATH:\$GOPATH/bin

source ~/expected/.env
EOF

curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
apt-get update
apt-get install -y migrate python-pip
pip install docker-compose
rm -rf /usr/lib/python2.7/dist-packages/OpenSSL
rm -rf /usr/lib/python2.7/dist-packages/pyOpenSSL-0.15.1.egg-info
pip install pyopenssl
