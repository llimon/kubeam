#!/usr/bin/env bash
##

curl -L -o kubectl.linux https://storage.googleapis.com/kubernetes-release/release/v1.9.0/bin/linux/amd64/kubectl

##
## create  self signed cert. We don't want to provide real information on cert signature
if [ ! -f server.key -o ! -f server.crt ]; then
  openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 \
    -subj "/C=US/ST=Denial/L=Springfield/O=Dis/CN=selfsigned.com" \
    -keyout server.key  -out server.crt
fi


