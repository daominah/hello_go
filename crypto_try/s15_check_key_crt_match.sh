#!/usr/bin/env bash

export privatePath=myorg0.key
export publicPath=myorg0.crt

openssl rsa --noout --modulus --in ${privatePath}

# one of 2 following command is good, depend on cert format
openssl x509 --inform pem --noout --modulus --in ${publicPath}
openssl x509 --inform der --noout --modulus --in ${publicPath}
