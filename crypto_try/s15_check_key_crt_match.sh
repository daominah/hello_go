#!/usr/bin/env bash

export privatePath=myorg0.key
export publicPath=myorg0.crt

openssl rsa --in --noout --modulus ${privatePath}

# one of 2 following command is good, depend on cert format
openssl x509 --inform pem --in --noout --modulus ${publicPath}
openssl x509 --inform der --in --noout --modulus ${publicPath}
